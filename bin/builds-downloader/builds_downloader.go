// er builds data download program.
// for certain settings:
// - target character
// - target weapon
// - some set of versions
// - some number of pages
// retrieve the data from api and merge with the current data set

package main

import (
	"er-builds/lib/cli"
	"er-builds/lib/erdata_builds"
	go_utils "er-builds/lib/utils"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/rs/zerolog"
)

const Pages int=80

func main() {
    go_utils.ConfigureDefaultZeroLogger()
    zerolog.SetGlobalLevel(zerolog.WarnLevel)
    var here string=go_utils.GetHereDirExe()

    var args cli.DataDownloaderArgs=cli.GetDataDownloaderCliArgs(here,"config")

    fmt.Println("versions:",args.Versions)
    fmt.Println()

    for i := range args.Selections {
        var character string=args.Selections[i].Character
        var weapon string=args.Selections[i].Weapon

        fmt.Printf(
            "character: %s\n",
            color.YellowString(character),
        )
        fmt.Printf("weapon: %s\n",color.BlueString(weapon))
        fmt.Println()

        var datadir string=filepath.Join(here,"data")

        var datafile string=erdata_builds.GetRouteDataFileName(
            character,
            weapon,
            datadir,
        )

        fmt.Printf("will write to data file: %s\n",color.YellowString(datafile))


        fmt.Println("getting data from api...")
        // retrieve new data for the char/weapon
        var newRoutes []erdata_builds.ErRoute2=erdata_builds.GetRouteData2Mt(
            character,
            weapon,
            Pages,
            args.Versions,

            5,
            3,
        )

        fmt.Printf("got %s routes\n",color.YellowString(strconv.Itoa(len(newRoutes))))

        // merge and save into the datafile
        erdata_builds.MergeDataIntoFile(newRoutes,datafile)
        fmt.Println()
    }

    fmt.Println("completed")
    fmt.Println("press ENTER to continue")
    go_utils.WaitForEnterKey()
}