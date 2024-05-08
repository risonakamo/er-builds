// functions for retrieving from aya.gg api

package aya_gg

import (
	"encoding/json"
	"io"
	"net/http"
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
	ItemSubcategoryId string
}

// get aya gg "all" data
func GetAyaGGAllData() ApiDataResponse {
	var e error

	var req *http.Request
	req,e=http.NewRequest(
		http.MethodGet,
		"https://aya.gg/erar/static/all",
		nil,
	)

	if e!=nil {
		panic(e)
	}

	var client http.Client=http.Client{}

	var resp *http.Response
	resp,e=client.Do(req)

	if e!=nil {
		panic(e)
	}

	defer resp.Body.Close()

	var data []byte
	data,e=io.ReadAll(resp.Body)

	if e!=nil {
		panic(e)
	}

	var allData ApiDataResponse
	e=json.Unmarshal(data,&allData)

	if e!=nil {
		panic(e)
	}

	return allData
}