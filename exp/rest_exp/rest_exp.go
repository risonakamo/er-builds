package main

import (
	"io"
	"net/http"
	"os"
)

// top level response when requesting for routes from route api
type ErRouteResponse struct {
    CharacterName string `json:"characterName"`
    ItemById map[string]ItemInfo `json:"itemById"`
    RecommendWeaponRouteDtoPage RoutesContainer `json:"recommendWeaponRouteDtoPage"`
}

// contains list of build routes
type RoutesContainer struct {
    Total int `json:"total"`
    Page int `json:"page"`
    HasNext bool `json:"hasNext"`
    Items []ErRoute `json:"items"`
}

// a single route
type ErRoute struct {
    Id int `json:"id"`
    Title string `json:"title"`
    Version string `json:"version"`

    WeaponIds []int `json:"weaponIds"`
    PathIds []int `json:"pathIds"`

    UpdateDtm int `json:"updateDtm"`
    Likes int `json:"v2Like"`
    WinRate float32 `json:"v2WinRate"`
}

// info about an item
type ItemInfo struct {
    Id int `json:"id"`
    Name string `json:"name"`
    ImageUrl string `json:"imageUrl"`
    BackgroundImageUrl string `json:"backgroundImageUrl"`
}

func main() {
	resp,e:=http.Get("https://er-node.dakgg.io/api/v0/routes?hl=en&character=Tia&weaponType=Bat")

    if e!=nil {
        panic(e)
    }

    defer resp.Body.Close()

    data,e:=io.ReadAll(resp.Body)

    wfile,e:=os.Create("test.json")

    if e!=nil {
        panic(e)
    }

    wfile.Write(data)
}