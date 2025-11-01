package tdm

import (
	"encoding/json"
	"fmt"
	"github.com/abrhoda/tdm/internal"
	"github.com/abrhoda/tdm/internal/foundry"
	"github.com/abrhoda/tdm/storage"
	"os"
	"path/filepath"
	"slices"
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
func unmarshalToFoundryModels[T foundry.FoundryModel](fullpath string) ([]T, error) {
	out := make([]T, 0)

	err := filepath.WalkDir(fullpath, func(fullpath string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error for entry %s. Error: %v", fullpath, err)
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

		var data T
		err = json.Unmarshal(content, &data)
		if err != nil {
			return err
		}

		out = append(out, data)
		return nil
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return out, nil
}

func unmarshalFoundryJournals(path string) ([]foundry.Journal, error) {
	packsDir := path + packs
	journals := make([]foundry.Journal, len(journalFiles))
	for i, file := range journalFiles {
		content, err := os.ReadFile(packsDir + file)
		if err != nil {
			fmt.Printf("Error reading journal file %s. Error: %v", packsDir+file, err)
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

// unmarshal should just take target type T as a generic, a dir path, and return []T. this is easier to split into different go routines and means
// that each []T can be unmarshalled, sanitized, filtered, and converted to a database model independently.
func filterFoundryModels[T foundry.FoundryModel](entities []T, licenses []LicenseOption, includeLegacy bool) []T {
	// no filtering to be done.
	if includeLegacy && slices.Contains(licenses, OpenGamingLicense) && slices.Contains(licenses, OpenRPGCreativeLicense) {
		return entities
	}

	out := make([]T, len(entities))
	for _, e := range entities {
		if !includeLegacy && e.IsLegacy() {
			continue
		}
		for _, l := range licenses {
			if e.HasProvidedLicense(string(l)) {
				out = append(out, e)
			}
		}
	}

	return out
}

// TODO implement
func processFoundryModel[T foundry.FoundryModel](entities []T, stripHtml bool, stripDiceExpressions bool, stripTags bool) ([]T, error) {
	// strip html from description and gm description
	// check description and gm description `@Check[<stat>|dc:<number>]`, `@Damage[XdXX[<type>]]`, `@Embed[...]`, `@UUID[...]` or `Compendium.pf2e...` tags and hydrate
	// check description and gm for `[[/r ...]]{...}` for dice expressions

	return entities, nil
}

func Build(cfg configuration) error {
	// route to correct foundry model type and put slice into `dataset` container
	var inMemoryDatastore storage.InMemoryDatastore
	for _, c := range cfg.content {
		for _, val := range contentsToDirs[string(c)] {
			fullpath := cfg.foundryDirectory + packs + val
			fmt.Printf("Loading content under %s\n", fullpath)
			switch val {
			case "backgrounds":
				b, err := unmarshalToFoundryModels[foundry.Background](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(b, cfg.licenses, cfg.includeLegacy)
				//writeAll(bgs)
			case "ancestries":
				a, err := unmarshalToFoundryModels[foundry.Ancestry](fullpath)
				if err != nil {
					return err
				}
				filtered := filterFoundryModels(a, cfg.licenses, cfg.includeLegacy)
				storageAncestries := make([]storage.Ancestry, len(filtered))
				for i, filtered := range filtered {
					storageAncestries[i], err = internal.ConvertAncestry(filtered)
					if err != nil {
						return err
					}
				}
				inMemoryDatastore.Ancestries = storageAncestries
				//writeAll(as)
			case "ancestryfeatures":
				af, err := unmarshalToFoundryModels[foundry.Feature](fullpath)
				if err != nil {
					return err
				}
				filtered := filterFoundryModels(af, cfg.licenses, cfg.includeLegacy)
				storageAncestryProperties := make([]storage.AncestryProperty, len(filtered))
				for i, filtered := range filtered {
					storageAncestryProperties[i], err = internal.ConvertAncestryProperty(filtered)
					if err != nil {
						return err
					}
				}
				inMemoryDatastore.AncestryProperties = storageAncestryProperties
			case "classfeatures":
				cf, err := unmarshalToFoundryModels[foundry.Feature](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(cf, cfg.licenses, cfg.includeLegacy)
			case "feats":
				f, err := unmarshalToFoundryModels[foundry.Feature](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(f, cfg.licenses, cfg.includeLegacy)
				//writeAll(fs)
			case "classes":
				c, err := unmarshalToFoundryModels[foundry.Class](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(c, cfg.licenses, cfg.includeLegacy)
				//writeAll(cs)
			case "equipment":
				e, err := unmarshalToFoundryModels[foundry.EquipmentEnvelope](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(e, cfg.licenses, cfg.includeLegacy)
			case "equipment-effects":
				ee, err := unmarshalToFoundryModels[foundry.EquipmentEffect](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(ee, cfg.licenses, cfg.includeLegacy)
			case "feat-effects":
				fe, err := unmarshalToFoundryModels[foundry.FeatEffect](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(fe, cfg.licenses, cfg.includeLegacy)
			case "heritages":
				h, err := unmarshalToFoundryModels[foundry.Heritage](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(h, cfg.licenses, cfg.includeLegacy)
			case "other-effects":
				oe, err := unmarshalToFoundryModels[foundry.OtherEffect](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(oe, cfg.licenses, cfg.includeLegacy)
			case "spell-effects":
				se, err := unmarshalToFoundryModels[foundry.SpellEffect](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(se, cfg.licenses, cfg.includeLegacy)
			case "spells":
				s, err := unmarshalToFoundryModels[foundry.Spell](fullpath)
				if err != nil {
					return err
				}
				_ = filterFoundryModels(s, cfg.licenses, cfg.includeLegacy)
			default:
				fmt.Printf("%s is not a supported content type right now.", val)
			}
		}
	}

	return nil
}
