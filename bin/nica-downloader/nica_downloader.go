// nica downloader program. using the same config as builds downloader, tries
// to download corresponding nica data for all builds of a target character/weapon. builds downloader
// needs to be run first to get the list of builds

package main

import (
	"er-builds/lib/dak_gg"
	"er-builds/lib/erdata_builds"
	"er-builds/lib/nica"
	go_utils "er-builds/lib/utils"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	// --- config
	var buildsDataDir string="../../data"
	var datafileName string="Mai-Whip.json"
	// --- end config

	go_utils.ConfigureDefaultZeroLogger()

	// --- auto config vars
	var selectedDataFile string=filepath.Join(buildsDataDir,datafileName)
	var nicaBuildsDir string=filepath.Join(buildsDataDir,"nica")
	os.MkdirAll(nicaBuildsDir,0755)



	var routedata []erdata_builds.ErRoute2=erdata_builds.ReadRouteDataFile(selectedDataFile)

	traitSkillsInfos:=dak_gg.GetTraitSkillsInfoMap()

	var nicaBuilds []nica.NicaBuild2

	for i := range routedata {
		fmt.Println("getting",routedata[i].Id)
		nicaBuilds=append(nicaBuilds,nica.GetBuild2(routedata[i].Id,traitSkillsInfos))
	}

	fmt.Println("writing file")
	nica.WriteNicaBuilds(
		filepath.Join(nicaBuildsDir,datafileName),
		nicaBuilds,
	)
}