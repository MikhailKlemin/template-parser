package model

type Text struct {
	Key     string   `json:"Key"`
	Value   *string  `json:"Value"`
	Context string   `json:"Context,omitempty"`
	Comment string   `json:"Comment,omitempty"`
	Sources []string `json:"Sources,omitempty"`
	State   string   `json:"State,omitempty"`
}

// Result represents the final export object
type Result struct {
	Language string `json:"language"`
	Texts    []Text `json:"texts"`
}
