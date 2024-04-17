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
    var here string=go_utils.GetHereDir()

    var datadir string=filepath.Join(here,"../../data")

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


    // ---- static ----
    app.Static("/",filepath.Join(here,"../../er-builds-web/build"))

    app.Listen(":4200")
}