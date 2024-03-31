// data manipulation functions

package erdata

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
)

// upgrade er route to v2
func upgradeErRoute(route ErRoute,itemInfo ItemInfoDict) ErRoute2 {
    var weapons []ItemInfo

    for i := range route.WeaponIds {
        var foundWeapon ItemInfo
        var found bool
        foundWeapon,found=itemInfo[fmt.Sprintf("%d",route.WeaponIds[i])]


        if !found {
            fmt.Println("failed to find item in item dict")
            fmt.Println("missing item:",route.WeaponIds[i])
            panic("missing item")
        }

        weapons=append(weapons,foundWeapon)
    }

    return ErRoute2 {
        ErRoute: route,
        WeaponInfos: weapons,
    }
}

// upgrade list of er routes
func upgradeErRouteList(routes []ErRoute,itemInfo ItemInfoDict) []ErRoute2 {
    var upgradedList []ErRoute2

    for i := range routes {
        upgradedList=append(upgradedList,upgradeErRoute(routes[i],itemInfo))
    }

    return upgradedList
}

// simplify er route response to just the routes
func extractErRoutes(routeResponse ErRouteResponse) []ErRoute2 {
    return upgradeErRouteList(
        routeResponse.RecommendWeaponRouteDtoPage.Items,
        routeResponse.ItemById,
    )
}

// remove all routes with duplicate ids from list of routes
func purgeDuplicateRoutes(routes []ErRoute2) []ErRoute2 {
    var seenRouteIds sets.Set[int]=sets.New[int]()

    var filteredRoutes []ErRoute2

    for i := range routes {
        // if seen it, skip
        if seenRouteIds.Has(routes[i].Id) {
            fmt.Println("found duplicate route:",routes[i].Id)
            continue
        }

        // otherwise, collect it
        seenRouteIds.Insert(routes[i].Id)
        filteredRoutes=append(filteredRoutes,routes[i])
    }

    return filteredRoutes
}