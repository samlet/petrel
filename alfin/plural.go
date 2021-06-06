package alfin

import (
	"github.com/fatih/camelcase"
	"strings"
)

var vowels = []string{"a", "e", "i", "o", "u"}

var pluralDictionary = map[string]string{
	"is":         "are",
	"analysis":   "analyses",
	"alumnus":    "alumni",
	"alumnae":    "alumni",
	"atlas":      "atlases",
	"appendix":   "appendices",
	"barrack":    "barracks",
	"beef":       "beefs",
	"brother":    "brothers",
	"cafe":       "cafes",
	"corpus":     "corpuses",
	"cow":        "cows",
	"ganglion":   "ganglions",
	"genus":      "genera",
	"graffito":   "graffiti",
	"loaf":       "loaves",
	"money":      "monies",
	"mongoose":   "mongooses",
	"move":       "moves",
	"mythos":     "mythoi",
	"niche":      "niches",
	"numen":      "numina",
	"octopus":    "octopuses",
	"opus":       "opuses",
	"ox":         "oxen",
	"penis":      "penises",
	"vagina":     "vaginas",
	"sex":        "sexes",
	"testis":     "testes",
	"turf":       "turfs",
	"tooth":      "teeth",
	"foot":       "feet",
	"cactus":     "cacti",
	"child":      "children",
	"criterion":  "criteria",
	"news":       "news",
	"datum":      "data",
	"deer":       "deer",
	"echo":       "echoes",
	"elf":        "elves",
	"embargo":    "embargoes",
	"foe":        "foes",
	"focus":      "foci",
	"fungus":     "fungi",
	"goose":      "geese",
	"hero":       "heroes",
	"hoof":       "hooves",
	"index":      "indices",
	"knife":      "knives",
	"leaf":       "leaves",
	"life":       "lives",
	"man":        "men",
	"mouse":      "mice",
	"nucleus":    "nuclei",
	"person":     "people",
	"phenomenon": "phenomena",
	"potato":     "potatoes",
	"self":       "selves",
	"syllabus":   "syllabi",
	"tomato":     "tomatoes",
	"torpedo":    "torpedoes",
	"veto":       "vetoes",
	"woman":      "women",
	"zero":       "zeroes",
}

var singularDictionary = map[string]string{
	"are":      "is",
	"analyses": "analysis",
	"alumni":   "alumnus",
	//"alumni": "alumnae", // for female - cannot have duplicate in map

	"genii":      "genius",
	"data":       "datum",
	"atlases":    "atlas",
	"appendices": "appendix",
	"barracks":   "barrack",
	"beefs":      "beef",
	"brothers":   "brother",
	"cafes":      "cafe",
	"corpuses":   "corpus",
	"cows":       "cow",
	"ganglions":  "ganglion",
	"genera":     "genus",
	"graffiti":   "graffito",
	"loaves":     "loaf",
	"monies":     "money",
	"mongooses":  "mongoose",
	"moves":      "move",
	"mythoi":     "mythos",
	"niches":     "niche",
	"numina":     "numen",
	"octopuses":  "octopus",
	"opuses":     "opus",
	"oxen":       "ox",
	"penises":    "penis",
	"vaginas":    "vagina",
	"sexes":      "sex",
	"testes":     "testis",
	"turfs":      "turf",
	"teeth":      "tooth",
	"feet":       "foot",
	"cacti":      "cactus",
	"children":   "child",
	"criteria":   "criterion",
	"news":       "news",
	"deer":       "deer",
	"echoes":     "echo",
	"elves":      "elf",
	"embargoes":  "embargo",
	"foes":       "foe",
	"foci":       "focus",
	"fungi":      "fungus",
	"geese":      "goose",
	"heroes":     "hero",
	"hooves":     "hoof",
	"indices":    "index",
	"knifes":     "knife",
	"leaves":     "leaf",
	"lives":      "life",
	"men":        "man",
	"mice":       "mouse",
	"nuclei":     "nucleus",
	"people":     "person",
	"phenomena":  "phenomenon",
	"potatoes":   "potato",
	"selves":     "self",
	"syllabi":    "syllabus",
	"tomatoes":   "tomato",
	"torpedoes":  "torpedo",
	"vetoes":     "veto",
	"women":      "woman",
	"zeroes":     "zero",
}

//https://www.socketloop.com/tutorials/golang-check-if-item-is-in-slice-array
func inVowels(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func AlreadyPluralized(test string) bool {
	if len(test) != 1 {

		// to handle words like genii, data and etc.
		if singularDictionary[test] != "" {

			return true
		}

		// put in some exceptions
		//if (string(test[len(test)-1]) != "s") || (string(test[len(test)-2]) != "ii") {
		if string(test[len(test)-1]) != "s" {

			if (string(test[len(test)-1:]) != "e") || (string(test[len(test)-1:]) != "y") {
				return false
			}

			if (string(test[len(test)-2:]) == "ch") || (string(test[len(test)-2:]) == "sh") || (string(test[len(test)-3:]) == "nes") {
				return false
			}

			if string(test[len(test)-3:]) == "ius" {
				return false
			}

			if pluralDictionary[test] == "" {
				return true
			}

		} else {
			return true
		}
	}
	return false

}

func Pluralize(singular string) string {

	plural := ""
	suffix := ""

	if singular != "" {

		root := singular

		// are we dealing with a single character?

		if len(singular) == 1 {
			return singular + "s"
		}

		plural = pluralDictionary[singular]

		// found an exact match in map
		if plural != "" {
			return plural
			//fmt.Println(plural)
		}

		// handle -sis
		if string(singular[len(singular)-3:]) == "sis" {
			root = string(singular[:len(singular)-3])
			suffix = "ses"
			plural = root + suffix
			return plural
		}

		// handle -z
		if string(singular[len(singular)-1:]) == "z" {
			root = string(singular[:len(singular)-1])
			suffix = "zes"
			plural = root + suffix
			return plural
		}

		// handle -ex
		if string(singular[len(singular)-2:]) == "ex" {
			root = string(singular[:len(singular)-2])
			suffix = "ices"
			plural = root + suffix
			return plural
		}

		// handle -ix
		if string(singular[len(singular)-2:]) == "ix" {
			root = string(singular[:len(singular)-2])
			suffix = "ices"
			plural = root + suffix
			return plural
		}

		// handle -us
		if string(singular[len(singular)-2:]) == "us" {
			root = string(singular[:len(singular)-2])
			suffix = "uses"
			plural = root + suffix
			return plural
		}

		// handle word such as dolly
		if (string(singular[len(singular)-1]) == "y") && !inVowels(string(singular[len(singular)-2]), vowels) {
			root = string(singular[:len(singular)-1])
			suffix = "ies"

		} else if string(singular[len(singular)-1]) == "s" {
			if inVowels(string(singular[len(singular)-2]), vowels) {
				if string(singular[len(singular)-3:]) == "ius" {
					root = string(singular[:len(singular)-2])
					suffix = "i"
				} else {
					root = string(singular[:len(singular)-1])
					suffix = "ses"
				}
			} else {
				suffix = "es"
			}
		} else if (string(singular[len(singular)-2:]) == "ch") || (string(singular[len(singular)-2:]) == "sh") {
			suffix = "es"
		} else {
			suffix = "s"
		}

		plural = root + suffix

		// sanity check
		if AlreadyPluralized(singular) {
			return singular
		} else {
			return plural
		}

	} else {
		return "Singular cannot be empty"
	}
}

func PluralizeTypeName(typeName string) string {
	words:=camelcase.Split(typeName)
	var result []string
	for i:=0;i<len(words)-1;i+=1{
		result=append(result, strings.ToLower(words[i]))
	}
	lastPart:=strings.ToLower(words[len(words)-1])
	result=append(result, Pluralize(lastPart))

	return strings.Join(result, "_")
}