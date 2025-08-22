package models

type CharacterDB struct {
	Name               string         `json:"name"`
	Rarity             string         `json:"rarity"`
	Afflatus           string         `json:"afflatus"`
	DamageType         string         `json:"damageType"`
	TierListCategory   string         `json:"tierListCategory"`
	TierListTags       string         `json:"tierListTags"`
	TierEuphoria       string         `json:"tierEuphoria,omitempty"` // optional
	Tags               []string       `json:"tags"`
	AvailableInGlobal  bool           `json:"availableInGlobal"`
	Rating             int            `json:"ratingsNew,omitempty"`  // converted
	TierComment        string         `json:"tierComment,omitempty"` // converted
	Skills             []SkillDB      `json:"skills,omitempty"`      // converted
	Insights           InsightsDB     `json:"insights"`              // converted
	Portray            PortrayDB      `json:"portray"`               // converted
	Pros               string         `json:"pros,omitempty"`
	Cons               string         `json:"cons,omitempty"`
	Materials          MaterialsDB    `json:"materials,omitzero"` // converted
	SuggestedPsychubes []string       `json:"psychubeSuggested,omitempty,omitzero"`
	PsychubeComments   string         `json:"psychubeComments,omitempty"`
	CharacterStats     CharacterStats `json:"characterStats"`
	Euphoria           []EuphoriaDB   `json:"euphoria,omitempty"`  // converted
	Resonance          []ResonanceDB  `json:"resonance,omitempty"` // converted
}

type StatusDB struct {
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
	Desc string `json:"description,omitempty"`
}

type PortrayDB struct {
	Level1 string `json:"level1"`
	Level2 string `json:"level2"`
	Level3 string `json:"level3"`
	Level4 string `json:"level4"`
	Level5 string `json:"level5"`
}

type SkillDB struct {
	Name     string     `json:"name"`
	Category string     `json:"category"`
	Desc     string     `json:"description"`
	Type     string     `json:"type,omitempty,omitzero"`
	Status   []StatusDB `json:"status,omitempty,omitzero"`
}

type InsightsDB struct {
	Name   string     `json:"name"`
	Level1 string     `json:"level1"`
	Level2 string     `json:"level2"`
	Level3 string     `json:"level3,omitzero"`
	Status []StatusDB `json:"status,omitempty,omitzero"`
}

type EuphoriaDB struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

type ResonanceDB struct {
	Name string `json:"name"`
	Code string `json:"code"`
	Desc string `json:"description"`
}

type MaterialsDB struct {
	Insight1 []Material     `json:"insight_1"`
	Insight2 []Material     `json:"insight_2"`
	Insight3 []Material     `json:"insight_3,omitempty"`
	Total    map[string]int `json:"total"`
}
