package parser

import (
	"os"
	"regexp"
	"strings"
	"template-parser/internal/model"

	"golang.org/x/net/html"
)

type NlsString struct {
	Text    string
	Context string
}

var nlsPipeRegex = regexp.MustCompile(`\{\{\s*['"](.+?)['"]\s*\|\s*nls\s*\}\}`)
var nlsFuncRegex = regexp.MustCompile(`\{\{\s*nls\.tr\(\s*['"](.+?)['"]\s*\)\s*\}\}`)

func getAttr(n *html.Node, key string) (string, bool) {
	key = strings.ToLower(key)
	for _, a := range n.Attr {
		if a.Key == key {
			return a.Val, true
		}
	}
	return "", false
}

func traverse(n *html.Node, currentContext string, results *[]NlsString) {
	if ctx, ok := getAttr(n, "nlsContext"); ok {
		currentContext = ctx
	}

	if n.Type == html.TextNode {
		matches := nlsPipeRegex.FindAllStringSubmatch(n.Data, -1)
		if len(matches) > 0 {
			for _, m := range matches {
				if len(m) > 1 {
					*results = append(*results, NlsString{
						Text:    m[1],
						Context: currentContext,
					})
				}
			}
		} else {
			// Only check function style if no pipe matches
			matchesFunc := nlsFuncRegex.FindAllStringSubmatch(n.Data, -1)
			for _, m := range matchesFunc {
				if len(m) > 1 {
					*results = append(*results, NlsString{
						Text:    m[1],
						Context: currentContext,
					})
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, currentContext, results)
	}
}

func ParseHTML(path string) (map[string][]model.Text, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		return nil, err
	}

	var finalResults = make(map[string][]model.Text)
	var results []NlsString
	traverse(doc, "", &results)

	for _, r := range results {
		//fmt.Printf("Text: %q, Context: %q\n", r.Text, r.Context)
		finalResults[path] = append(finalResults[path], model.Text{Key: r.Text, Context: r.Context})
	}

	return finalResults, nil
}
