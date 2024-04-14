package main

import (
	"fmt"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v3"
)

func main() {
    var here string=getHereDir()

    var app *fiber.App=fiber.New(fiber.Config{
        CaseSensitive: true,
        ErrorHandler: func(c fiber.Ctx, err error) error {
            fmt.Println("fiber error")
            fmt.Println(err)
            return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
        },
    })

    app.Static("/",filepath.Join(here,"../er-builds-web/build"))

    app.Listen(":4200")
}

// get directory of main function
func getHereDir() string {
    var selfFilepath string
    _, selfFilepath, _, _ = runtime.Caller(0)

    return filepath.Dir(selfFilepath)
}