// functions implementing filesystem api for nica

package nica

import (
	go_utils "er-builds/lib/utils"
	"fmt"
)

// create filename to access a nica builds json file
func GetNicaBuildsFilename(character string,weapon string) string {
	return fmt.Sprintf("%s-%s.json",character,weapon)
}

// write nica builds file
func WriteNicaBuilds(filename string,data []NicaBuild2) {
	var e error=go_utils.WriteJson(filename,data)

	if e!=nil {
		panic(e)
	}
}

// read nica builds file
func ReadNicaBuilds(filename string) []NicaBuild2 {
	var data []NicaBuild2
	var e error
	data,e=go_utils.ReadJson[[]NicaBuild2](filename)

	if e!=nil {
		panic(e)
	}

	return data
}