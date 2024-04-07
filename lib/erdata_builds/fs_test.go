// testing fs functions

package erdata_builds

import (
	"fmt"
	"testing"
)

// retrieve data and try to write it to file.
// should work if the file doesn't already exist, and after it
// does exist (running it more than once),
// the other times should update the file by a small amount
func Test_writeTest(t *testing.T) {
    var data []ErRoute2=getRouteDataMultiPage("Elena","Rapier",1,10)

    var filtered []ErRoute2=filterByVersion(
        data,
        []string{
            "1.17.0",
            "1.18.0",
        },
    )

    fmt.Println("found routes:",len(filtered))

    MergeDataIntoFile(filtered,"testdata/elena.json")
}