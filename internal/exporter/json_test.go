package exporter_test

import (
	"bytes"
	"encoding/json"
	"template-parser/internal/exporter"
	"template-parser/internal/model"
	"testing"
)

func TestExportToJSONWriter(t *testing.T) {
	value1 := "Hello"
	value2 := "World"
	value3 := "Optional text"

	texts := []model.Text{
		{
			Key:     "greeting",
			Value:   &value1,
			Context: "Used in homepage",
			Comment: "Friendly greeting",
			Sources: []string{"ui", "backend"},
			State:   "approved",
		},
		{
			Key:     "farewell",
			Value:   &value2,
			Context: "Used in email",
			Sources: []string{"backend"},
			State:   "pending",
		},
		{
			Key:   "optional_text",
			Value: &value3,
			// Context, Comment, Sources and State are omitted
		},
		{
			Key:   "nil_value_text",
			Value: nil,
			State: "draft",
		},
	}

	lang := "en"
	var buf bytes.Buffer

	err := exporter.ExportToJSONWriter(texts, lang, &buf)
	if err != nil {
		t.Fatalf("ExportToJSONWriter returned error: %v", err)
	}

	var result model.Result
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if result.Language != lang {
		t.Errorf("expected language %q, got %q", lang, result.Language)
	}

	if len(result.Texts) != len(texts) {
		t.Errorf("expected %d texts, got %d", len(texts), len(result.Texts))
	}

	for i, txt := range texts {
		if result.Texts[i].Key != txt.Key {
			t.Errorf("expected Key %v, got %v", txt.Key, result.Texts[i].Key)
		}
		if (txt.Value == nil && result.Texts[i].Value != nil) ||
			(txt.Value != nil && result.Texts[i].Value != nil && *txt.Value != *result.Texts[i].Value) {
			t.Errorf("expected Value %v, got %v", txt.Value, result.Texts[i].Value)
		}
		if result.Texts[i].Context != txt.Context {
			t.Errorf("expected Context %v, got %v", txt.Context, result.Texts[i].Context)
		}
		if result.Texts[i].Comment != txt.Comment {
			t.Errorf("expected Comment %v, got %v", txt.Comment, result.Texts[i].Comment)
		}
		if len(result.Texts[i].Sources) != len(txt.Sources) {
			t.Errorf("expected Sources %v, got %v", txt.Sources, result.Texts[i].Sources)
		}
		if result.Texts[i].State != txt.State {
			t.Errorf("expected State %v, got %v", txt.State, result.Texts[i].State)
		}
	}
}
