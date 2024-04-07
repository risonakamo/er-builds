// functions to fetch from er api

package erdata_builds

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// fetch routes for a character
func getRouteData(character string,weapon string,page int) ErRouteResponse {
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

    var routeObj ErRouteResponse
    json.Unmarshal(data,&routeObj)

    return routeObj
}

// main fetch function, fetching and transforming into the better looking data type.
// get routes from multiple pages, merging them all into ErRoute2 list
func getRouteDataMultiPage(
    character string,
    weapon string,
    pageStart int,
    pageEnd int,
) []ErRoute2 {
    var routes []ErRoute2

    for i := pageStart; i<=pageEnd ; i++ {
        routes=append(
            routes,
            extractErRoutes(getRouteData(character,weapon,i))...
        )
    }

    return purgeDuplicateRoutes(routes)
}