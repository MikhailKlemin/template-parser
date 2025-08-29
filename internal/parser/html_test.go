package parser

import (
	"fmt"
	"path/filepath"
	"testing"
)

func TestParseAngularFile(t *testing.T) {
	path := filepath.Join("..", "..", "testdata", "angular", "sample.component.html")
	fmt.Println(path)
	results, err := ParseHTML(path)
	if err != nil {
		t.Fatalf("ParseAngularFile failed: %v", err)
	}

	expected := [][]string{
		{"Recipe", "components.recipeInfo"},
		{"Iterations", "components.recipeInfo"},
		{"Photobleaching", "photobleaching"},
		{"Exposure", "exposure"},
		{"Number of accumulations", "components.recipeInfo"},
	}

	/*if len(results) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(results))
	}*/

	for _, texts := range results {
		for i, text := range texts {
			if text.Key != expected[i][0] {
				t.Errorf("expected %q, got %q", expected[i][0], text.Key)
			}
			if text.Context != expected[i][1] {
				t.Errorf("expected %q, got %q", expected[i][1], text.Context)
			}

		}
	}
}
