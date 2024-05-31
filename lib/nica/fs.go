// functions implementing filesystem api for nica

package nica

import (
	go_utils "er-builds/lib/utils"
	"errors"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
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

// read nica builds file. if file does not exist, return empty without crashing
func ReadNicaBuilds(filename string) []NicaBuild2 {
	var data []NicaBuild2
	var e error
	data,e=go_utils.ReadJson[[]NicaBuild2](filename)

	if errors.Is(e,os.ErrNotExist) {
		log.Warn().Msgf("nica build did not exist - ignoring: %s",filename)
		return []NicaBuild2{}
	}

	if e!=nil {
		panic(e)
	}

	return data
}