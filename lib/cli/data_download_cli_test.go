package cli

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_readSelectConfig(t *testing.T) {
    res:=readCharactersSelectConfig("../../config/chars.yml")
    pp.Print(res)
}