package nica

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_apiGet(t *testing.T) {
	result:=getBuild(50069)

	pp.Print(result)
}