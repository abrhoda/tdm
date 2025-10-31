package tdm

import (
	"encoding/json"
	"fmt"
	"github.com/abrhoda/tdm/internal"
	"github.com/abrhoda/tdm/internal/foundry"
	"github.com/abrhoda/tdm/storage"
	"os"
	"path/filepath"
	"strings"
)

const packs = "/packs/"
var journalFiles = []string{"journals/ancestries.json", "journals/archetypes.json", "journals/classes.json"}

var contentsToDirs = map[string][]string{
	//"actions": {"actions"},
	"ancestries":  {"ancestries", "ancestryfeatures"},
	"backgrounds": {"backgrounds"},
	//"bestiaries": {.. list them all ..},
	"classes": {"classes", "classfeatures"},
	//"conditions": {"conditions"},
	//"deities": {"deities"},
	"effects":   {"other-effects"},
	"equipment": {"equipment", "equipment-effects"},
	"feats":     {"feats", "feat-effects"},
	"heritages": {"heritages"},
	//"hazards":" {"hazards"}
	"spells": {"spells", "spell-effects"},

	// TODO others to include:
	// hazards, other-effects (this is like aid), deities, conditions, bestiaries, actions
}



// TODO out slice should have a capacity to avoid reallocations when adding elements.
func walkDir[T foundry.FoundryModel](fullpath string, noLegacyContent bool, licenses []string) ([]T, error) {
	out := make([]T, 0)

	err := filepath.WalkDir(fullpath, func(fullpath string, dirEntry os.DirEntry, err error) error {
		if err != nil {
fmt.Printf("Error for entry %s. Error: %v", fullpath, err)
			fmt.Printf("got error: %v\n", err)
			return err
		}

		if dirEntry.IsDir() || !strings.HasSuffix(dirEntry.Name(), ".json") || dirEntry.Name() == "_folders.json" {
			fmt.Printf("Not processing %s\n", dirEntry.Name())
			return nil
		}

		content, err := os.ReadFile(fullpath)
		if err != nil {
			fmt.Printf("Error reading entry %s. Error: %v", fullpath, err)
			return err
		}

		fmt.Printf("DEBUG: processing file: %s\n", fullpath)
		var data T
		err = json.Unmarshal(content, &data)
		//fmt.Printf("result: %v\n", data)
		if err != nil {
			return err
		}

		// filter out legacy content if needed.
		if noLegacyContent && data.IsLegacy() {
			fmt.Printf("noLegacyContent (%t) && data.IsLegacy (%t) is true", noLegacyContent, data.IsLegacy())
			return nil
		}

		// ensure `data`'s license is in the provided licenses.
		for _, l := range licenses {
			if data.HasProvidedLicense(l) {
				out = append(out, data)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return out, nil
}

func readJournalFiles(partialpath string) ([]foundry.Journal, error) {
	journals := make([]foundry.Journal, len(journalFiles))
	
	for i, file := range journalFiles {
		content, err := os.ReadFile(partialpath + file)
		if err != nil {
			fmt.Printf("Error reading journal file %s. Error: %v", partialpath + file, err)
			return nil, err
		}

		var j foundry.Journal
		err = json.Unmarshal(content, &j)
		if err != nil {
			return nil, err
		}

		journals[i] = j
	}
	

	return journals, nil
}

func unmarshalFoundryJsonFiles(path string, contents []string, licenses []string, noLegacyContent bool) (*foundry.Dataset, error) {
	var dataset foundry.Dataset

	dataset.Journals = make([]foundry.Journal, len(journalFiles))
	packsDir := path + packs
	for i, file := range journalFiles {
		content, err := os.ReadFile(packsDir + file)
		if err != nil {
			fmt.Printf("Error reading journal file %s. Error: %v", packsDir + file, err)
			return nil, err
		}

		var j foundry.Journal
		err = json.Unmarshal(content, &j)
		if err != nil {
			return nil, err
		}

		dataset.Journals[i] = j
	}


	// create <absPath>/packs/<content paths> to walk and walk them using there matching foundry type
	for _, c := range contents {
		for _, val := range contentsToDirs[c] {
			fullpath := packsDir + val
			fmt.Printf("Loading content under %s\n", fullpath)
			switch val {
			case "backgrounds":
				b, err := walkDir[foundry.Background](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Backgrounds = b
				//writeAll(bgs)
			case "ancestries":
				a, err := walkDir[foundry.Ancestry](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Ancestries = a
				//writeAll(as)
			case "ancestryfeatures":
				af, err := walkDir[foundry.Feature](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.AncestryFeatures = af
			case "classfeatures":
				cf, err := walkDir[foundry.Feature](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.ClassFeatures = cf
			case "feats":
				f, err := walkDir[foundry.Feature](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Feats = f
				//writeAll(fs)
			case "classes":
				c, err := walkDir[foundry.Class](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Classes = c
				//writeAll(cs)
			case "equipment":
				e, err := walkDir[foundry.EquipmentEnvelope](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Equipment = e
			case "equipment-effects":
				ee, err := walkDir[foundry.EquipmentEffect](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.EquipmentEffects = ee
			case "feat-effects":
				fe, err := walkDir[foundry.FeatEffect](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.FeatEffects = fe
			case "heritages":
				h, err := walkDir[foundry.Heritage](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Heritages = h
			case "other-effects":
				oe, err := walkDir[foundry.OtherEffect](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.OtherEffects = oe
			case "spell-effects":
				se, err := walkDir[foundry.SpellEffect](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.SpellEffects = se
			case "spells":
				s, err := walkDir[foundry.Spell](fullpath, noLegacyContent, licenses)
				if err != nil {
					return nil, err
				}
				dataset.Spells = s
			default:
				fmt.Printf("%s is not a supported content type right now.", val)
			}
		}
	}

	return &dataset, nil
}

func Build() error {
	return nil
}
