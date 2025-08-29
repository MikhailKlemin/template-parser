package exporter

import (
	"encoding/json"
	"os"
	"template-parser/internal/model"
)

func ExportToJSON(texts []model.Text, lang, outPath string) error {

	var result model.Result
	result.Language = lang
	result.Texts = texts

	f, err := os.Create(outPath)
	if err != nil {
		return err
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", "    ")
	return enc.Encode(result)
}
