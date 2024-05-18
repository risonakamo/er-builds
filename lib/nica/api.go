// functions for accessing nica api

package nica

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

// cleaned up version of a raw build response
type NicaBuild struct {
    Id int
    WeaponCodes []int
    TraitCodes []int
    LateGameItemCodes []int
    TacticalSkillGroupCode int
    Paths []int
}

// retrieve target build id from nica api
func getBuild(buildId int) BuildResponseRaw {
    var reqUrl string=fmt.Sprintf(
        "https://api.nicashow.fun/build/%d",
        buildId,
    )

    var e error
    var req *http.Request
    req,e=http.NewRequest(
        http.MethodGet,
        reqUrl,
        nil,
    )

    if e!=nil {
        panic(e)
    }

    var client http.Client=http.Client{}

    var resp *http.Response
    resp,e=client.Do(req)

    if e!=nil {
        panic(e)
    }

    defer resp.Body.Close()

    var data []byte
    data,e=io.ReadAll(resp.Body)

    if e!=nil {
        panic(e)
    }

    var parsedData BuildResponseRaw
    e=json.Unmarshal(data,&parsedData)

    if e!=nil {
        panic(e)
    }

    return parsedData
}