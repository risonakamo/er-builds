// funcs helping with file read/write

package go_utils

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// read a yaml file and return result
func ReadYaml[DataT any](filename string) (DataT,error) {
	var data []byte
	var e error
	data,e=os.ReadFile(filename)

	if errors.Is(e,fs.ErrNotExist) {
		log.Info().Msgf("file not found: %s",filename)
		var def DataT
		return def,e
	}

	if e!=nil {
		var def DataT
		return def,e
	}

	var parsedData DataT
	yaml.Unmarshal(data,&parsedData)

	return parsedData,nil
}

// overwrite target yml file with a new file
func WriteYaml(filename string,data any) error {
	var wfile *os.File
	var e error
	wfile,e=os.Create(filename)

	if e!=nil {
		panic(e)
	}

	defer wfile.Close()

	var ymldata []byte
	ymldata,e=yaml.Marshal(data)

	if e!=nil {
		panic(e)
	}

	wfile.Write(ymldata)
	return nil
}

// read json file and return result
func ReadJson[DataT any](filename string) (DataT,error) {
	var data []byte
	var e error
	data,e=os.ReadFile(filename)

	if errors.Is(e,fs.ErrNotExist) {
		log.Info().Msgf("file not found: %s",filename)
		var def DataT
		return def,e
	}

	if e!=nil {
		var def DataT
		return def,e
	}

	var parsedData DataT
	json.Unmarshal(data,&parsedData)

	return parsedData,nil
}

// overwrite target json file with a new file
func WriteJson(filename string,data any) error {
	var wfile *os.File
	var e error
	wfile,e=os.Create(filename)

	if e!=nil {
		panic(e)
	}

	defer wfile.Close()

	var jsondata []byte
	jsondata,e=json.Marshal(data)

	if e!=nil {
		panic(e)
	}

	wfile.Write(jsondata)
	return nil
}