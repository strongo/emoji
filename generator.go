package main

import (
	"log"
	"os"
	"strings"
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"sort"
	"io"
)

func main() {
	gemojis, err := getGemoji(gemojiDBJsonURL)
	if err != nil {
		log.Fatalln(err)
	}

	if err = newGoGenerator().Generate(gemojis); err != nil {
		log.Fatalln(err)
	}
}

type GemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

const gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"

func getGemoji(url string) (gs []GemojiEmoji, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	emojiFile, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(emojiFile, &gs); err != nil {
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
			// "tag":      make(map[string]io.Writer, 20),
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
	if _, err = fmt.Fprintf(f, `package emojis

var Category%v = []string {
`, strings.Title(id)); err != nil {
		return
	}
	return
}

func (g goGenerator) Generate(gemojis []GemojiEmoji) (err error) {
	defer func() {
		for _, writers := range g.writers {
			for _, w := range writers {
				if _, err = fmt.Fprintln(w, "}"); err != nil {
					return
				}
				if err = w.(io.Closer).Close(); err != nil {
					return
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
		// for _, tag := range gemoji.Tags {
		// 	if err = g.writeItem("tag", tag, gemoji); err != nil {
		// 		return
		// 	}
		// }
	}

	categories := make([]string, len(g.writers["category"]))
	{
		var i int
		for category, _ := range g.writers["category"] {
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

	if _, err = fmt.Fprintf(categoryFile, "\t\"%v\", // %v, tags=%v, aliases=%v\n", gemoji.Emoji, gemoji.Description, gemoji.Tags, gemoji.Aliases); err != nil {
		return err
	}
	return err
}

func (g goGenerator) generateMapFile(categories []string, gemojis []GemojiEmoji) (err error) {
	const fileName = "./go/emoji_all_map.go"
	os.Remove(fileName)
	f, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()

	if _, err = f.WriteString(`package emojis

var ByCategory = map[string][]string {
`); err != nil {
		return
	}

	sort.Strings(categories)
	for _, category := range categories {
		if _, err = fmt.Fprintf(f, "\t"+`"%v": Category%v,`+"\n", category, strings.Title(category)); err != nil {
			return
		}
	}

	if _, err = f.WriteString(`}

var All = map[string]Info {
`); err != nil {
		return
	}

	for _, gemoji := range gemojis {
		if gemoji.Emoji == "" {
			continue
		}
		fmt.Fprintf(f, "\t"+`"%v": {Description: "%v", Category: "%v"},`+"\n", gemoji.Emoji, gemoji.Description, gemoji.Category)
	}

	if _, err = f.WriteString("}"); err != nil {
		return
	}

	return
}

func (goGenerator) generateConstants(gemojis []GemojiEmoji) (err error) {
	const fileName = "./go/emoji/constants.go"
	os.Remove(fileName)
	f, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer func() {
		f.Close()
	}()

	if _, err = f.WriteString(`package emoji

const (
`); err != nil {
		return
	}
	for _, gemoji := range gemojis {
		if gemoji.Description == "" {
			continue
		}
		var name string
		if gemoji.Emoji == "🇹🇷" {
			name = "TurkeyCountry"
		} else if gemoji.Emoji == "🏘" {
			name = "House2"
		} else {
			desc := strings.Split(gemoji.Description, " ")
			for i, d := range desc {
				if d == "" {
					continue
				}
				d = strings.Replace(d, "&", "And", -1)
				d = strings.Replace(d, "’", "", -1)
				d = strings.Replace(d, "-", "", -1)
				d = strings.Replace(d, ".", "", -1)
				d = strings.Replace(d, ":", "", -1)
				d = strings.Replace(d, ";", "", -1)
				d = strings.Replace(d, ",", "", -1)
				d = strings.Replace(d, "“", "", -1)
				d = strings.Replace(d, "”", "", -1)
				d = strings.Replace(d, "(", "", -1)
				d = strings.Replace(d, ")", "", -1)
				d = strings.Replace(d, "!", "", -1)
				d = strings.Replace(d, "#", "Number", -1)
				d = strings.Replace(d, "*", "Star", -1)
				d = strings.Replace(d, "1st", "First", -1)
				d = strings.Replace(d, "2nd", "Second", -1)
				d = strings.Replace(d, "3rd", "Third", -1)
				desc[i] = strings.Title(d)
			}
			name = strings.Join(desc, "")
		}

		fmt.Fprintf(f, "\t%v = \"%v\"\n", name, gemoji.Emoji)
	}
	if _, err = f.WriteString(")"); err != nil {
		return
	}
	return
}
