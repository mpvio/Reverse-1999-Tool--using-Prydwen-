package models

type ItemList struct {
	Data Data `json:"data"`
}

type Data struct {
	Characters Content `json:"allContentfulReverseCharacter,omitzero"`
	Psychubes  Content `json:"allContentfulReversePsychube,omitzero"`
}

type Content struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	// slug and name also used in character file for recommended psychubes
	Slug string `json:"slug"`
	Name string `json:"name"`
	// exclusive to psychubes
	Rarity               string            `json:"rarity,omitempty"`
	Stats                Stats             `json:"stats,omitzero"`
	Description1         DescriptionAsText `json:"descriptionLevel1,omitzero"`
	Description5         DescriptionAsText `json:"descriptionLevel5,omitzero"`
	Tags                 []string          `json:"tags,omitempty,omitzero"`
	NotAvailableInGlobal bool              `json:"notAvailableInGlobal,omitempty"`
}

// exclusive to psychubes
type Stats struct {
	Atk        SingleStat `json:"atk,omitzero"`
	Hp         SingleStat `json:"hp,omitzero"`
	MentalDef  SingleStat `json:"mental_def,omitzero"`
	RealityDef SingleStat `json:"reality_def,omitzero"`
	Custom     SingleStat `json:"custom,omitzero"`
}

type SingleStat struct {
	Base int `json:"base,omitempty,omitzero"`
	Max  int `json:"max,omitempty,omitzero"`
	// only used by "custom"
	Name string `json:"name,omitempty,omitzero"`
}

// for psychubes and full characters
type DescriptionAsText struct {
	Raw string `json:"raw,omitempty"`
}
