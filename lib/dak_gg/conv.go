// funcs for converting dak gg data types to other forms

package dak_gg

import "strings"

// keyed version of trait skill list.
// key: trait skill id
// val: the trait skill
type TraitSkillMap map[int]TraitSkill

// key list of trait skills by each one's id
func GroupTraitSkillsById(traitSkills []TraitSkill) TraitSkillMap {
	var result TraitSkillMap=make(TraitSkillMap)

	for i := range traitSkills {
		result[traitSkills[i].Id]=traitSkills[i]
	}

	return result
}

// fix image url fields in trait skills to formalise with other urls. removes the cdn.dak.gg
// in the front.
// \/\/cdn.dak.gg/assets/er/game-assets/1.22.0/TraitSkillIcon_31200.png
// becomes
// /er/game-assets/1.22.0/TraitSkillIcon_31200.png
// MUTATES the input array
func fixTraitSkillsUrls(traitSkills []TraitSkill) []TraitSkill {
	for i := range traitSkills {
		traitSkills[i].ImageUrl=strings.TrimPrefix(
			traitSkills[i].ImageUrl,
			"//cdn.dak.gg/assets",
		)
	}

	return traitSkills
}