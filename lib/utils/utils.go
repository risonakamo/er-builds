// general shared go util funcs

package go_utils

import (
	"path/filepath"
	"runtime"
)

// get directory of main function
func GetHereDir() string {
    var selfFilepath string
    _, selfFilepath, _, _ = runtime.Caller(0)

    return filepath.Dir(selfFilepath)
}