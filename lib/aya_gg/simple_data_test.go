package aya_gg

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_dataConvertTest(t *testing.T) {
    data:=GetAyaGGAllData()

    result:=ParseToSimpleCharData(data)

    pp.Print(result)
}