// funcs implementing cli for data downloader tool

package cli

import (
	"fmt"
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/kr/pretty"
)

// character/weapon pair
type CharacterWeapon struct {
    Character string
    Weapon string
}

func GetDataDownloaderCliArgs() {
	var parser *argparse.Parser=argparse.NewParser(
        "data_download",
        "builds data downloader tool",
    )

    charsString:=parser.String(
        "c",
        "characters",
        &argparse.Options{
            Help: "comma separated list of Character,Weapon to gather data for. "+
                "Separate each Character,Weapon with spaces",
        },
    )

    var e error=parser.Parse(os.Args)

    if e!=nil {
        panic(e)
    }

    var charslist []CharacterWeapon=parseCharsString(*charsString)
    pretty.Print(charslist)
}

// parse a space-seperated character,weapon list.
// example of a proper list: "Tia,Bat DebiMarlene,TwoHandSword Mai,Whip"
func parseCharsString(charString string) []CharacterWeapon {
    // split into Character,Weapon
    var split1 []string=strings.Split(charString," ")

    var pairs []CharacterWeapon

    for i := range split1 {
        var split2 []string=strings.Split(split1[i],",")

        if len(split2)!=2 {
            fmt.Println("failed to split character/weapon pair")
            fmt.Println("bad string:",split1[i])
            continue
        }

        pairs=append(pairs,CharacterWeapon{
            Character: split2[0],
            Weapon: split2[1],
        })
    }

    return pairs
}