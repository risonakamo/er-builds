// funcs dealing with dak gg urls

package dak_gg

import "fmt"

// create url to access an item icon. does not include the base url to match
// with other urls
func CreateItemIconUrl(itemId int,gameVersion string) string {
    return fmt.Sprintf("/assets/er/game-assets/%s/ItemIcon_%d.png",gameVersion,itemId)
}