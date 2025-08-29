package combiner_test

import (
	"reflect"
	"template-parser/internal/combiner"
	"template-parser/internal/model"
	"testing"
)

func TestMergeTextsByKeyAndContext(t *testing.T) {
	src := map[string][]model.Text{
		"file1.json": {
			{Key: "greeting", Value: ptr("Hello"), Context: "homepage"},
			{Key: "farewell", Value: ptr("Bye"), Context: "email"},
		},
		"file2.json": {
			{Key: "greeting", Value: ptr("Hello"), Context: "homepage"},
			{Key: "optional_text", Value: ptr("Optional")},
		},
		"file3.json": {
			{Key: "greeting", Value: ptr("Hello"), Context: "email"},
		},
	}

	expected := []model.Text{
		{Key: "greeting", Value: ptr("Hello"), Context: "homepage", Sources: []string{"file1.json", "file2.json"}},
		{Key: "farewell", Value: ptr("Bye"), Context: "email", Sources: []string{"file1.json"}},
		{Key: "optional_text", Value: ptr("Optional"), Sources: []string{"file2.json"}},
		{Key: "greeting", Value: ptr("Hello"), Context: "email", Sources: []string{"file3.json"}},
	}

	result := combiner.MergeTextsByKeyAndContext(src)

	for _, exp := range expected {
		found := false
		for _, r := range result {
			if exp.Key == r.Key && exp.Context == r.Context &&
				reflect.DeepEqual(exp.Sources, r.Sources) &&
				((exp.Value == nil && r.Value == nil) || (exp.Value != nil && r.Value != nil && *exp.Value == *r.Value)) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected text %+v not found in result", exp)
		}
	}

	if len(result) != len(expected) {
		t.Errorf("expected %d merged texts, got %d", len(expected), len(result))
	}
}

func ptr(s string) *string {
	return &s
}
