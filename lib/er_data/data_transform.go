// data manipulation functions

package erdata

import "fmt"

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