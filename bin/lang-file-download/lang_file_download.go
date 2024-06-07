// lang file downloader. when run, looks for dev-config file in config folder to download
// langfile from api. places lang file in config folder

package main

import (
	"er-builds/lib/oer_api"
	go_utils "er-builds/lib/utils"
	"fmt"
	"path/filepath"

	"github.com/rs/zerolog"
)

func main() {
    go_utils.ConfigureDefaultZeroLogger()
    zerolog.SetGlobalLevel(zerolog.WarnLevel)

    var here string=go_utils.GetHereDirExe()

    // --- config
    var devConfigPath string=filepath.Join(here,"config/dev-config.yml")

    var langfilePath string=filepath.Join(here,"config/saved-langfile.txt")
    // --- end config

    var devConfig oer_api.OerApiConfig=oer_api.ReadOerConfig(devConfigPath)

    if len(devConfig.ApiKey)==0 {
        panic("failed to read api key")
    }

    fmt.Println("getting lang file from api...")
    var langfileString string=oer_api.GetLanguageFile(
        devConfig.ApiKey,
        oer_api.ErLang_english,
    )

    fmt.Println("writing lang file...")
    oer_api.WriteLangFile(langfilePath,langfileString)

    fmt.Println("completed")
}