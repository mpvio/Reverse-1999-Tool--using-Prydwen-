package models

import (
	"fmt"
	"golangR99/constants"
	"strings"
)

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

func (r Raw) GetString() string {
	return parseContentArr(r.Content)
}

func (r Raw) GetDiff(other Raw) string {
	return contentArrDiff(r.Content[0].Content, other.Content[0].Content)
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

func (ci ContentItem) HasValue() bool {
	return ci.Value != ""
}

func (ci ContentItem) Equals(other ContentItem) (bool, string) {
	if ci.NodeType != other.NodeType {
		return false, TYPE
	}
	if len(ci.Content) != len(other.Content) {
		return false, SIZE
	}
	if strings.TrimSpace(ci.Value) != strings.TrimSpace(other.Value) {
		return false, VALUE
	}
	// iterate through both contents:
	for i, val := range ci.Content {
		equal, how := val.Equals(other.Content[i])
		if !equal {
			return equal, how
		}
	}
	return true, ""
}

func (ci ContentItem) GetDiff(other ContentItem) string {
	equal, _ := ci.Equals(other)
	// fmt.Println("INNER:", ci.GetString(), other.GetString(), equal)
	if equal {
		// fmt.Println("EQUAL")
		return ci.GetString()
	} else {
		// fmt.Println("INEQUAL", ci.HasValue(), other.HasValue())
		if ci.HasValue() && other.HasValue() {
			// fmt.Println("BOTH HAVE VALUES")
			return fmt.Sprintf("{%s} -> {%s}", ci.GetString(), other.GetString())
		}
		// iterate through both contentItem lists within ci and other:
		// fmt.Println("ITERATING!")
		return contentArrDiff(ci.Content, other.Content)
	}
}

type Mark struct {
	Type string `json:"type,omitempty"`
}

// helper functions for []contentItem
func parseContentArr(ci []ContentItem) string {
	resp := ""
	for _, val := range ci {
		resp += val.GetString()
	}
	return resp
}

func contentArrDiff(a, b []ContentItem) string {
	resp := ""
	for i, val := range a {
		// fmt.Println("CHECKING:", val.GetString(), b[i].GetString())
		resp += val.GetDiff(b[i])
	}
	// fmt.Println("CURRENTLY:", resp)
	if len(b) > len(a) {
		start := len(a)
		end := len(b)
		extra := ""
		for i := start; i < end; i++ {
			// fmt.Println("ADDING:", b[i].GetString())
			extra += b[i].GetString()
		}
		resp += "++" + extra
	}
	return resp
}
