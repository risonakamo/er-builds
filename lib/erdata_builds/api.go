// functions to fetch from er api

package erdata_builds

import (
	"fmt"

	"github.com/imroc/req/v3"
	"github.com/rs/zerolog/log"
)

// top level api data fetch function using most filter features.
// retrieves data for char/weapon with specified number of pages and version filtering.
func GetRouteData2(
    character string,
    weapon string,
    pages int,
    versions []string,
    client *req.Client,
) []ErRoute2 {
    var routes []ErRoute2=getRouteDataMultiPage(
        character,
        weapon,
        0,
        pages,
        true,
        client,
    )

    return filterByVersion(
        routes,
        versions,
    )
}

// main fetch function, fetching and transforming into the better looking data type.
// get routes from multiple pages, merging them all into ErRoute2 list.
// if earlystop given, stop when get a result that is 0 length
func getRouteDataMultiPage(
    character string,
    weapon string,
    pageStart int,
    pageEnd int,
    earlyStop bool,
    client *req.Client,
) []ErRoute2 {
    var routes []ErRoute2

    for i := pageStart; i<pageEnd ; i++ {
        log.Info().Msgf("getting page: %d/%d",i+1,pageEnd)
        var newRoutes []ErRoute2=extractErRoutes(getRouteData(character,weapon,i,client))
        log.Info().Msgf("got %d unfiltered routes",len(newRoutes))
        // pp.Print(newRoutes)

        // since even after getting the api it still can contain some non matching weapons
        newRoutes=filterByWeapon(newRoutes,weapon)
        log.Info().Msg(fmt.Sprintf("-> got %d routes",len(newRoutes)))

        if earlyStop && len(newRoutes)==0 {
            log.Info().Msg("got no routes. stopping data retrieval")
            break
        }

        routes=append(
            routes,
            newRoutes...
        )
    }

    return purgeDuplicateRoutes(routes)
}

// fetch routes for a character.
func getRouteData(
    character string,
    weapon string,
    page int,
    client *req.Client,
) ErRouteResponse {
    var e error
    var result ErRouteResponse
    _,e=client.R().
        AddQueryParam("hl","en").
        AddQueryParam("character",character).
        AddQueryParam("weaponType",weapon).
        AddQueryParam("page",fmt.Sprintf("%d",page)).
        SetSuccessResult(&result).
        Get("https://er-node.dakgg.io/api/v0/routes")

    if e!=nil {
        panic(e)
    }

    return result
}