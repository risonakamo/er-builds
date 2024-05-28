// implements funcs dealing with nica api
package nica

import "er-builds/lib/dak_gg"

// get and convert to nica build 2
func GetBuild2(buildId int,traitSkills dak_gg.TraitSkillMap) NicaBuild2 {
    return upgradeNicaBuildTo2(
        convRawToNicaBuild(
            getBuild(buildId),
        ),
        traitSkills,
    )
}