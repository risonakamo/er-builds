// funcs implementing cli for data downloader tool

package cli

import (
	"os"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// data downloader args. just list of character weapon combos to use
type DataDownloaderArgs struct {
    Selections []CharacterWeapon
}

// character/weapon pair
type CharacterWeapon struct {
    Character string
    Weapon string
}

// selection of characters.
// key: character name
// val: weapon names for that character
type CharactersSelection map[string][]string

// get cli args for data downloader tool
func GetDataDownloaderCliArgs() DataDownloaderArgs {
	var parser *argparse.Parser=argparse.NewParser(
        "data_download",
        "builds data downloader tool",
    )

    var charsString *string=parser.String(
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

    var charslist []CharacterWeapon

    if len(*charsString)>0 {
        charslist=parseCharsString(*charsString)
    } else {
        var charsConfig CharactersSelection=readCharactersSelectConfig("download-builds-config.yml")
        charslist=characterSelectionsToPairs(charsConfig)
    }

    return DataDownloaderArgs {
        Selections: charslist,
    }
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
            log.Warn().
                Str("bad input",split1[i]).
                Msg("failed to split character/weapon pair, skipping")

            continue
        }

        pairs=append(pairs,CharacterWeapon{
            Character: split2[0],
            Weapon: split2[1],
        })
    }

    return pairs
}

// read character selection yml file
func readCharactersSelectConfig(filepath string) CharactersSelection {
    var data []byte
    var e error
    data,e=os.ReadFile(filepath)

    if e!=nil {
        panic(e)
    }

    var parsedData CharactersSelection=make(CharactersSelection)
    e=yaml.Unmarshal(data,&parsedData)

    if e!=nil {
        panic(e)
    }

    return parsedData
}

// convert characters selection dict into character/weapon array
func characterSelectionsToPairs(selections CharactersSelection) []CharacterWeapon {
    var res []CharacterWeapon

    // for all selections. selections[i] is character name
    for i := range selections {
        // for all weapons of a character.
        // selections[i][i2] is a weapon
        for i2 := range selections[i] {
            res=append(res,CharacterWeapon{
                Character: i,
                Weapon: selections[i][i2],
            })
        }
    }

    return res
}