// nica downloader program. using the same config as builds downloader, tries
// to download corresponding nica data for all builds of a target character/weapon. builds downloader
// needs to be run first to get the list of builds

package main

import (
	"er-builds/lib/dak_gg"
	"er-builds/lib/erdata_builds"
	"er-builds/lib/nica"
)

func main() {
	var routedata []erdata_builds.ErRoute2=erdata_builds.
		ReadRouteDataFile("../../data/Mai-Whip.json")

	traitSkillsInfos:=dak_gg.GetTraitSkillsInfoMap()

	var nicaBuilds []nica.NicaBuild2

	for i := range routedata {
		nicaBuilds=append(nicaBuilds,nica.GetBuild2(routedata[i].Id,traitSkillsInfos))
	}
}