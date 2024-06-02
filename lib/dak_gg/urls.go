// funcs dealing with dak gg urls

package dak_gg

import "fmt"

// todo: move this to a config file
const GameVersion string="1.22.1"

// create url to access an item icon. does not include the base url to match
// with other urls
func CreateItemIconUrl(itemId int) string {
    return fmt.Sprintf("/assets/er/game-assets/%s/ItemIcon_%d.png",GameVersion,itemId)
}