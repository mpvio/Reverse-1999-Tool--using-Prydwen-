package models

import (
	"fmt"
	"golangR99/constants"
)

func (c Character) Convert() CharacterDB {
	return c.Result.Data.CurrentUnit.Nodes[0].Convert()
}

func (c CharacterNode) Convert() CharacterDB {
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
		Skills:             ConvertSlice(c.Skills),
		Insights:           c.Insights.Convert(),
		Portray:            c.Portray.Convert(),
		Pros:               c.Pros.ConvertToJson().GetString(),
		Cons:               c.Cons.ConvertToJson().GetString(),
		Materials:          c.Materials.Convert(),
		SuggestedPsychubes: convertNodesToStrings(c.SuggestedPsychubes),
		PsychubeComments:   c.PsychubeComments.ConvertToJson().GetString(),
		CharacterStats:     c.CharacterStats, // keep the same for now
		Euphoria:           ConvertSlice(c.Euphoria),
		Resonance:          ConvertSlice(c.Resonance),
	}
}

func (i Inheritance) Convert() InsightsDB {
	return InsightsDB{
		Name:   i.Name,
		Level1: i.Level1.ConvertToJson().GetString(),
		Level2: i.Level2.ConvertToJson().GetString(),
		Level3: i.Level3.ConvertToJson().GetString(),
		Status: ConvertSlice(i.Status),
	}
}

func (s Status) Convert() StatusDB {
	desc := s.Desc.ConvertToJson().GetString()
	return StatusDB{
		Name: s.Name,
		Type: s.Type,
		Desc: desc,
	}
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

func (s Skill) Convert() SkillDB {
	return SkillDB{
		Name:     s.Name,
		Category: s.Category,
		Desc:     s.DescToString(),
		Type:     s.TypeToString(),
		Status:   ConvertSlice(s.Status),
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

func (r Rating) Convert() int {
	return r.Base
}

func (t TierComment) Convert() string {
	return t.TierComment
}

// todo: add validator (to Material too) to remove Insight3 if it's just "Material": 0
func (m Materials) Convert() MaterialsDB {
	total := make(map[string]int)
	// only 5 & 6* characters have Insight 3 materials
	insight3valid := materialSliceValidity(m.Insight3)
	insights := [][]Material{m.Insight1, m.Insight2}
	var insight3 []Material
	if insight3valid {
		insights = append(insights, m.Insight3)
		insight3 = m.Insight3
	}
	// add insight costs
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
		Insight3: insight3, // [] if insight3valid == false
		Total:    total,
	}
}

func (m Materials) GetValidInsights() [][]Material {
	var insights [][]Material
	if materialSliceValidity(m.Insight1) {
		insights = append(insights, m.Insight1)
	}
	if materialSliceValidity(m.Insight2) {
		insights = append(insights, m.Insight2)
	}
	if materialSliceValidity(m.Insight3) {
		insights = append(insights, m.Insight3)
	}
	return insights
}

func (m Material) Valid() bool {
	if m.Name == "Material" || m.Amount == 0 {
		return false
	}
	return true
}

func (e Euphoria) Convert() EuphoriaDB {
	return EuphoriaDB{
		Name: e.Name,
		Desc: e.Desc.ConvertToJson().GetString(),
	}
}

func (r Resonance) Convert() ResonanceDB {
	return ResonanceDB{
		Name: r.Name,
		Code: r.Code,
		Desc: r.Desc.ConvertToJson().GetString(),
	}
}

// helper function(s)
func materialSliceValidity(materials []Material) bool {
	for _, m := range materials {
		if !m.Valid() {
			return false
		}
	}
	return true
}

type Convertible[T any] interface {
	Convert() T
}

// "S ~[]E" = take any slice S whose elements are type E
// "E Convertible[T]" = E implements Convertible and returns a type T (constraint)
// "T any" = the aforementioned type T
func ConvertSlice[S ~[]E, E Convertible[T], T any](slice S) []T {
	return MapSlice(slice, func(e E) T { return e.Convert() })
}

// "mapper func(E) T" = function that takes parameter of type E and returns type T
// E has no constraint to enable direct calls to MapSlice (e.g. for types w/o Convert, like Node)
func MapSlice[S ~[]E, E any, T any](slice S, mapper func(E) T) []T {
	if slice == nil {
		return nil
	}
	result := make([]T, len(slice))
	for i, val := range slice {
		result[i] = mapper(val)
	}
	return result
}

func convertNodesToStrings(nodes []Node) []string {
	return MapSlice(nodes, func(n Node) string { return n.Name })
}
