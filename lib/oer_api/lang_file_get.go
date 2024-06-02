// funcs to get specific data out of lang file

package oer_api

import (
	"errors"
	"strconv"
)

// get the name of an item. error if could not find.
func GetItemName(langdict OerLangDict,itemId int) (string,error) {
    var val string
    var in bool
    val,in=langdict.Nested["Item"].Nested["Name"].Fields[strconv.Itoa(itemId)]

    if !in {
        return "",errors.New("failed to find item id")
    }

    return val,nil
}