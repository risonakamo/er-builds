// er builds data download program.
// for certain settings:
// - target character
// - target weapon
// - some set of versions
// - some number of pages
// retrieve the data from api and merge with the current data set.
// how to use: while in this dir, `go run <this file>`

package main

import (
	"er-builds/lib/erdata_builds"
	go_utils "er-builds/lib/utils"
	"fmt"
	"path/filepath"
)

func main() {
    // --- config ---
    var character string="Tia"
    var weapon string="Bat"
    var pages int=50
    var versions []string=[]string{
        "1.19.0",
        "1.18.0",
        "1.20.0",
    }
    // --- end config ---

    fmt.Println("character:",character)
    fmt.Println("weapon:",weapon)
    fmt.Println("versions:",versions)
    fmt.Println()


    var here string=go_utils.GetHereDir()

    var datadir string=filepath.Join(here,"../../data")

    var datafile string=erdata_builds.GetRouteDataFileName(
        character,
        weapon,
        datadir,
    )

    fmt.Println("will write to data file:",datafile)


    fmt.Println("getting data from api...")
    // retrieve new data for the char/weapon
    var newRoutes []erdata_builds.ErRoute2=erdata_builds.GetRouteData2(
        character,
        weapon,
        pages,
        versions,
    )

    fmt.Println("writing data")
    // merge and save into the datafile
    erdata_builds.MergeDataIntoFile(newRoutes,datafile)
}