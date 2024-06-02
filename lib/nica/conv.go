// funcs converting nica data formats

package nica

import (
	"encoding/json"
	"er-builds/lib/dak_gg"
	"er-builds/lib/erdata_builds"
	"er-builds/lib/oer_api"

	"github.com/rs/zerolog/log"
)

// nica build 2s keyed by their id
// key: id of nica build
// val: the nica build
type NicaBuild2Dict map[int]NicaBuild2

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
	// original build id
	Id int

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

// parse nica build contents to create nica build 2.
// requires additional info:
// - list of all possible trait skills from dakgg
// - optional: langfile dict. if not available, skips parsing of late game items
func upgradeNicaBuildTo2(
	build NicaBuild,
	traitSkills dak_gg.TraitSkillMap,
	langDict oer_api.OerLangDict,
) NicaBuild2 {
	var itemInfos []erdata_builds.ItemInfo2

	// converting trait selections to item info2s
	for i := range build.TraitCodes {
		var traitInfo dak_gg.TraitSkill
		var in bool
		traitInfo,in=traitSkills[build.TraitCodes[i]]

		if !in {
			continue
		}

		itemInfos=append(itemInfos,erdata_builds.ItemInfo2{
			ItemInfo: erdata_builds.ItemInfo{
				Id: traitInfo.Id,
				Name: traitInfo.Name,
				Tooltip: traitInfo.Tooltip,
				ImageUrl: traitInfo.ImageUrl,
				BackgroundImageUrl: "",
			},

			ItemType: erdata_builds.ItemType_augment,
			WeaponName: "",
		})
	}

	// converting late game item codes to item info2s
	var itemId int
	for _,itemId = range build.LateGameItemCodes {
		var itemName string
		var e error
		itemName,e=oer_api.GetItemName(langDict,itemId)

		if e!=nil {
			log.Warn().Msgf("failed to find late game item code: %i",itemId)
			continue
		}

		itemInfos=append(itemInfos,erdata_builds.ItemInfo2{
			ItemInfo: erdata_builds.ItemInfo{
				Id: itemId,
				Name: itemName,
				Tooltip: "",
				ImageUrl: dak_gg.CreateItemIconUrl(itemId),
				BackgroundImageUrl: "",
			},

			ItemType: erdata_builds.ItemType_late	,
			WeaponName: "",
		})
	}

	return NicaBuild2{
		Id: build.Id,
		ItemInfos: itemInfos,
	}
}

// parse an int array string
// if input str is empty, returns empty array
func parseIntArrayStr(intarrayStr string) []int {
	if len(intarrayStr)==0 {
		return []int{}
	}

	var result []int
	var e error=json.Unmarshal([]byte(intarrayStr),&result)

	if e!=nil {
		log.Error().Msgf("bad input int array input: %v",intarrayStr)
		panic(e)
	}

	return result
}

// parse a map of int arrays into a merged list of int arrays.
// all the keys of the map will be lost.
// used for the lategame item codes, which is stored as multi level
// int arrays, but we don't care about the levels
// if input str is empty, returns empty array
func parseIntMapStr(intMapStr string) []int {
	if len(intMapStr)==0 {
		return []int{}
	}

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

// convert nica build array into map
func groupNicaBuildsById(builds []NicaBuild2) NicaBuild2Dict {
	var result NicaBuild2Dict=make(NicaBuild2Dict)
	for buildI := range builds {
		result[builds[buildI].Id]=builds[buildI]
	}

	return result
}