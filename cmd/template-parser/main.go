package main

import (
	"flag"
	"fmt"
	"maps"
	"strings"
	"template-parser/internal/combiner"
	"template-parser/internal/exporter"
	"template-parser/internal/model"
	"template-parser/internal/parser"
	"template-parser/internal/walker"
)

func main() {

	dir := flag.String("dir", ".", "Directory to scan")
	//workers := flag.Int("workers", 4, "Number of concurrent workers")
	//lang := flag.String("lang", "en", "Target language code")
	out := flag.String("out", "translations.json", "Output JSON file")
	//flag.BoolVar(&verbose, "v", false, "Enable verbose logging")
	flag.Parse()

	files, err := walker.Walk(*dir)
	if err != nil {
		panic(err)
	}

	var result = make(map[string][]model.Text)
	var merged []model.Text
	for _, f := range files {
		if strings.HasSuffix(f, ".html") {
			m, err := parser.ParseHTML(f)
			if err != nil {
				fmt.Println("Error parsing", f, ":", err)
				continue
			}
			maps.Copy(result, m)
		}
		if strings.HasSuffix(f, ".ts") {
			m, err := parser.ParseTS(f)
			if err != nil {
				fmt.Println("Error parsing", f, ":", err)
				continue
			}
			maps.Copy(result, m)
		}
	}

	merged = combiner.MergeTextsByKeyAndContext(result)
	exporter.ExportToJSON(merged, "en", *out)

}
