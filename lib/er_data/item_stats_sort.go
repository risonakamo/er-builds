// functions for sorting item statistics lists

package erdata

import "sort"

// sorts list of item stats by their total count. mutates.
func sortByTotalBuilds(itemStats []ItemsStatistics) {
    sort.Slice(
        itemStats,
        func(i int, j int) bool {
            return itemStats[i].Total<itemStats[j].Total
        },
    )
}