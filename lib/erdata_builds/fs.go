// filesystem functions for writing and reading data

package erdata_builds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

// given er route data, try to merge it into the data in the target file,
// and save the file. the data should be .json.
// if the file does not exist, will be created
// merged data will ensure no duplicate Route ids.
// ensures the datafile's dir exists
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
        fmt.Sprintf("%s-%s.json",character,weapon),
    )
}

// parse filename into er data file descriptor, or error if bad name
func parseRouteDataFileName(filename string) (ErDataFileDescriptor,error) {
    var reg *regexp.Regexp=regexp.MustCompile(`(\w+)-(\w+).json`)

    var matches []string=reg.FindStringSubmatch(filename)

    if len(matches)!=3 {
        return ErDataFileDescriptor{},errors.New("bad match")
    }

    return ErDataFileDescriptor {
        Character: matches[1],
        Weapon: matches[2],

        Filename: filepath.Base(filename),
    },nil
}

// overwrite target file with the provided data
func writeRouteDataFile(data []ErRoute2,datafile string) {
    os.MkdirAll(filepath.Dir(datafile),0755)

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