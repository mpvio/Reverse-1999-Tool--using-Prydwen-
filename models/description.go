package models

type Raw struct {
	Data     map[string]any `json:"data,omitempty"` // or use struct{} if you know it's always empty
	Content  []ContentItem  `json:"content,omitempty"`
	NodeType string         `json:"nodeType,omitempty"`
}

type ContentItem struct {
	Data     map[string]any `json:"data,omitempty"`
	Content  []ContentItem  `json:"content,omitempty"`
	Marks    []Mark         `json:"marks,omitempty"`
	Value    string         `json:"value,omitempty"`
	NodeType string         `json:"nodeType,omitempty"`
}

type Mark struct {
	Type string `json:"type,omitempty"`
}
