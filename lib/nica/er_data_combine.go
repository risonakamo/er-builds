// functions to combine nica data with erdata_builds ErRoute2

package nica

import (
	"er-builds/lib/erdata_builds"

	"k8s.io/apimachinery/pkg/util/sets"
)

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

// given list of er routes and list of nica builds, return a list of IDs
// that are present in the routes data, but NOT in the nica data.
func NicaBuildDiff(
    routes []erdata_builds.ErRoute2,
    nicaBuilds []NicaBuild2,
) []int {
    var nicaBuildIds sets.Set[int]=sets.New[int]()

    // collect ids of all existing nica builds
    for nicaBuildI := range nicaBuilds {
        nicaBuildIds.Insert(nicaBuilds[nicaBuildI].Id)
    }

    var routeIds sets.Set[int]=sets.New[int]()

    // collect ids of all routes
    for routeI := range routes {
        routeIds.Insert(routes[routeI].Id)
    }

    // if remove the nica ids from the route ids, then what we have left is route ids
    // that do not have a corresponding nica id
    return routeIds.Difference(nicaBuildIds).UnsortedList()
}