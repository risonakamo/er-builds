// erdata lib tests

package erdata

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp/v3"
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