// filesystem functions for writing and reading data

package erdata_builds

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
)

// given er route data, try to merge it into the data in the target file,
// and save the file. the data should be .json.
// if the file does not exist, will be created
// merged data will ensure no duplicate Route ids
func MergeDataIntoFile(data []ErRoute2,datafile string) {
    var readData []ErRoute2=readRouteDataFile(datafile)
    var originalLen int=len(readData)

    readData=append(readData,data...)

    readData=purgeDuplicateRoutes(readData)
    fmt.Println("data file changed by:",len(readData)-originalLen)

    writeRouteDataFile(readData,datafile)
    fmt.Println("updated",datafile)
}

// read er route data file. if file does not exist, returns empty
func readRouteDataFile(datafile string) []ErRoute2 {
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