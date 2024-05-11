// exe to generate the builds download yml. generates to the config dir

package main

import (
	"er-builds/lib/aya_gg"
	go_utils "er-builds/lib/utils"
	"fmt"
	"path/filepath"
)

func main() {
	var saveLocation string="../../config/chars.yml"

	fmt.Println("retrieving data...")
	var data aya_gg.ApiDataResponse=aya_gg.GetAyaGGAllData()

	fmt.Println("parsing data...")
	var simpleData aya_gg.SimpleCharDataDict=aya_gg.ParseToSimpleCharData(data)

	var fullSaveLocation string
	fullSaveLocation,_=filepath.Abs(saveLocation)
	fmt.Println("saving to:",fullSaveLocation)
	aya_gg.WriteSimpleDataFile(saveLocation,simpleData)

	go_utils.CommentAllInFile(saveLocation)
}