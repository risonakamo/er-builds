package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	req,e:=http.NewRequest(http.MethodGet,"https://dak.gg/er/routes?character=Tia&weaponType=Bat",nil)

	if e!=nil {
		panic(e)
	}

	req.Header.Add("Cookie","locale=en")

	client:=&http.Client {}

	resp,e:=client.Do(req)

	if e!=nil {
		panic(e)
	}

	defer resp.Body.Close()

	doc,e:=goquery.NewDocumentFromReader(resp.Body)

	if e!=nil {
		panic(e)
	}

	doc.Find("section h2").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Text())
		if s.Text()=="Select Saved Plans" {
			fmt.Println(s.Next().Html())
			s.Next().Children().Each(func(i int, s2 *goquery.Selection) {
				s2.Find("img").Each(func(i int, s3 *goquery.Selection) {
					fmt.Println(s3.Attr("alt"))
				})
			})
		}
	})
}