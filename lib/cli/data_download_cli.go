// funcs implementing cli for data downloader tool

package cli

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/akamensky/argparse"
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

// data downloader args. just list of character weapon combos to use
type DataDownloaderArgs struct {
    Selections []CharacterWeapon
    Versions []string
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

// config yml used to select versions and characters/weapons to download
type CharactersSelectionConfig struct {
    Versions []string
    CharacterSelections CharactersSelection `yaml:"characters"`
}

// get cli args for data downloader tool. configsdir should be relative to here dir
func GetDataDownloaderCliArgs(
    hereDir string,
    configsDir string,
) DataDownloaderArgs {
	var parser *argparse.Parser=argparse.NewParser(
        "data_download",
        "builds data downloader tool. "+
        "If -c not given, will read from config file config/chars.yml. Otherwise, "+
        "only uses the character/weapons specified by -c. Use double quotes (\") and spaces to select "+
        "multiple character/weapons",
    )

    var charsString *string=parser.String(
        "c",
        "characters",
        &argparse.Options{
            Help: "comma separated list of Character,Weapon to gather data for. "+
                "Separate each Character,Weapon with spaces. Example: \"Tia,Bat Mai,Whip\"",
        },
    )

    var e error=parser.Parse(os.Args)

    if e!=nil {
        panic(e)
    }

    var charslist []CharacterWeapon

    var config CharactersSelectionConfig=ReadCharactersSelectConfig(
        filepath.Join(hereDir,configsDir,"chars.yml"),
    )

    if len(*charsString)>0 {
        charslist=parseCharsString(*charsString)
    } else {
        charslist=characterSelectionsToPairs(config.CharacterSelections)
    }

    return DataDownloaderArgs {
        Selections: charslist,
        Versions: config.Versions,
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
func ReadCharactersSelectConfig(filepath string) CharactersSelectionConfig {
    var data []byte
    var e error
    data,e=os.ReadFile(filepath)

    if e!=nil {
        log.Fatal().Err(e).Msg("failed to find character config")
    }

    var parsedData CharactersSelectionConfig
    e=yaml.Unmarshal(data,&parsedData)

    if e!=nil {
        log.Fatal().Err(e).Msg("failed to parse character config")
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