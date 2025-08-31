package exporter_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"template-parser/internal/exporter"
	"template-parser/internal/model"
)

func TestExportToJSONWriter_Full(t *testing.T) {
	value1 := "Hello"
	value2 := "World"
	value3 := "Optional text"

	texts := []model.Text{
		{Key: "greeting", Value: &value1, Context: "Used in homepage", Comment: "Friendly greeting", Sources: []string{"ui", "backend"}, State: "approved"},
		{Key: "farewell", Value: &value2, Context: "Used in email", Sources: []string{"backend"}, State: "pending"},
		{Key: "optional_text", Value: &value3},
		{Key: "nil_value_text", Value: nil, State: "draft"},
	}

	lang := "en"
	var buf bytes.Buffer

	opts := exporter.ExportOptions{Pretty: false, IncludeSources: true}
	err := exporter.ExportToJSONWriter(texts, lang, &buf, opts)
	if err != nil {
		t.Fatalf("ExportToJSONWriter returned error: %v", err)
	}

	var result struct {
		Language string       `json:"Language"`
		Texts    []model.Text `json:"Texts"`
	}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("failed to unmarshal JSON: %v", err)
	}

	if result.Language != lang {
		t.Errorf("expected language %q, got %q", lang, result.Language)
	}
	if len(result.Texts) != len(texts) {
		t.Errorf("expected %d texts, got %d", len(texts), len(result.Texts))
	}
}

func TestExportToJSONWriter_Minimal(t *testing.T) {
	value := "Hello"
	texts := []model.Text{
		{Key: "greeting", Value: &value, Sources: []string{"ui", "backend"}},
	}

	var buf bytes.Buffer
	opts := exporter.ExportOptions{Pretty: false, IncludeSources: false}
	err := exporter.ExportToJSONWriter(texts, "en", &buf, opts)
	if err != nil {
		t.Fatalf("ExportToJSONWriter returned error: %v", err)
	}

	// minimal struct with no Sources field
	var result struct {
		Language string `json:"Language"`
		Texts    []struct {
			Key   string  `json:"Key"`
			Value *string `json:"Value"`
		} `json:"Texts"`
	}
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Fatalf("failed to unmarshal minimal JSON: %v", err)
	}

	if len(result.Texts) != 1 {
		t.Fatalf("expected 1 text, got %d", len(result.Texts))
	}
	if result.Texts[0].Key != "greeting" {
		t.Errorf("expected Key 'greeting', got %q", result.Texts[0].Key)
	}
	if result.Texts[0].Value == nil || *result.Texts[0].Value != "Hello" {
		t.Errorf("expected Value 'Hello', got %v", result.Texts[0].Value)
	}
}
