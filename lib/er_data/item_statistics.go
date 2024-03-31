// functions for performing analysis on ErRoute2 lists

package erdata

// filter list of routes to only routes with a specified version
func filterByVersion(routes []ErRoute2,version string) []ErRoute2 {
    var newroutes []ErRoute2

    for i := range routes {
        if routes[i].Version==version {
            newroutes=append(newroutes,routes[i])
        }
    }

    return newroutes
}

// collect item stats from all routes
func computeItemStatistics(routes []ErRoute2) ItemStatisticsDict {
    var collectedStats ItemStatisticsDict=make(ItemStatisticsDict)

    // for all routes
    for i := range routes {
        var route ErRoute2=routes[i]

        // for all items in a single route
        for i2 := range routes[i].ItemInfos {
            var item ItemInfo2=routes[i].ItemInfos[i2]

            var in bool
            _,in=collectedStats[item.Id]

            // initialise if have not seen this item before. use the route's
            // statistics
            if !in {
                collectedStats[item.Id]=&ItemsStatistics {
                    Item: item,

                    Total: 1,
                    AverageWinRate: route.WinRate,
                    HighestWinRate: route.WinRate,
                }
            } else {
                var stats *ItemsStatistics=collectedStats[item.Id]
                stats.Total+=1
                stats.TotalWinRate+=route.WinRate
                stats.AverageWinRate=stats.TotalWinRate/float32(stats.Total)
                stats.HighestWinRate=max(stats.HighestWinRate,route.WinRate)
            }
        }
    }

    return collectedStats
}

// group dict of item statistics into grouped item stats obj
func groupItemStatistics(itemStats ItemStatisticsDict) GroupedItemStatistics {
    var grouped GroupedItemStatistics=make(GroupedItemStatistics)

    var itemStat *ItemsStatistics
    for _,itemStat = range itemStats {
        var itemType ItemType=itemStat.Item.ItemType

        var in bool
        _,in=grouped[itemType]

        // initialise group if not seen the item type yet
        if !in {
            grouped[itemType]=[]ItemsStatistics{}

        // otherwise add to the group
        } else {
            grouped[itemType]=append(grouped[itemType],*itemStat)
        }
    }

    return grouped
}