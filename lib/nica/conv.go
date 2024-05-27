// funcs converting nica data formats

package nica

import (
	"encoding/json"
	"er-builds/lib/erdata_builds"
)

// cleaned up version of a raw build response
type NicaBuild struct {
    Id int
	// purple weapons
    WeaponCodes []int
	// augments. size 6, includes both main and sub augments
    TraitCodes []int
	// all items chosen in lategame items, might not all be weapons/armour
    LateGameItemCodes []int

	// not sure what this is yet. seems too short to be tac skill code
    TacticalSkillGroupCode int

	// probably the route locations
    Paths []int
}

// upgraded form of nica build
type NicaBuild2 struct {
	NicaBuild

	// converting available information in the nica build to list of
	// item infos
	ItemInfos []erdata_builds.ItemInfo2
}

// convert raw build to cleaned up build
func convRawToNicaBuild(rawBuild BuildResponseRaw) NicaBuild {
	return NicaBuild{
		Id: rawBuild.RecommendWeaponRoute.Id,
		WeaponCodes: parseIntArrayStr(rawBuild.RecommendWeaponRoute.WeaponCodes),
		TraitCodes: parseIntArrayStr(rawBuild.RecommendWeaponRoute.TraitCodes),
		LateGameItemCodes: parseIntMapStr(rawBuild.RecommendWeaponRoute.LateGameItemCodes),
		TacticalSkillGroupCode: rawBuild.RecommendWeaponRoute.TacticalSkillGroupCode,
		Paths: parseIntArrayStr(rawBuild.RecommendWeaponRoute.Paths),
	}
}

// parse nica build contents to create nica build 2
func upgradeNicaBuildTo2(build NicaBuild) NicaBuild2 {

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

// parse a map of int arrays into a merged list of int arrays.
// all the keys of the map will be lost.
// used for the lategame item codes, which is stored as multi level
// int arrays, but we don't care about the levels
func parseIntMapStr(intMapStr string) []int {
	var mapResult map[string][]int
	var e error=json.Unmarshal([]byte(intMapStr),&mapResult)

	if e!=nil {
		panic(e)
	}

	var collected []int
	for i := range mapResult {
		collected=append(collected,mapResult[i]...)
	}

	return collected
}