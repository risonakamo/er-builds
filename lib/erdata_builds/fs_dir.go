// functions for managing er data dir

package erdata_builds

import "os"

// scan directory and get all data
func GetErDataFiles(datadir string) []ErDataFileDescriptor {
	var dirItems []os.DirEntry
	var e error

	dirItems,e=os.ReadDir(datadir)

	if e!=nil {
		panic(e)
	}

	var foundFiles []ErDataFileDescriptor

	for i := range dirItems	{
		// only evaluate files
		if dirItems[i].IsDir() {
			continue
		}

		var foundfile ErDataFileDescriptor
		foundfile,e=parseRouteDataFileName(dirItems[i].Name())

		if e!=nil {
			continue
		}

		foundFiles=append(foundFiles,foundfile)
	}

	return foundFiles
}