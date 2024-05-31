// nica downloader program. using the same config as builds downloader, tries
// to download corresponding nica data for all builds of a target character/weapon. builds downloader
// needs to be run first to get the list of builds
//
// looks for folders relative to exe

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

	"github.com/imroc/req/v3"
)

func main() {
	go_utils.ConfigureDefaultZeroLogger()

	var here string=go_utils.GetHereDirExe()

	// --- config
	var buildsDataDir string=filepath.Join(here,"data")

	var nicaBuildsDir string=filepath.Join(buildsDataDir,"nica")
	os.MkdirAll(nicaBuildsDir,0755)

	var charSelectConfig cli.CharactersSelectionConfig=cli.
		ReadCharactersSelectConfig(
			filepath.Join(here,"config/chars.yml"),
		)
	// --- end config


	var traitSkillsInfos dak_gg.TraitSkillMap=dak_gg.GetTraitSkillsInfoMap()

	var client *req.Client=req.C()

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
					client,
				))
			}

			fmt.Println("writing file")
			var nicaBuildFilename string=nica.GetNicaBuildsFilename(character,weapon)
			nica.WriteNicaBuilds(
				filepath.Join(nicaBuildsDir,nicaBuildFilename),
				nicaBuilds,
			)
		}
	}
}