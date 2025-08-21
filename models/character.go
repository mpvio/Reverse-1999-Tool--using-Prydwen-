package models

import "encoding/json"

type Character struct {
	Result Result `json:"result"`
}

type Result struct {
	Data CharacterData `json:"data"`
}

type CharacterData struct {
	CurrentUnit CurrentUnit `json:"currentUnit"`
}

type CurrentUnit struct {
	Nodes []CharacterNode `json:"nodes"`
}

type CharacterNode struct {
	// Slug   string `json:"slug"`
	Name               string            `json:"name"`
	Rarity             string            `json:"rarity"`
	Afflatus           string            `json:"afflatus"`
	DamageType         string            `json:"damageType"`
	TierListCategory   string            `json:"tierListCategory"`
	TierListTags       string            `json:"tierListTags"`
	TierEuphoria       string            `json:"tierEuphoria,omitempty"` // optional
	Tags               []string          `json:"tags"`
	AvailableInGlobal  bool              `json:"availableInGlobal"`
	Rating             Rating            `json:"ratingsNew,omitzero"`
	TierComment        TierComment       `json:"tierComment,omitzero"`
	Skills             []Skill           `json:"skills,omitempty"`
	Insights           Inheritance       `json:"inheritance"`
	Portray            Portray           `json:"portray"`
	Pros               DescriptionAsText `json:"pros,omitzero"`
	Cons               DescriptionAsText `json:"cons,omitzero"`
	Materials          Materials         `json:"materials,omitzero"`
	SuggestedPsychubes []Node            `json:"psychubeSuggested,omitempty"`
	PsychubeComments   DescriptionAsText `json:"psychubeComments,omitzero"`
	CharacterStats     CharacterStats    `json:"characterStats"`
	Euphoria           []Euphoria        `json:"euphoria,omitempty"`
	Resonance          []Resonance       `json:"resonance,omitempty"`
}

type Inheritance struct {
	Name   string            `json:"name"`
	Level1 DescriptionAsText `json:"level1"`
	Level2 DescriptionAsText `json:"level2"`
	Level3 DescriptionAsText `json:"level3,omitzero"`
	Status []Status          `json:"status,omitempty,omitzero"`
}

type Status struct {
	Name string            `json:"name"`
	Type string            `json:"type,omitempty"`
	Desc DescriptionAsText `json:"description,omitzero"`
}

type Portray struct {
	Level1 DescriptionAsText `json:"level1"`
	Level2 DescriptionAsText `json:"level2"`
	Level3 DescriptionAsText `json:"level3"`
	Level4 DescriptionAsText `json:"level4"`
	Level5 DescriptionAsText `json:"level5"`
}

type Skill struct {
	Name     string            `json:"name"`
	Category string            `json:"category"`
	Desc1    DescriptionAsText `json:"description1"`
	Desc2    DescriptionAsText `json:"description2,omitzero"`
	Desc3    DescriptionAsText `json:"description3,omitzero"`
	Type1    string            `json:"type1,omitempty,omitzero"`
	Type2    string            `json:"type2,omitempty,omitzero"`
	Type3    string            `json:"type3,omitempty,omitzero"`
	Status   []Status          `json:"status,omitempty,omitzero"`
}

type Rating struct {
	Base int `json:"base,omitempty,omitzero"`
}

type TierComment struct {
	TierComment string `json:"tierComment,omitempty,omitzero"`
}

type Materials struct {
	Insight1 []Material `json:"insight_1"`
	Insight2 []Material `json:"insight_2"`
	Insight3 []Material `json:"insight_3,omitempty"`
}

type Material struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type CharacterStats struct {
	Atk        SingleCharacterStat `json:"atk"`
	Hp         SingleCharacterStat `json:"hp"`
	RealityDef SingleCharacterStat `json:"reality_def"`
	MentalDef  SingleCharacterStat `json:"mental_def"`
	Crit       SingleCharacterStat `json:"crit"`
	CritRate   SingleCharacterStat `json:"crit_rate"`
	CritDMG    SingleCharacterStat `json:"crit_dmg"`
}

type SingleCharacterStat struct {
	Insight0Min json.Number `json:"insight_0_min"`
	Insight1Max json.Number `json:"insight_1_max"`
	Insight2Max json.Number `json:"insight_2_max"`
	Insight3Max json.Number `json:"insight_3_max,omitempty"`
}

type Euphoria struct {
	Name string            `json:"name"`
	Desc DescriptionAsText `json:"description"`
}

type Resonance struct {
	Name string            `json:"name"`
	Code string            `json:"code"`
	Desc DescriptionAsText `json:"description"`
}
