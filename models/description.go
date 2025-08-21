package models

import "golangR99/constants"

type Raw struct {
	Data     map[string]any `json:"data,omitempty"` // remove?
	Content  []ContentItem  `json:"content,omitempty"`
	NodeType string         `json:"nodeType,omitempty"`
}

func (r Raw) GetString() string {
	return parseContentArr(r.Content)
}

type ContentItem struct {
	Data     map[string]any `json:"data,omitempty"` // remove?
	Content  []ContentItem  `json:"content,omitempty"`
	Marks    []Mark         `json:"marks,omitempty"` // remove?
	Value    string         `json:"value,omitempty"`
	NodeType string         `json:"nodeType,omitempty"`
}

func (ci ContentItem) GetString() string {
	switch ci.NodeType {
	case constants.TEXT:
		return ci.Value
	default:
		return parseContentArr(ci.Content)
	}
}

type Mark struct {
	Type string `json:"type,omitempty"`
}

// helper function for contentItem
func parseContentArr(ci []ContentItem) string {
	resp := ""
	for _, val := range ci {
		resp += val.GetString()
	}
	return resp
}
