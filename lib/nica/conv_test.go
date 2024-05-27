package nica

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_convToBuild(t *testing.T) {
	rawbuild:=getBuild(4529)

	result:=convRawToNicaBuild(rawbuild)

	pp.Print(result)
}