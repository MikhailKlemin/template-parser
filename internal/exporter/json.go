package exporter

import (
	"encoding/json"
	"io"
	"template-parser/internal/model"
)

/* func ExportToJSON(texts []model.Text, lang, outPath string) error {

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
*/

func ExportToJSONWriter(texts []model.Text, lang string, w io.Writer) error {
	var result model.Result
	result.Language = lang
	result.Texts = texts

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	return enc.Encode(result)
}
