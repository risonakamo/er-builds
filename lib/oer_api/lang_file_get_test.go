package oer_api

import (
	"fmt"
	"testing"
)

func Test_getItem(t *testing.T) {
	langfile:=readLangFile("test/saved-langfile.txt")
	langdict:=parseLangFile(langfile)

	got,e:=getItemName(langdict,502401)

    if e!=nil {
        panic(e)
    }

    fmt.Println(got)

    got,e=getItemName(langdict,101201)

    if e!=nil {
        panic(e)
    }

    fmt.Println(got)

    // this should fail
    got,e=getItemName(langdict,1047200)

    if e==nil {
        t.Error("found item when it shouldn't have")
    }
}