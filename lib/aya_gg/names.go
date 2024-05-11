// name conversion entities

package aya_gg

// mapping of some aya gg names to another name
// key: aya gg name
// val: other name
type WeaponNameMapDict map[string]string
var WeaponNameMap WeaponNameMapDict=WeaponNameMapDict{
	"DirectFire":"Shuriken",
	"HighAngleFire":"Throw",
	"OneHandSword":"Dagger",
}

// do weapon name conversion
func convertWeaponName(weapon string) string {
	var val string
	var in bool
	val,in=WeaponNameMap[weapon]

	if in {
		return val
	}

	return weapon
}