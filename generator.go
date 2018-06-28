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
	if _, err = fmt.Fprintf(f, `package emoji

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

	if err = g.generateMapFile(); err != nil {
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

func (g goGenerator) generateMapFile() (err error) {
	emojiMapFileName := "./go/emoji_all_map.go"
	os.Remove(emojiMapFileName)
	emojiMapFile, err := os.Create(emojiMapFileName)
	if err != nil {
		return
	}
	defer func() {
		emojiMapFile.Close()
	}()
	if _, err = emojiMapFile.WriteString(`package emoji

var ByCategory = map[string][]string {
`); err != nil {
		return
	}

	categoryWriters := g.writers["category"]
	categories := make([]string, len(categoryWriters))

	var i int
	for category, _ := range categoryWriters {
		categories[i] = category
		i++
	}
	sort.Strings(categories)
	for _, category := range categories {
		if _, err = fmt.Fprintf(emojiMapFile, "\t"+`"%v": Category%v,`+"\n", category, strings.Title(category)); err != nil {
			return
		}
	}
	if _, err = emojiMapFile.WriteString("}"); err != nil {
		return
	}

	return nil
}
