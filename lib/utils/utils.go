// general shared go util funcs

package go_utils

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// when called, gives the location of the file that called this function
func GetHereDir() string {
    var selfFilepath string
    _, selfFilepath, _, _ = runtime.Caller(1)

    return filepath.Dir(selfFilepath)
}

// give folder location of the exe that calls this func
func GetHereDirExe() string {
    var exePath string
    var e error
    exePath,e=os.Executable()

    if e!=nil {
        panic(e)
    }

    return filepath.Dir(exePath)
}

// set zerolog global logger default options
func ConfigureDefaultZeroLogger() {
    log.Logger=log.Output(zerolog.ConsoleWriter{
        Out:os.Stdout,
    })
}

// open chrome to specified url
func OpenChrome(url string) {
    var cmd *exec.Cmd=exec.Command("chrome.exe",url)
    var err error=cmd.Run()

    if err!=nil {
        panic(err)
    }
}

// try to open web url or file with default program.
// essentially runs program like it was double clicked
func OpenTargetWithDefaultProgram(url string) {
    var cmd *exec.Cmd=exec.Command("cmd","/c","start",url)
    var e error=cmd.Run()

    if e!=nil {
        panic(e)
    }
}

// pause until any key is pressed
func WaitForAnyKey() {
    var buf []byte=make([]byte,1)
    os.Stdin.Read(buf)
}

// pause until enter key is pressed
func WaitForEnterKey() {
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

// add # to all lines in a target file
func CommentAllInFile(filename string) {
    var file *os.File
    var e error
    file,e=os.Open(filename)

    if e!=nil {
        panic(e)
    }
    defer file.Close()

    var reader *bufio.Reader=bufio.NewReader(file)

    var modifedFile string=""
    for {
        var line string
        line,e=reader.ReadString('\n')

        if e!=nil {
            break
        }

        modifedFile+="# "+line
    }

    file.Close()

    var newFile *os.File
    newFile,e=os.Create(filename)
    newFile.WriteString(modifedFile)
}