// functions implementing the simplified data format

package aya_gg

// simple char data in dict form.
// key: char name
// val: list of weapons of that character
type SimpleCharDataDict map[string][]string

// names of characters keyed by their ids
// key: char id
// val: char name
type CharsByCharId map[int]string

// weapons organised by their character id.
// key: char id
// val: list of weapons that char uses
type WeaponsByCharId map[int][]string

// simplified collected form of character info
type SimpleChardata struct {
    Name string
    Weapon string
}

// convert aya gg data response into list of simple char data
func parseToSimpleCharData(data ApiDataResponse) SimpleCharDataDict {
    var weaponsDict WeaponsByCharId=groupWeaponsByCharIds(data.Result.CharacterWeapons)
    var charsDict CharsByCharId=groupCharsByCharId(data.Result.Characters)

    return convertWeaponsDictToSimpleCharsDict(weaponsDict,charsDict)
}

// convert api character weapon to dict of char ids and the weapons of that char
func groupWeaponsByCharIds(charWeapons []ApiCharacterWeapon) WeaponsByCharId {
    var weaponsByCharId WeaponsByCharId=make(WeaponsByCharId)

    var i int
    for i = range charWeapons {
        var charId int=charWeapons[i].CharacterId

        var in bool
        _,in=weaponsByCharId[charId]

        if !in {
            weaponsByCharId[charId]=[]string{}
        }

        weaponsByCharId[charId]=append(
            weaponsByCharId[charId],
            charWeapons[i].ItemSubcategoryId,
        )
    }

    return weaponsByCharId
}

// key characters by their id
func groupCharsByCharId(chars []ApiCharacter) CharsByCharId {
    var result CharsByCharId=make(CharsByCharId)

    var i int
    for i = range chars {
        result[chars[i].Id]=chars[i].Profile.Name
    }

    return result
}

// convert weapons by char id dict to simple char dict, which has the key of the
// dict be the name instead of the id
func convertWeaponsDictToSimpleCharsDict(
    weaponsDict WeaponsByCharId,
    charsDict CharsByCharId,
) SimpleCharDataDict {
    var result SimpleCharDataDict=make(SimpleCharDataDict)

    var charId int
    var weapons []string
    for charId,weapons = range weaponsDict {
        result[charsDict[charId]]=weapons
    }

    return result
}