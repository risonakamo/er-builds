// er data types

package erdata_builds

// item info dict
// key: item's id
// val: the item
type ItemInfoDict map[string]ItemInfo

// map of multiple item statistics
// key: item's ID
// val: the item's statistics
type ItemStatisticsDict map[int]*ItemsStatistics

// list of item statistics grouped by their type
// key: the item type
// val: all item statistics objs that have this type
type GroupedItemStatistics map[ItemType][]ItemsStatistics

// possible types of items
type ItemType string
const (
    ItemType_weapon ItemType="weapon"
    ItemType_head ItemType="head"
    ItemType_chest ItemType="chest"
    ItemType_arm ItemType="arm"
    ItemType_leg ItemType="leg"
    ItemType_tacskill ItemType="tacskill"
    ItemType_augment ItemType="augment"
)

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

    ItemInfos []ItemInfo2
    MainWeapon string
}

// info about an item
type ItemInfo struct {
    Id int `json:"id"`
    Name string `json:"name"`

    Tooltip string `json:"tooltip"`

    ImageUrl string `json:"imageUrl"`
    BackgroundImageUrl string `json:"backgroundImageUrl"`
}

// upgraded item with type field added. item info 2 can represent
//
// - equipmment (armour/weapon)
// - tac skill
// - augment (main and sub)
//
// this is because item info is just the id, name, tooltip, and icon, of which
// all of these have it
type ItemInfo2 struct {
    ItemInfo

    ItemType ItemType `json:"itemType"`

    // filled out if item type was not arm, leg, chest, head. otherwise empty.
    WeaponName string
}

// statistics of a certain item. computed from a list of ErRoute2
type ItemsStatistics struct {
    Item ItemInfo2 `json:"itemInfo"`

    // number of times this item showed up
    Total int `json:"totalBuilds"`
    // out of the number of builds used to calculate these stats, what is
    // the percentage of this item appearing in those builds
    BuildsPercentage float32 `json:"buildsPercentage"`

    Likes int `json:"likes"`

    TotalWinRate float32 `json:"totalWinRate"`
    AverageWinRate float32 `json:"averageWinRate"`
    HighestWinRate float32 `json:"highestWinRate"`
}

// information about a er routes datafile. currently all header data is
// stored in the filename. eventually, might add a header to the json? or store
// seperately?
type ErDataFileDescriptor struct {
    Character string `json:"character"`
    Weapon string `json:"weapon"`

    // is NOT full file path. only filename with extension
    Filename string `json:"filename"`
}