// oer api config file access api

package oer_api

import go_utils "er-builds/lib/utils"

// yml config file that needs to be provided to use oer api
type OerApiConfig struct {
	ApiKey string `yaml:"apiKey"`
}

// get oer config
func readOerConfig(filename string) OerApiConfig {
	var res OerApiConfig
	var e error
	res,e=go_utils.ReadYaml[OerApiConfig](filename)

	if e!=nil {
		panic(e)
	}

	return res
}