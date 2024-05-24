// data manipulation functions

package erdata_builds

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

    var upgradedItemInfos []ItemInfo2=upgradeAllItems(weapons)

    return ErRoute2 {
        ErRoute: route,
        ItemInfos: upgradedItemInfos,
        MainWeapon: getMainWeapon(upgradedItemInfos),
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
// creates new array
func purgeDuplicateRoutes(routes []ErRoute2) []ErRoute2 {
    var seenRouteIds sets.Set[int]=sets.New[int]()

    var filteredRoutes []ErRoute2

    for i := range routes {
        // if seen it, skip
        if seenRouteIds.Has(routes[i].Id) {
            // fmt.Println("found duplicate route:",routes[i].Id)
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
    var itemType ItemType
    var weaponName string
    itemType,weaponName=extractItemType(item)

    return ItemInfo2 {
        ItemInfo: item,
        ItemType: itemType,
        WeaponName: weaponName,
    }
}

// try to determine item type. defaults to weapon.
// if the weapon is "weapon" type, also returns the actual weapon type value. otherwise, empty.
//
// todo: get list of all possible weapons to confirm that something is weapon. right
// now it defaults to weapon if it doesn't recognise the type - might be an issue if
// one of the non weapon types changes
func extractItemType(item ItemInfo) (ItemType,string) {
    // matches something like: "Epic / Chest\n\nDefense +16\nSkill Amplification..."
    // tries to extract the 2nd word which is the item type
    // [0]: whole match
    // [1]: the item type
    var reg *regexp.Regexp=regexp.MustCompile(`\w+ \/ (.*)\n`)
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
            return ItemType_chest,""
        case "Head":
            return ItemType_head,""
        case "Arm","Arm/Accessory":
            return ItemType_arm,""
        case "Leg":
            return ItemType_leg,""
    }

    return ItemType_weapon,extractedType
}

// convert list of items into item info 2
func upgradeAllItems(items []ItemInfo) []ItemInfo2 {
    var newitems []ItemInfo2

    for i := range items {
        newitems=append(newitems,upgradeItemInfo(items[i]))
    }

    return newitems
}

// get the main weapon from list of item infos. main weapon is the first item to
// have weapon filled out. removes all white spaces in the name
func getMainWeapon(items []ItemInfo2) string {
    for i := range items {
        if len(items[i].WeaponName)>0 {
            var convertedName string
            var in bool
            convertedName,in=WeaponNameToShortName[items[i].WeaponName]

            if !in {
                convertedName=items[i].WeaponName
            }

            return convertedName
        }
    }

    panic("could not find main weapon")
}

// filter routes down to only ones with certain weapon. for cleaning up the api call
func filterByWeapon(routes []ErRoute2,weapon string) []ErRoute2 {
    var result []ErRoute2

    for i := range routes {
        // pp.Print(routes[i])
        if routes[i].MainWeapon==weapon {
            result=append(result,routes[i])
        }
    }

    return result
}