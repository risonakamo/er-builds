// er data types

package erdata

// item info dict
// key: item's id
// val: the item
type ItemInfoDict map[string]ItemInfo

// possible types of items
type ItemType string
const ItemType_weapon ItemType="weapon"
const ItemType_head ItemType="head"
const ItemType_chest ItemType="chest"
const ItemType_arm ItemType="arm"
const ItemType_leg ItemType="leg"

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

// enhanced er route with items converted into obj form and upgraded
type ErRoute2 struct {
    ErRoute

    WeaponInfos []ItemInfo2
}

// info about an item
type ItemInfo struct {
    Id int `json:"id"`
    Name string `json:"name"`

    Tooltip string `json:"tooltip"`

    ImageUrl string `json:"imageUrl"`
    BackgroundImageUrl string `json:"backgroundImageUrl"`
}

// upgraded item with type field added
type ItemInfo2 struct {
    ItemInfo

    ItemType ItemType
}