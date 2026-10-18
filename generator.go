package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	title           = cases.Title(language.English)
	gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
	logFatal        = log.Fatalln
)

func main() {
	gemojis, err := getGemoji(gemojiDBJsonURL)
	if err != nil {
		logFatal(err)
		return
	}

	if err = newGoGenerator().Generate(gemojis); err != nil {
		logFatal(err)
		return
	}
}

type GemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

func getGemoji(url string) (gs []GemojiEmoji, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = res.Body.Close()
	}()

	emojiFile, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(emojiFile, &gs); err != nil {
		return nil, err
	}
	return
}

type goGenerator struct {
	writers map[string]map[string]io.Writer
}

func newGoGenerator() goGenerator {
	return goGenerator{
		writers: map[string]map[string]io.Writer{
			"category": make(map[string]io.Writer, 20),
		},
	}
}

func (g goGenerator) getWriter(kind, id string) (f io.Writer, err error) {
	if id == "" {
		return nil, nil
	}
	writers := g.writers[kind]
	id = strings.ToLower(id)
	var ok bool
	if f, ok = writers[id]; ok {
		return
	}
	fileName := fmt.Sprintf("./go/%v_%v.go", kind, id)
	if f, err = os.Create(fileName); err != nil {
		return
	}
	writers[id] = f
	_, err = fmt.Fprintf(f, `package emojis

var Category%v = []string {
`, title.String(id))
	return
}

func (g goGenerator) Generate(gemojis []GemojiEmoji) (err error) {
	defer func() {
		for _, writers := range g.writers {
			for _, w := range writers {
				_, _ = fmt.Fprintln(w, "}")
				if closer, ok := w.(io.Closer); ok {
					_ = closer.Close()
				}
			}
		}
	}()

	for _, gemoji := range gemojis {
		if gemoji.Emoji == "" {
			continue
		}
		if gemoji.Category == "" {
			gemoji.Category = "uncategorized"
		}
		if err = g.writeItem("category", gemoji.Category, gemoji); err != nil {
			return
		}
	}

	categories := make([]string, len(g.writers["category"]))
	{
		var i int
		for category := range g.writers["category"] {
			categories[i] = category
			i++
		}
	}

	if err = g.generateMapFile(categories, gemojis); err != nil {
		return
	}

	if err = g.generateConstants(gemojis); err != nil {
		return
	}

	return nil
}

func (g goGenerator) writeItem(kind, id string, gemoji GemojiEmoji) error {
	categoryFile, err := g.getWriter(kind, id)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(categoryFile, "\t\"%v\", // %v, tags=%v, aliases=%v\n", gemoji.Emoji, gemoji.Description, gemoji.Tags, gemoji.Aliases)
	return err
}

func (g goGenerator) generateMapFile(categories []string, gemojis []GemojiEmoji) (err error) {
	const fileName = "./go/emoji_all_map.go"
	var buf strings.Builder
	buf.WriteString(`package emojis

var ByCategory = map[string][]string {
`)

	sort.Strings(categories)
	for _, category := range categories {
		fmt.Fprintf(&buf, "\t\"%v\": Category%v,\n", category, title.String(category))
	}

	buf.WriteString(`}

var All = map[string]Info {
`)

	for _, gemoji := range gemojis {
		if gemoji.Emoji == "" {
			continue
		}
		fmt.Fprintf(&buf, "\t\"%v\": {Description: \"%v\", Category: \"%v\"},\n", gemoji.Emoji, gemoji.Description, gemoji.Category)
	}

	buf.WriteString("}")
	return os.WriteFile(fileName, []byte(buf.String()), 0o644)
}

func (goGenerator) generateConstants(gemojis []GemojiEmoji) (err error) {
	const fileName = "./go/emoji/constants.go"
	var buf strings.Builder
	buf.WriteString(`package emoji

const (
`)
	for _, gemoji := range gemojis {
		if gemoji.Description == "" {
			continue
		}
		var name string
		switch gemoji.Emoji {
		case "🇹🇷":
			name = "TurkeyCountry"
		case "🏘":
			name = "House2"
		default:
			desc := strings.Split(gemoji.Description, " ")
			for i, d := range desc {
				if d == "" {
					continue
				}
				d = strings.ReplaceAll(d, "&", "And")
				d = strings.ReplaceAll(d, "’", "")
				d = strings.ReplaceAll(d, "-", "")
				d = strings.ReplaceAll(d, ".", "")
				d = strings.ReplaceAll(d, ":", "")
				d = strings.ReplaceAll(d, ";", "")
				d = strings.ReplaceAll(d, ",", "")
				d = strings.ReplaceAll(d, "“", "")
				d = strings.ReplaceAll(d, "”", "")
				d = strings.ReplaceAll(d, "(", "")
				d = strings.ReplaceAll(d, ")", "")
				d = strings.ReplaceAll(d, "!", "")
				d = strings.ReplaceAll(d, "#", "Number")
				d = strings.ReplaceAll(d, "*", "Star")
				d = strings.ReplaceAll(d, "1st", "First")
				d = strings.ReplaceAll(d, "2nd", "Second")
				d = strings.ReplaceAll(d, "3rd", "Third")
				desc[i] = title.String(d)
			}
			name = strings.Join(desc, "")
		}

		fmt.Fprintf(&buf, "\t%v = \"%v\"\n", name, gemoji.Emoji)
	}
	buf.WriteString(")")
	return os.WriteFile(fileName, []byte(buf.String()), 0o644)
}
