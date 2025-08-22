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

func (r Raw) Get3Diff(b, c Raw) string {
	return content3Diff(r.Content[0].Content, b.Content[0].Content, c.Content[0].Content)
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

func (ci ContentItem) Get3Diff(b, c ContentItem) string {
	ab, _ := ci.Equals(b)
	bc, _ := b.Equals(c)
	// a == b == c
	if ab && bc {
		return ci.GetString()
	}
	// a = b, b != c
	if ab && !bc {
		if ci.HasValue() && c.HasValue() {
			return fmt.Sprintf("[1 -> 3]{%s} -> {%s}", ci.GetString(), c.GetString())
		} else {
			return fmt.Sprintf("[1 -> 3]{%s}", contentArrDiff(ci.Content, c.Content))
		}
	}
	// a != b, b = c
	if !ab && bc {
		if ci.HasValue() && b.HasValue() {
			return fmt.Sprintf("[1 -> 2]{%s} -> {%s}", ci.GetString(), b.GetString())
		} else {
			return fmt.Sprintf("[1 -> 2]{%s}", contentArrDiff(ci.Content, b.Content))
		}
	}
	// a != b != c
	aValue, bValue, cValue := ci.HasValue(), b.HasValue(), c.HasValue()
	var aStr, bStr, cStr string

	if aValue {
		aStr = ci.GetString()
	}
	if bValue {
		bStr = b.GetString()
	}
	if cValue {
		cStr = c.GetString()
	}

	// if all have values, just get their strings
	if aValue && bValue && cValue {
		return fmt.Sprintf("[1 -> 2 -> 3]{%s} -> {%s} -> {%s}", aStr, bStr, cStr)
	}
	// else, use content3Diff to get diffs for contentItems that don't.
	if aValue && bValue && !cValue {
		return fmt.Sprintf("[1 -> 2 -> 3]{%s} -> {%s} {%s}", aStr, bStr, content3Diff(ci.Content, b.Content, c.Content))
	}
	if aValue && !bValue && !cValue {
		return fmt.Sprintf("[1 -> 2 -> 3]{%s} {%s}", aStr, content3Diff(ci.Content, b.Content, c.Content))
	} else {
		return fmt.Sprintf("[1 -> 2 -> 3]{%s}", content3Diff(ci.Content, b.Content, c.Content))
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

func content3Diff(a, b, c []ContentItem) string {
	resp := ""
	// assuming len a <= len b <= len c
	for i, val := range a {
		resp += val.Get3Diff(b[i], c[i])
	}
	// arrays are different sizes
	if !(len(a) == len(b) && len(b) == len(c)) {
		alen := len(a)
		blen := len(b)
		clen := len(c)
		extra := ""
		for i := alen; i < blen; i++ {
			extra += b[i].GetDiff(c[i])
		}
		if clen > blen {
			resp += "++[2]" + extra
			remainder := ""
			for i := blen; i < clen; i++ {
				remainder += c[i].GetString()
			}
			resp += "++[3]" + remainder
		} else {
			resp += "++[2 & 3]" + extra
		}

	}
	return resp
}
