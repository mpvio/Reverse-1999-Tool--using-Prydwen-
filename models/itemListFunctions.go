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

func (n Node) ConvertToPsychube() PsychubeDB {
	desc1 := n.Description1.ConvertToJson()
	desc5 := n.Description5.ConvertToJson()
	return PsychubeDB{
		Name:              n.Name,
		Rarity:            n.Rarity,
		AvailableInGlobal: !n.NotAvailableInGlobal,
		Stats:             n.Stats,
		Desc:              desc1.GetDiff(desc5),
		Tags:              n.Tags,
	}
}

func (dt DescriptionAsText) ConvertToJson() Raw {
	var desc Raw
	if dt.Raw == "" {
		return Raw{}
	}
	err := json.Unmarshal([]byte(dt.Raw), &desc)
	if err != nil {
		log.Fatal(err)
	}
	return desc
}
