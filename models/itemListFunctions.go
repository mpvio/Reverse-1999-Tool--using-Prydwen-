package models

import (
	"encoding/json"
	"golangR99/constants"
	"log"
)

func (n Node) GetType() string {
	if n.Stats == (Stats{}) {
		return constants.CHARACTER
	}
	return constants.PSYCHUBE
}

func (dt DescriptionAsText) ConvertToJson() Raw {
	var desc Raw
	err := json.Unmarshal([]byte(dt.Raw), &desc)
	if err != nil {
		log.Fatal(err)
	}
	return desc
}
