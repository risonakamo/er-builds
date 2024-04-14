// filesystem functions for writing and reading data

package erdata_builds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// given er route data, try to merge it into the data in the target file,
// and save the file. the data should be .json.
// if the file does not exist, will be created
// merged data will ensure no duplicate Route ids
func MergeDataIntoFile(data []ErRoute2,datafile string) {
    var readData []ErRoute2=ReadRouteDataFile(datafile)
    var originalLen int=len(readData)

    readData=append(readData,data...)

    readData=purgeDuplicateRoutes(readData)
    fmt.Println("data file changed by:",len(readData)-originalLen)

    writeRouteDataFile(readData,datafile)
    fmt.Println("updated",datafile)
}

// read er route data file. if file does not exist, returns empty
func ReadRouteDataFile(datafile string) []ErRoute2 {
    var data []byte
    var e error
    data,e=os.ReadFile(datafile)

    if errors.Is(e,fs.ErrNotExist) {
        fmt.Println("route data file was missing, returning empty")
        return []ErRoute2{}
    }

    if e!=nil {
        panic(e)
    }

    var parsedData []ErRoute2
    json.Unmarshal(data,&parsedData)
    return parsedData
}

// generate the correct name to access a character/weapon's route data.
// data files should be made with this name
// todo: watch out for characters/weapons with spaces. convert to underscores?
func GetRouteDataFileName(
    character string,
    weapon string,
    dataDir string,
) string {
    return filepath.Join(
        dataDir,
        fmt.Sprintf("%s-%s",character,weapon),
    )
}

// overwrite target file with the provided data
func writeRouteDataFile(data []ErRoute2,datafile string) {
    var wfile *os.File
    var e error
    wfile,e=os.Create(datafile)

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
}

// write item statistics to json. mostly used for debug, don't actually want to
// ever save these except maybe for cache optimisation
func writeItemStatistics(itemStats GroupedItemStatistics,filename string) {
    var wfile *os.File
    var e error
    wfile,e=os.Create(filename)

    if e!=nil {
        panic(e)
    }

    defer wfile.Close()

    var jsondata []byte
    jsondata,e=json.Marshal(itemStats)

    if e!=nil {
        panic(e)
    }

    wfile.Write(jsondata)
}