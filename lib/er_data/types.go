// er data types

package erdata

// item info dict
// key: item's id
// val: the item
type ItemInfoDict map[string]ItemInfo

// top level response when requesting for routes from route api
type ErRouteResponse struct {
    CharacterName string `json:"characterName"`
    ItemById ItemInfoDict `json:"itemById"`
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

// enhanced er route with weapons filled in
type ErRoute2 struct {
    ErRoute

    WeaponInfos []ItemInfo
}

// info about an item
type ItemInfo struct {
    Id int `json:"id"`
    Name string `json:"name"`
    ImageUrl string `json:"imageUrl"`
    BackgroundImageUrl string `json:"backgroundImageUrl"`
}