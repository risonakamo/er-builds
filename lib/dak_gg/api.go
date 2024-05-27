// func implementing dak api access

package dak_gg

import "github.com/imroc/req/v3"

// raw response from calling trait skills info api
type TraitSkillsInfoRaw struct {
	TraitSkills []TraitSkill `json:"traitSkills"`
}

// a single trait skill (augment)
type TraitSkill struct {
	Id int
	Name string
	Tooltip string
	Group string
	Type string
	ImageUrl string
	Active bool
}

// variant of get trait skills info. converts into map
func GetTraitSkillsInfoMap() TraitSkillMap {
	return GroupTraitSkillsById(GetTraitSkillsInfo())
}

// get all trait skill infos from dak as list of trait skills
func GetTraitSkillsInfo() []TraitSkill {
	var client *req.Client=req.C()

	var result TraitSkillsInfoRaw
	var e error
	_,e=client.R().
		SetSuccessResult(&result).
		Get("https://er-node.dakgg.io/api/v1/data/trait-skills?hl=en")

	if e!=nil {
		panic(e)
	}

	return fixTraitSkillsUrls(result.TraitSkills)
}