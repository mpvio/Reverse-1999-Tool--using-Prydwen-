package models

import (
	"encoding/json"
	"fmt"
	"golangR99/constants"
)

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
	Rating             Rating            `json:"ratingsNew,omitzero"`  // converted
	TierComment        TierComment       `json:"tierComment,omitzero"` // converted
	Skills             []Skill           `json:"skills,omitempty"`     // converted
	Insights           Inheritance       `json:"inheritance"`          // converted
	Portray            Portray           `json:"portray"`              // converted
	Pros               DescriptionAsText `json:"pros,omitzero"`
	Cons               DescriptionAsText `json:"cons,omitzero"`
	Materials          Materials         `json:"materials,omitzero"` // converted
	SuggestedPsychubes []Node            `json:"psychubeSuggested,omitempty"`
	PsychubeComments   DescriptionAsText `json:"psychubeComments,omitzero"`
	CharacterStats     CharacterStats    `json:"characterStats"`
	Euphoria           []Euphoria        `json:"euphoria,omitempty"`  // converted
	Resonance          []Resonance       `json:"resonance,omitempty"` // converted
}

func (c CharacterNode) Convert() CharacterDB {
	// return
	return CharacterDB{
		Name:               c.Name,
		Rarity:             c.Rarity,
		Afflatus:           c.Afflatus,
		DamageType:         c.DamageType,
		TierListCategory:   c.TierListCategory,
		TierListTags:       c.TierListTags,
		TierEuphoria:       c.TierEuphoria,
		Tags:               c.Tags,
		AvailableInGlobal:  c.AvailableInGlobal,
		Rating:             c.Rating.Base,
		TierComment:        c.TierComment.TierComment,
		Skills:             convertSkillArr(c.Skills),
		Insights:           c.Insights.Convert(),
		Portray:            c.Portray.Convert(),
		Pros:               c.Pros.ConvertToJson().GetString(),
		Cons:               c.Cons.ConvertToJson().GetString(),
		Materials:          c.Materials.Convert(),
		SuggestedPsychubes: convertNodesToStrings(c.SuggestedPsychubes),
		PsychubeComments:   c.PsychubeComments.ConvertToJson().GetString(),
		CharacterStats:     c.CharacterStats, // keep the same for now
		Euphoria:           convertEuphoriaArr(c.Euphoria),
		Resonance:          convertResonanceArr(c.Resonance),
	}
}

type Inheritance struct {
	Name   string            `json:"name"`
	Level1 DescriptionAsText `json:"level1"`
	Level2 DescriptionAsText `json:"level2"`
	Level3 DescriptionAsText `json:"level3,omitzero"`
	Status []Status          `json:"status,omitempty,omitzero"`
}

func (i Inheritance) Convert() InsightsDB {
	return InsightsDB{
		Name:   i.Name,
		Level1: i.Level1.ConvertToJson().GetString(),
		Level2: i.Level2.ConvertToJson().GetString(),
		Level3: i.Level3.ConvertToJson().GetString(),
		Status: convertStatusArr(i.Status),
	}
}

type Status struct {
	Name string            `json:"name"`
	Type string            `json:"type,omitempty"`
	Desc DescriptionAsText `json:"description,omitzero"`
}

func (s Status) Convert() StatusDB {
	desc := s.Desc.ConvertToJson().GetString()
	return StatusDB{
		Name: s.Name,
		Type: s.Type,
		Desc: desc,
	}
}

type Portray struct {
	Level1 DescriptionAsText `json:"level1"`
	Level2 DescriptionAsText `json:"level2"`
	Level3 DescriptionAsText `json:"level3"`
	Level4 DescriptionAsText `json:"level4"`
	Level5 DescriptionAsText `json:"level5"`
}

func (p Portray) Convert() PortrayDB {
	return PortrayDB{
		Level1: p.Level1.ConvertToJson().GetString(),
		Level2: p.Level2.ConvertToJson().GetString(),
		Level3: p.Level3.ConvertToJson().GetString(),
		Level4: p.Level4.ConvertToJson().GetString(),
		Level5: p.Level5.ConvertToJson().GetString(),
	}
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

func (s Skill) Convert() SkillDB {
	return SkillDB{
		Name:     s.Name,
		Category: s.Category,
		Desc:     s.DescToString(),
		Type:     s.TypeToString(),
		Status:   convertStatusArr(s.Status),
	}
}

func (s Skill) DescToString() string {
	if s.Category == constants.ULTIMATE {
		return s.Desc1.ConvertToJson().GetString()
	} else {
		sd1, sd2, sd3 := s.Desc1.ConvertToJson(), s.Desc2.ConvertToJson(), s.Desc3.ConvertToJson()
		return sd1.Get3Diff(sd2, sd3)
	}
}
func (s Skill) TypeToString() string {
	s12 := s.Type1 == s.Type2
	s23 := s.Type2 == s.Type3
	if s12 && s23 {
		return s.Type1
	}
	if s12 && !s23 {
		return fmt.Sprintf("[1 -> 3]{%s} -> {%s}", s.Type1, s.Type3)
	}
	if !s12 && s23 {
		return fmt.Sprintf("[1 -> 2]{%s} -> {%s}", s.Type1, s.Type2)
	} else {
		return fmt.Sprintf("[1 -> 2 -> 3]{%s} -> {%s} -> {%s}", s.Type1, s.Type2, s.Type3)
	}
}

type Rating struct {
	Base int `json:"base,omitempty,omitzero"`
}

func (r Rating) Convert() int {
	return r.Base
}

type TierComment struct {
	TierComment string `json:"tierComment,omitempty,omitzero"`
}

func (t TierComment) Convert() string {
	return t.TierComment
}

type Materials struct {
	Insight1 []Material `json:"insight_1"`
	Insight2 []Material `json:"insight_2"`
	Insight3 []Material `json:"insight_3,omitempty"`
}

func (m Materials) Convert() MaterialsDB {
	total := make(map[string]int)
	insights := [][]Material{m.Insight1, m.Insight2, m.Insight3}
	for _, insight := range insights {
		for _, v := range insight {
			_, ok := total[v.Name]
			if ok {
				// add to already saved value
				total[v.Name] += v.Amount
			} else {
				total[v.Name] = v.Amount
			}
		}
	}
	return MaterialsDB{
		Insight1: m.Insight1,
		Insight2: m.Insight2,
		Insight3: m.Insight3,
		Total:    total,
	}
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

func (e Euphoria) Convert() EuphoriaDB {
	return EuphoriaDB{
		Name: e.Name,
		Desc: e.Desc.ConvertToJson().GetString(),
	}
}

type Resonance struct {
	Name string            `json:"name"`
	Code string            `json:"code"`
	Desc DescriptionAsText `json:"description"`
}

func (r Resonance) Convert() ResonanceDB {
	return ResonanceDB{
		Name: r.Name,
		Code: r.Code,
		Desc: r.Desc.ConvertToJson().GetString(),
	}
}

// helper function(s)
func convertStatusArr(statuses []Status) []StatusDB {
	result := make([]StatusDB, len(statuses))
	for i, val := range statuses {
		result[i] = val.Convert()
	}
	return result
}

func convertSkillArr(skills []Skill) []SkillDB {
	result := make([]SkillDB, len(skills))
	for i, val := range skills {
		result[i] = val.Convert()
	}
	return result
}

func convertEuphoriaArr(euphorias []Euphoria) []EuphoriaDB {
	result := make([]EuphoriaDB, len(euphorias))
	for i, val := range euphorias {
		result[i] = val.Convert()
	}
	return result
}

func convertResonanceArr(resonances []Resonance) []ResonanceDB {
	result := make([]ResonanceDB, len(resonances))
	for i, val := range resonances {
		result[i] = val.Convert()
	}
	return result
}

func convertNodesToStrings(nodes []Node) []string {
	result := make([]string, len(nodes))
	for i, val := range nodes {
		result[i] = val.Name
	}
	return result
}
