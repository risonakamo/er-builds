package aya_gg

import (
	"testing"

	"github.com/kr/pretty"
)

func Test_apiTest(t *testing.T) {
	data:=GetAyaGGAllData()
    pretty.Print(data)
}