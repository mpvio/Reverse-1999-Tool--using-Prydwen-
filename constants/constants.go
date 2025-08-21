package constants

import (
	"fmt"
	"strings"
)

const (
	// endpoints
	CHARACTERS_ENDPOINT = "https://www.prydwen.gg/page-data/sq/d/4273101454.json"
	PSYCHUBES_ENDPOINT  = "https://www.prydwen.gg/page-data/sq/d/2760202169.json"
	// node types
	DOCUMENT  = "document"
	PARAGRAPH = "paragraph"
	TEXT      = "text"
	// formatting
	ITALIC = "italic"
	BOLD   = "bold"
	// item types
	CHARACTER = "character"
	PSYCHUBE  = "psychube"
)

func GET_CHARACTER_ENDPOINT(name string) string {
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")
	return fmt.Sprintf("https://www.prydwen.gg/page-data/re1999/characters/%s/page-data.json", name)
}
