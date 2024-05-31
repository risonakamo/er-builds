// functions for accessing nica api

package nica

import (
	"strconv"

	"github.com/imroc/req/v3"
)

// raw response from nica api
type BuildResponseRaw struct {
    RecommendWeaponRoute struct {
        Id int
        WeaponCodes string
        TraitCodes string
        LateGameItemCodes string
        TacticalSkillGroupCode int
        Paths string
    }
}

// retrieve target build id from nica api. v2 of func
func getBuild(buildId int,client *req.Client) BuildResponseRaw {
    var result BuildResponseRaw
    var e error
    _,e=client.R().
        SetPathParam("buildId",strconv.Itoa(buildId)).
        SetSuccessResult(&result).
        Get("https://api.nicashow.fun/build/{buildId}")

    if e!=nil {
        panic(e)
    }

    return result
}