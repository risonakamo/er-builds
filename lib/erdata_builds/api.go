// functions to fetch from er api

package erdata_builds

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rs/zerolog/log"
)

// top level api data fetch function using most filter features.
// retrieves data for char/weapon with specified number of pages and version filtering.
func GetRouteData2(
    character string,
    weapon string,
    pages int,
    versions []string,
) []ErRoute2 {
    var routes []ErRoute2=getRouteDataMultiPage(
        character,
        weapon,
        0,
        pages,
        true,
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
) []ErRoute2 {
    var routes []ErRoute2

    for i := pageStart; i<pageEnd ; i++ {
        log.Info().Msgf("getting page: %d/%d",i+1,pageEnd)
        var newRoutes []ErRoute2=extractErRoutes(getRouteData(character,weapon,i))
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
) ErRouteResponse {
    var e error

    // create the request
    var req *http.Request
    req,e=http.NewRequest(
        http.MethodGet,
        "https://er-node.dakgg.io/api/v0/routes",
        nil,
    )

    if e!=nil {
        panic(e)
    }

    // fill in url queries
    var query url.Values=req.URL.Query()
    query.Add("hl","en")
    query.Add("character",character)
    query.Add("weaponType",weapon)
    query.Add("page",fmt.Sprintf("%d",page))

    req.URL.RawQuery=query.Encode()
    log.Info().Msgf("making url request: %s",req.URL)

    // make the request
    var client http.Client=http.Client{}

    var resp *http.Response
    resp,e=client.Do(req)

    if e!=nil {
        panic(e)
    }

    defer resp.Body.Close()

    // parse request into obj
    var data []byte
    data,e=io.ReadAll(resp.Body)

    if e!=nil {
        panic(e)
    }

    var routeObj ErRouteResponse
    json.Unmarshal(data,&routeObj)

    return routeObj
}