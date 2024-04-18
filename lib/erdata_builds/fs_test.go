// testing fs functions

package erdata_builds

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp/v3"
)

// retrieve data and try to write it to file.
// should work if the file doesn't already exist, and after it
// does exist (running it more than once),
// the other times should update the file by a small amount
func Test_writeTest(t *testing.T) {
    // var character string="Elena"
    // var weapon string="Rapier"
    var character string="Tia"
    var weapon string="Bat"

    var data []ErRoute2=getRouteDataMultiPage(
        character,
        weapon,
        1,
        10,
        true,
    )

    var filtered []ErRoute2=filterByVersion(
        data,
        []string{
            "1.17.0",
            "1.18.0",
        },
    )

    fmt.Println("found routes:",len(filtered))

    var filename string=GetRouteDataFileName(character,weapon,"testdata")

    MergeDataIntoFile(filtered,filename)
}

// test scanning the testdata dir for datafiles. only works if had written some
func Test_dataDirScan(t *testing.T) {
    var datafiles []ErDataFileDescriptor=GetErDataFiles("testdata")

    pp.Print(datafiles)
}