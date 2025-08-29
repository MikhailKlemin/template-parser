package combiner

import "template-parser/internal/model"

func MergeTextsByKeyAndContext(src map[string][]model.Text) []model.Text {
	merged := make(map[string]model.Text)

	for path, texts := range src {
		for _, t := range texts {
			id := t.Key + "||" + t.Context

			if existing, ok := merged[id]; ok {
				existing.Sources = append(existing.Sources, path)
				merged[id] = existing
			} else {
				t.Sources = []string{path}
				merged[id] = t
			}
		}
	}

	result := make([]model.Text, 0, len(merged))
	for _, v := range merged {
		result = append(result, v)
	}

	return result
}
