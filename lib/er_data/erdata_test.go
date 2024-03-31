package erdata

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp/v3"
)

// test getting data and printing it
func Test_getData(t *testing.T) {
    data:=getRouteData("Elena","Rapier",1)

    pp.Print(data)
}

// test multi page get
func Test_multiGet(t *testing.T) {
    data1:=getRouteDataMultiPage("Elena","Rapier",1,1)
    data2:=getRouteDataMultiPage("Elena","Rapier",1,5)

    fmt.Println(len(data1))
    fmt.Println(len(data2))
}