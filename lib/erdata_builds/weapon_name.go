// functions dealing with weapon name

package erdata_builds

// convert long form weapon name that appears in tooltips to
// short name needed for api calls
var WeaponNameToShortName map[string]string=map[string]string{
    "Two-handed Sword":"TwoHandSword",
    "Assault Rifle":"AssaultRifle",
    "Sniper Rifle":"SniperRifle",
    "Dual Swords":"DualSword",
    "VF Prosthetic":"VFArm",
    "Crossbow":"CrossBow",
}