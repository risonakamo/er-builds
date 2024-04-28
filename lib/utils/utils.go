// general shared go util funcs

package go_utils

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// when called, gives the location of the file that called this function
func GetHereDir() string {
    var selfFilepath string
    _, selfFilepath, _, _ = runtime.Caller(1)

    return filepath.Dir(selfFilepath)
}

// give folder location of the exe that calls this func
func GetHereDirExe() string {
    var exePath string
    var e error
    exePath,e=os.Executable()

    if e!=nil {
        panic(e)
    }

    return filepath.Dir(exePath)
}

// set zerolog global logger default options
func ConfigureDefaultZeroLogger() {
    log.Logger=log.Output(zerolog.ConsoleWriter{
        Out:os.Stdout,
    })
}