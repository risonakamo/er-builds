// data manipulation functions

package erdata

import (
	"fmt"
	"regexp"

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
        ItemInfos: upgradeAllItems(weapons),
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

// add additional fields to item info
func upgradeItemInfo(item ItemInfo) ItemInfo2 {
    return ItemInfo2 {
        ItemInfo: item,
        ItemType: extractItemType(item),
    }
}

// try to determine item type. defaults to weapon
// todo: get list of all possible weapons to confirm that something is weapon
func extractItemType(item ItemInfo) ItemType {
    // matches something like: "Epic / Chest\n\nDefense +16\nSkill Amplification..."
    // tries to extract the 2nd word which is the item type
    // [0]: whole match
    // [1]: the item type
    var reg *regexp.Regexp=regexp.MustCompile(`\w+ \/ (\w+)`)
    var matches []string=reg.FindStringSubmatch(item.Tooltip)

    if len(matches)!=2 {
        fmt.Println("regex extract failed")
        fmt.Println("bad string:",item.Tooltip)
        fmt.Println("matches:",matches)
        panic("bad match")
    }

    var extractedType string=matches[1]

    switch extractedType {
        case "Chest":
            return ItemType_chest
        case "Head":
            return ItemType_head
        case "Arm":
            return ItemType_arm
        case "Leg":
            return ItemType_leg
    }

    return ItemType_weapon
}

// convert list of items into item info 2
func upgradeAllItems(items []ItemInfo) []ItemInfo2 {
    var newitems []ItemInfo2

    for i := range items {
        newitems=append(newitems,upgradeItemInfo(items[i]))
    }

    return newitems
}