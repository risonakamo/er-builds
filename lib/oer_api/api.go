package oer_api

import (
	"github.com/imroc/req/v3"
	"github.com/k0kubun/pp/v3"
)

// reponse from get language file api. this response further contains another url
// that needs to be called for the actual language file
type LanguageFileUrlResponse struct {
	Data struct {
		L10Path string
	}
}

// enum of all possible language values usable with oer api
type ErLang=string
const (
	ErLang_english ErLang="English"
	ErLang_japanese ErLang="Japanese"
)

func getLanguageFile(apiKey string,language ErLang) {
	var client *req.Client=req.C()

	var result LanguageFileUrlResponse
	var e error
	var resp *req.Response
	resp,e=client.R().
		SetPathParam("lang",language).
		SetHeader("x-api-key",apiKey).
		SetHeader("accept","application/json").
		SetSuccessResult(&result).
		Get("https://open-api.bser.io/v1/l10n/{lang}")

	if e!=nil {
		panic(e)
	}

	pp.Print(resp.String())

	// var result2 string
	// _,e=client.R().
	// 	SetSuccessResult(&result2).
	// 	Get(result.Data.L10Path)

	// if e!=nil {
	// 	panic(e)
	// }

	// pp.Print(result2)
}