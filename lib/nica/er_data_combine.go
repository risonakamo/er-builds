// functions to combine nica data with erdata_builds ErRoute2

package nica

import "er-builds/lib/erdata_builds"

// combines nica data with routes. adds all items from nica data to the routes.
// MUTATES the original routes array.
func CombineWithErRoutes(
    routes []erdata_builds.ErRoute2,
    nicaData []NicaBuild2,
) []erdata_builds.ErRoute2 {
    var nicaDataDict NicaBuild2Dict=groupNicaBuildsById(nicaData)

    // for each route, if it has a corresponding nica build, combine that nica build's
    // item infos into the route
    for routeI := range routes {
        var foundNicaBuild NicaBuild2
        var in bool
        foundNicaBuild,in=nicaDataDict[routes[routeI].Id]

        if !in {
            continue
        }

        routes[routeI].ItemInfos=append(
            routes[routeI].ItemInfos,
            foundNicaBuild.ItemInfos...,
        )
    }

    return routes
}