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
	"strconv"

	"github.com/fatih/color"
	"github.com/imroc/req/v3"
	"github.com/rs/zerolog"
)

const Workers int=5

func main() {
	go_utils.ConfigureDefaultZeroLogger()
	zerolog.SetGlobalLevel(zerolog.WarnLevel)

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

	// for all characters, and all weapons of the character
	for character := range charSelectConfig.CharacterSelections {
		for weaponI := range charSelectConfig.CharacterSelections[character] {
			var weapon string=charSelectConfig.CharacterSelections[character][weaponI]

			fmt.Println()
			fmt.Printf("getting nica builds: %s %s\n",
				color.BlueString(character),
				color.BlueString(weapon),
			)

			// determine the route data file
			var routeDataFilename string=erdata_builds.GetRouteDataFileName(
				character,
				weapon,
				buildsDataDir,
			)

			// determine nica builds filename
			var nicaBuildFilename string=filepath.Join(
				nicaBuildsDir,
				nica.GetNicaBuildsFilename(character,weapon),
			)

			// read the route data file to know what build ids to get
			var routedata []erdata_builds.ErRoute2=erdata_builds.ReadRouteDataFile(routeDataFilename)

			// read the nica builds data to know what builds we already have
			var existingNicaBuilds []nica.NicaBuild2=nica.ReadNicaBuilds(
				nicaBuildFilename,
			)

			// getting new builds
			var buildsToGet []int=nica.NicaBuildDiff(routedata,existingNicaBuilds)

			if len(buildsToGet)==0 {
				fmt.Println(color.YellowString("no new builds to get"))
				continue
			}

			fmt.Printf("getting %s builds...\n",
				color.CyanString(strconv.Itoa(len(buildsToGet))),
			)

			var newNicaBuilds []nica.NicaBuild2=nica.GetBuilds2_mt(
				buildsToGet,
				traitSkillsInfos,
				client,

				Workers,
			)

			if len(newNicaBuilds)==0 {
				color.Yellow("no new builds downloaded")
				continue
			}

			// merge new builds with the existing builds. don't care about de-duplication for now,
			// since all the builds we get should not already be in the current nica builds
			existingNicaBuilds=append(existingNicaBuilds,newNicaBuilds...)



			fmt.Println("saving nica build file:",
				color.GreenString(nicaBuildFilename),
			)

			nica.WriteNicaBuilds(
				nicaBuildFilename,
				existingNicaBuilds,
			)
		}
	}

	fmt.Println()
	fmt.Println("completed")
	fmt.Println("press ENTER to continue")
	go_utils.WaitForEnterKey()
}