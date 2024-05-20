// functions for retrieving from aya.gg api

package aya_gg

import (
	"github.com/imroc/req/v3"
)

// top level response from aya gg api
type ApiDataResponse struct {
	Result struct {
		Characters []ApiCharacter
		CharacterWeapons []ApiCharacterWeapon
	}
}

// a character in api response
type ApiCharacter struct {
	Id int
	Profile struct {
		Name string
	}
}

// a character-weapon definition in api response
type ApiCharacterWeapon struct {
	CharacterId int

	// the name of the weapon
	ItemSubcategoryId string
}

// get aya gg "all" data
func GetAyaGGAllData() ApiDataResponse {
	var client *req.Client=req.C()

	var result ApiDataResponse
	var e error
	_,e=client.R().
		SetSuccessResult(&result).
		Get("https://aya.gg/erar/static/all")

	if e!=nil {
		panic(e)
	}

	return result
}