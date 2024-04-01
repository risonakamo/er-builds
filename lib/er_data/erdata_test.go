// erdata lib tests

package erdata

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp/v3"
	"github.com/kr/pretty"
)

// test getting data and printing it
func Test_getData(t *testing.T) {
    var data ErRouteResponse=getRouteData("Elena","Rapier",1)

    pp.Print(data)
}

// test multi page get
func Test_multiGet(t *testing.T) {
    var data1 []ErRoute2=getRouteDataMultiPage("Elena","Rapier",1,1)
    var data2 []ErRoute2=getRouteDataMultiPage("Elena","Rapier",1,5)

    fmt.Println(len(data1))
    fmt.Println(len(data2))

    pp.Print(data2)
}

// test going through item pipeline
func Test_itemStatPipeline(t *testing.T) {
    var data []ErRoute2=getRouteDataMultiPage("Elena","Rapier",1,10)

    var filtered []ErRoute2=filterByVersion(
        data,
        []string{
            "1.17.0",
            "1.18.0",
        },
    )

    // pp.Print(filtered)
    fmt.Println("found routes:",len(filtered))

    var itemStats ItemStatisticsDict=computeItemStatistics(filtered)

    fmt.Println("unique items:",len(itemStats))
    pretty.Println(itemStats)

    var groupedStats GroupedItemStatistics=groupItemStatistics(itemStats)

    fmt.Println("grouped")
    pretty.Println(groupedStats)
}