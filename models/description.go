package models

// for determining differences
const (
	TYPE  = "type"
	SIZE  = "size"
	VALUE = "value"
)

type Raw struct {
	Data     map[string]any `json:"data,omitempty"` // remove?
	Content  []ContentItem  `json:"content,omitempty"`
	NodeType string         `json:"nodeType,omitempty"` // remove?
}

type ContentItem struct {
	Data     map[string]any `json:"data,omitempty"` // remove?
	Content  []ContentItem  `json:"content,omitempty"`
	Marks    []Mark         `json:"marks,omitempty"` // remove?
	Value    string         `json:"value,omitempty"`
	NodeType string         `json:"nodeType,omitempty"`
}

type Mark struct {
	Type string `json:"type,omitempty"`
}
