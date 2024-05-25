package oer_api

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_readConfig(t *testing.T) {
	res:=readOerConfig("test/oer-config.yml")

	pp.Print(res)
}
