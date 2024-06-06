// functions for performing analysis on ErRoute2 lists

package erdata_builds

import (
	"fmt"

	"k8s.io/apimachinery/pkg/util/sets"
)

// top level item statistics computation func. computes item statistics and groups
// by the type.
func ComputeAllItemStatistics(routes []ErRoute2) GroupedItemStatistics {
    var itemstats ItemStatisticsDict=computeItemStatistics(routes)
    return groupItemStatistics(itemstats)
}

// filter list of routes to only routes with any of the specified versions
func filterByVersion(routes []ErRoute2,versions []string) []ErRoute2 {
    var versionsSet sets.Set[string]=sets.New[string](versions...)
    var newroutes []ErRoute2

    for i := range routes {
        if versionsSet.Has(routes[i].Version) {
            newroutes=append(newroutes,routes[i])
        }
    }

    return newroutes
}

// collect item stats from all routes
func computeItemStatistics(routes []ErRoute2) ItemStatisticsDict {
    var collectedStats ItemStatisticsDict=make(ItemStatisticsDict)

    // for all routes
    for routeI := range routes {
        var route ErRoute2=routes[routeI]

        var seenItems sets.Set[string]=sets.New[string]()
        for itemInfoI := range routes[routeI].ItemInfos {
            var item ItemInfo2=routes[routeI].ItemInfos[itemInfoI]

            var itemId string=createItemStatisticsDictId(item.Id,item.ItemType)

            if seenItems.Has(itemId) {
                continue
            }

            seenItems.Insert(itemId)

            var in bool
            _,in=collectedStats[itemId]

            // initialise if have not seen this item before. use the route's
            // statistics
            if !in {
                collectedStats[itemId]=&ItemsStatistics {
                    Item: item,

                    Total: 0,
                    Likes: 0,
                    BuildsPercentage: 0,

                    BuildLikeRatio: 0,

                    TotalWinRate: 0,
                    AverageWinRate: 0,
                    HighestWinRate: 0,
                }
            }

            var stats *ItemsStatistics=collectedStats[itemId]
            stats.Total+=1
            stats.Likes+=route.Likes
            stats.BuildsPercentage=(float32(stats.Total)/float32(len(routes)))*100
            stats.BuildLikeRatio=float32(stats.Likes)/float32(stats.Total)
            stats.TotalWinRate+=route.WinRate
            stats.AverageWinRate=stats.TotalWinRate/float32(stats.Total)
            stats.HighestWinRate=max(stats.HighestWinRate,route.WinRate)
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
        }

        // always add the item to the group
        grouped[itemType]=append(grouped[itemType],*itemStat)
    }

    return grouped
}

// create id to index into item statistics dict
func createItemStatisticsDictId(itemId int,itemType ItemType) string {
    return fmt.Sprintf("%d_%s",itemId,itemType)
}