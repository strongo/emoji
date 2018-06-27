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
	categoryFiles map[string]*os.File
}

func newGoGenerator() goGenerator {
	return goGenerator{
		categoryFiles: make(map[string]*os.File, 20),
	}
}

func (g goGenerator) getCategoryFile(category string) (f *os.File, err error) {
	if category == "" {
		return nil, nil
	}
	category = strings.ToLower(category)
	var ok bool
	if f, ok = g.categoryFiles[category]; ok {
		return
	}
	fileName := "./go/category_" + category + ".go"
	if f, err = os.Create(fileName); err != nil {
		return
	}
	g.categoryFiles[category] = f
	if _, err = fmt.Fprintf(f, `package emoji

var Category%v = []string {
`, strings.Title(category)); err != nil {
		return
	}
	return
}

func (g goGenerator) Generate(gemojis []GemojiEmoji) (err error) {

	defer func() {
		for _, categoryFile := range g.categoryFiles {
			if _, err = fmt.Fprintln(categoryFile, "}"); err != nil {
				return
			}
			categoryFile.Close()
		}
	}()
	for _, gemoji := range gemojis {
		if gemoji.Emoji == "" {
			continue
		}
		if gemoji.Category == "" {
			gemoji.Category = "uncategorized"
		}
		categoryFile, err := g.getCategoryFile(gemoji.Category)
		if err != nil {
			return err
		}

		if _, err = fmt.Fprintf(categoryFile, "\t\"%v\", // %v, tags=%v, aliases=%v\n", gemoji.Emoji, gemoji.Description, gemoji.Tags, gemoji.Aliases); err != nil {
			return err
		}
	}

	if err = g.generateMapFile(); err != nil {
		return
	}

	return nil
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
	categories := make([]string, len(g.categoryFiles))

	var i int
	for category, _ := range g.categoryFiles {
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
