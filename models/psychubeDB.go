package models

type PsychubeDB struct {
	Name              string   `json:"name"`
	Rarity            string   `json:"rarity,omitempty"`
	AvailableInGlobal bool     `json:"availableInGlobal,omitempty"`
	Stats             Stats    `json:"stats,omitzero"`
	Desc              string   `json:"description,omitempty"`
	Tags              []string `json:"tags,omitempty,omitzero"`
}
