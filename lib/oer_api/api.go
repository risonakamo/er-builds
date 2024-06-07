// funcs matching oer api

package oer_api

import (
	"github.com/imroc/req/v3"
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

// get raw language file from api. language file is a giant string
func GetLanguageFile(apiKey string,language ErLang) string {
	var client *req.Client=req.C()

	var result LanguageFileUrlResponse
	var e error
	_,e=client.R().
		SetPathParam("lang",language).
		SetHeader("x-api-key",apiKey).
		SetHeader("accept","application/json").
		SetSuccessResult(&result).
		Get("https://open-api.bser.io/v1/l10n/{lang}")

	if e!=nil {
		panic(e)
	}

	if len(result.Data.L10Path)==0 {
		panic("lang file url was empty")
	}

	var resp *req.Response
	resp,e=client.R().
		Get(result.Data.L10Path)

	if e!=nil {
		panic(e)
	}

	return resp.String()
}