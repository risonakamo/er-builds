// implements funcs dealing with nica api
package nica

import (
	"er-builds/lib/dak_gg"

	"github.com/imroc/req/v3"
)

// get and convert to nica build 2
func GetBuild2(
    buildId int,
    traitSkills dak_gg.TraitSkillMap,
    client *req.Client,
) NicaBuild2 {
    return upgradeNicaBuildTo2(
        convRawToNicaBuild(
            getBuild(buildId,client),
        ),
        traitSkills,
    )
}