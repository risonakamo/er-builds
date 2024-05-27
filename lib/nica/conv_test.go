package nica

import (
	"er-builds/lib/dak_gg"
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_convToBuild(t *testing.T) {
	rawbuild:=getBuild(4529)

	result:=convRawToNicaBuild(rawbuild)

	pp.Print(result)
}

func Test_convToBuild2(t *testing.T) {
	rawbuild:=getBuild(4529)
	build1:=convRawToNicaBuild(rawbuild)

	traitSkillsInfos:=dak_gg.GetTraitSkillsInfoMap()

	res:=upgradeNicaBuildTo2(build1,traitSkillsInfos)

	pp.Print(res)
}