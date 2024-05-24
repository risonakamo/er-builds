package go_utils

import (
	"fmt"
	"testing"

	"github.com/k0kubun/pp/v3"
)

type TestYaml struct {
	Huh int `json:"huh" yaml:"huh"`
	Something string `json:"something" yaml:"something"`
	Thing []string `json:"thing" yaml:"thing"`
}

func Test_readYaml(t *testing.T) {
	var res TestYaml
	var e error
	res,e=ReadYaml[TestYaml]("test/test.yml")

	if e!=nil {
		panic(e)
	}

	pp.Print(res)
}

func Test_readWriteChain(t *testing.T) {
	var res TestYaml
	var e error
	res,e=ReadYaml[TestYaml]("test/test.yml")

	if e!=nil {
		panic(e)
	}

	fmt.Println("result of yml read")
	pp.Print(res)
	fmt.Println()

	e=WriteJson("test/test.json",res)

	if e!=nil {
		panic(e)
	}

	var res2 TestYaml
	res2,e=ReadJson[TestYaml]("test/test.json")

	if e!=nil {
		panic(e)
	}

	fmt.Println("result of json read")
	pp.Print(res2)
	fmt.Println()
}