package erdata

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

// test getting data and printing it
func Test_getData(t *testing.T) {
    data:=getRouteData("Elena","Rapier",1)

    pp.Print(data)
}