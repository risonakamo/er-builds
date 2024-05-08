// main exported functions

package aya_gg

// save aya gg to a yml file. give the full filename with extension
func GetAyaGGToFile(filename string) {
	var data ApiDataResponse=GetAyaGGAllData()
	var simpleData SimpleCharDataDict=parseToSimpleCharData(data)
	writeSimpleDataFile(filename,simpleData)
}