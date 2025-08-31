package exporter

import (
	"encoding/json"
	"io"
	"template-parser/internal/model"
)

type ExportOptions struct {
	Pretty         bool
	IncludeSources bool
}

type jsonTextMinimal struct {
	Key   string  `json:"Key"`
	Value *string `json:"Value"`
}

type jsonTextFull struct {
	Key     string   `json:"Key"`
	Value   *string  `json:"Value"`
	Sources []string `json:"Sources,omitempty"`
}

type jsonResult[T any] struct {
	Language string `json:"Language"`
	Texts    []T    `json:"Texts"`
}

func ExportToJSONWriter(texts []model.Text, lang string, w io.Writer, opts ExportOptions) error {
	var data any

	if opts.IncludeSources {
		out := make([]jsonTextFull, len(texts))
		for i, t := range texts {
			out[i] = jsonTextFull{
				Key:     t.Key,
				Value:   t.Value,
				Sources: t.Sources,
			}
		}
		data = jsonResult[jsonTextFull]{Language: lang, Texts: out}
	} else {
		out := make([]jsonTextMinimal, len(texts))
		for i, t := range texts {
			out[i] = jsonTextMinimal{
				Key:   t.Key,
				Value: t.Value,
			}
		}
		data = jsonResult[jsonTextMinimal]{Language: lang, Texts: out}
	}

	enc := json.NewEncoder(w)
	if opts.Pretty {
		enc.SetIndent("", "    ")
	}
	return enc.Encode(data)
}
