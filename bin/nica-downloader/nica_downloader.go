// nica downloader program. using the same config as builds downloader, tries
// to download corresponding nica data for all builds of a target character/weapon. builds downloader
// needs to be run first to get the list of builds

package main

import (
	"er-builds/lib/cli"
	"er-builds/lib/dak_gg"
	"er-builds/lib/erdata_builds"
	"er-builds/lib/nica"
	go_utils "er-builds/lib/utils"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	go_utils.ConfigureDefaultZeroLogger()

	// --- config
	var buildsDataDir string="../../data"

	var nicaBuildsDir string=filepath.Join(buildsDataDir,"nica")
	os.MkdirAll(nicaBuildsDir,0755)

	var charSelectConfig cli.CharactersSelectionConfig=cli.
		ReadCharactersSelectConfig("../../config/chars.yml")
	// --- end config


	var traitSkillsInfos dak_gg.TraitSkillMap=dak_gg.GetTraitSkillsInfoMap()

	for character := range charSelectConfig.CharacterSelections {
		for weaponI := range charSelectConfig.CharacterSelections[character] {
			var weapon string=charSelectConfig.CharacterSelections[character][weaponI]

			fmt.Println("getting for:",character,weapon)

			var routeDataFilename string=erdata_builds.GetRouteDataFileName(
				character,
				weapon,
				buildsDataDir,
			)

			var routedata []erdata_builds.ErRoute2=erdata_builds.ReadRouteDataFile(routeDataFilename)

			var nicaBuilds []nica.NicaBuild2

			for routeI := range routedata {
				fmt.Println("getting",routedata[routeI].Id)
				nicaBuilds=append(nicaBuilds,nica.GetBuild2(
					routedata[routeI].Id,
					traitSkillsInfos,
				))
			}

			fmt.Println("writing file")
			nica.WriteNicaBuilds(
				filepath.Join(nicaBuildsDir,routeDataFilename),
				nicaBuilds,
			)
		}
	}
}