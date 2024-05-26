// funcs converting nica data formats

package nica

import "encoding/json"

// cleaned up version of a raw build response
type NicaBuild struct {
    Id int
    WeaponCodes []int
    TraitCodes []int
    LateGameItemCodes []int
    TacticalSkillGroupCode int
    Paths []int
}

// convert raw build to cleaned up build
func convRawToNicaBuild(rawBuild BuildResponseRaw) NicaBuild {
	return NicaBuild{
		Id: rawBuild.RecommendWeaponRoute.Id,
		WeaponCodes: parseIntArrayStr(rawBuild.RecommendWeaponRoute.WeaponCodes),
		TraitCodes: parseIntArrayStr(rawBuild.RecommendWeaponRoute.TraitCodes),
		LateGameItemCodes: parseIntArrayStr(rawBuild.RecommendWeaponRoute.LateGameItemCodes),
		TacticalSkillGroupCode: rawBuild.RecommendWeaponRoute.TacticalSkillGroupCode,
		Paths: parseIntArrayStr(rawBuild.RecommendWeaponRoute.Paths),
	}
}

// parse an int array string
func parseIntArrayStr(intarrayStr string) []int {
	var result []int
	var e error=json.Unmarshal([]byte(intarrayStr),&result)

	if e!=nil {
		panic(e)
	}

	return result
}