// general shared go util funcs

package go_utils

import (
	"path/filepath"
	"runtime"
)

// when called, gives the location of the file that called this function
func GetHereDir() string {
    var selfFilepath string
    _, selfFilepath, _, _ = runtime.Caller(1)

    return filepath.Dir(selfFilepath)
}