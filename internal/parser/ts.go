package parser

import (
	"os"
	"regexp"
	"template-parser/internal/model"
)

var tsRegex = regexp.MustCompile(`nls\.tr\(\s*"(.*?)"\s*\)|nls\.tr\(\s*'(.*?)'\s*\)`)

func ParseTS(path string) (map[string][]model.Text, error) {
	var finalResults = make(map[string][]model.Text)

	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	matches := tsRegex.FindAllSubmatch(file, -1)

	for _, m := range matches {
		var key string
		if len(m[1]) > 0 {
			key = string(m[1])
		} else if len(m[2]) > 0 {
			key = string(m[2])
		}
		//fmt.Println("Found:", key)
		finalResults[path] = append(finalResults[path], model.Text{Key: key, Context: ""})
	}

	return finalResults, nil
}
