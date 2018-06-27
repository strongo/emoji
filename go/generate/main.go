package main

import (
	"io/ioutil"
	"log"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"fmt"
)

const gemojiDBJsonURL = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"

type GemojiEmoji struct {
	Aliases     []string `json:"aliases"`
	Category    string   `json:"category"`
	Description string   `json:"description"`
	Emoji       string   `json:"emoji"`
	Tags        []string `json:"tags"`
}

type TemplateData struct {
	PkgName string
	CodeMap map[string]string
}

const templateMapCode = `
package {{.PkgName}}

// NOTE: THIS FILE WAS PRODUCED BY THE
// EMOJICODEMAP CODE GENERATION TOOL (github.com/kyokomi/generateEmojiCodeMap)
// DO NOT EDIT

// Mapping from character to concrete escape code.
var emojiCodeMap = map[string]string{
	{{range $key, $val := .CodeMap}}":{{$key}}:": {{$val}},
{{end}}
}
`

//var pkgName string
//var fileName string
//
//func init() {
//	log.SetFlags(log.Llongfile)
//
//	flag.StringVar(&pkgName, "pkg", "main", "output package")
//	flag.StringVar(&fileName, "o", "emoji_codemap.go", "output file")
//	flag.Parse()
//}

func main() {
	gemojis, err := getGemoji(gemojiDBJsonURL)
	if err != nil {
		log.Fatalln(err)
	}

	if err = newGenerator().Generate(gemojis); err != nil {
		log.Fatalln(err)
	}
}

type generator struct {
	categoryFiles map[string]*os.File
}

func newGenerator() generator {
	return generator{
		categoryFiles: make(map[string]*os.File, 20),
	}
}

func (g generator) getCategoryFile(category string) (f *os.File, err error) {
	if category == "" {
		return nil, nil
	}
	category = strings.ToLower(category)
	var ok bool
	if f, ok = g.categoryFiles[category]; ok {
		return
	}
	fileName := "../category_" + category + ".go"
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

func (g generator) Generate(gemojis []GemojiEmoji) (err error) {
	emojiMapFileName := "../emoji_all_map.go"
	os.Remove(emojiMapFileName)
	emojiMapFile, err := os.Create(emojiMapFileName)
	if err != nil {
		return
	}
	defer func() {
		if _, err = emojiMapFile.WriteString("}"); err != nil {
			return
		}
		emojiMapFile.Close()
	}()
	if _, err = emojiMapFile.WriteString(`package emoji

var Categories = map[string][]string {
`); err != nil {
		return
	}

	defer func() {
		for _, categoryFile := range g.categoryFiles {
			if _, err = fmt.Fprintln(categoryFile, "}"); err != nil {
				return
			}
			categoryFile.Close()
		}
	}()
	for _, gemoji := range gemojis {
		if gemoji.Category != "" {
			categoryFile, err := g.getCategoryFile(gemoji.Category)
			if err != nil {
				return err
			}

			if _, err = fmt.Fprintf(categoryFile, "\t\"%v\", // %v, tags=%v, aliases=%v\n", gemoji.Emoji, gemoji.Description, gemoji.Tags, gemoji.Aliases); err != nil {
				return err
			}
		}
	}
	return nil
}

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

//func generateJson(pkgName string) ([]byte, error) {
//
//	// Read Emoji file
//
//
//	emojiCodeMap := make(map[string]string)
//	for _, gemoji := range gs {
//		for _, a := range gemoji.Aliases {
//			emojiCodeMap[a] = fmt.Sprintf("%+q", gemoji.Emoji)
//		}
//	}
//
//	// Template GenerateSource
//
//	var buf bytes.Buffer
//	t := template.Must(template.New("template").Parse(templateMapCode))
//	if err := t.Execute(&buf, TemplateData{pkgName, emojiCodeMap}); err != nil {
//		return nil, err
//	}
//
//	// gofmt
//
//	bts, err := format.Source(buf.Bytes())
//	if err != nil {
//		fmt.Println(string(buf.Bytes()))
//		return nil, fmt.Errorf("gofmt: %s", err)
//	}
//
//	return bts, nil
//}
