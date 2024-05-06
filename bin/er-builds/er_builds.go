// ER builds api web server

package main

import (
	"er-builds/lib/erdata_builds"
	go_utils "er-builds/lib/utils"
	"fmt"
	"path/filepath"

	"github.com/gofiber/fiber/v3"
)

func main() {
    var here string=go_utils.GetHereDirExe()

    var datadir string=filepath.Join(here,"data")

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive: true,
        ErrorHandler: func(c fiber.Ctx, err error) error {
            fmt.Println("fiber error")
            fmt.Println(err)
            return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        },
    })


    // ---- apis ----
    // get item statistics for a target character
    app.Get("/get-builds",func(c fiber.Ctx) error {
        var character string=c.Query("character")
        var weapon string=c.Query("weapon")

        var datafileName string=erdata_builds.GetRouteDataFileName(
            character,
            weapon,
            datadir,
        )

        var routeData []erdata_builds.ErRoute2=erdata_builds.ReadRouteDataFile(
            datafileName,
        )

        var itemStatistics erdata_builds.GroupedItemStatistics=
            erdata_builds.ComputeAllItemStatistics(
                routeData,
            )

        return c.JSON(itemStatistics)
    })

    // get the available datafiles
    app.Get("/get-datafiles",func(c fiber.Ctx) error {
        var datafiles []erdata_builds.ErDataFileDescriptor=erdata_builds.GetErDataFiles(
            datadir,
        )

        return c.JSON(datafiles)
    })

    // open download config
    app.Get("/open-downloader-config",func(c fiber.Ctx) error {
        fmt.Println("Opening config file")
        go_utils.OpenTargetWithDefaultProgram(filepath.Join(here,"config/download-builds-config.yml"))

        return c.SendString("completed")
    })

    // run downloader program
    app.Get("/run-downloader",func(c fiber.Ctx) error {
        fmt.Println("running downloader")
        go_utils.OpenTargetWithDefaultProgram(filepath.Join(here,"builds-downloader.exe"))

        return c.SendString("completed")
    })


    // ---- static ----
    app.Static("/",filepath.Join(here,"er-builds-web/build"))

    go_utils.OpenTargetWithDefaultProgram("http://localhost:4200")
    app.Listen(":4200")
}