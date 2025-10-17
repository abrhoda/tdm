package main

import (
	"encoding/json"
	"fmt"
	"github.com/abrhoda/tdm/models"
	"os"
	"path/filepath"
	"strings"
)

const packs = "/packs/"

var contentsToDirs = map[string][]string{
	//"actions": {"actions"},
	"ancestries":  {"ancestries", "ancestryfeatures"},
	"backgrounds": {"backgrounds"},
	//"bestiaries": {.. list them all ..},
	"classes": {"classes", "classfeatures"},
	//"conditions": {"conditions"},
	//"deities": {"deities"},
	"effects": {"other-effects"},
	"equipment": {"equipment", "equipment-effects"},
	"feats":     {"feats", "feat-effects"},
	"heritages": {"heritages"},
	//"hazards":" {"hazards"}
	//"spells": {"spells", "spell-effects"},

	// TODO others to include:
	// hazards, other-effects (this is like aid), deities, conditions, bestiaries, actions
}

var allContents = []string{
	//"actions",
	"ancestries",
	"backgrounds",
	//"bestiaries",
	"classes",
	//"conditions",
	//"deities",
	"equipment",
	"feats",
	//"hazards",
	"heritages",
	//"hero-point-deck",
	"effects",
	//"spells",
}
var allLicenses = []string{"ogl", "orc"}

// TODO out slice should have a capacity to avoid reallocations when adding elements.
func walkDir[T models.FoundryModel](path string, noLegacyContent bool, licenses []string) ([]T, error) {
	out := make([]T, 0)

	err := filepath.WalkDir(path, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("Error for entry %s. Error: %v", path, err)
			fmt.Printf("got error: %v\n", err)
			return err
		}

		if dirEntry.IsDir() || !strings.HasSuffix(dirEntry.Name(), ".json") || dirEntry.Name() == "_folders.json" {
			fmt.Printf("Not processing %s\n", dirEntry.Name())
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Printf("Error reading entry %s. Error: %v", path, err)
			return err
		}

		var data T
		err = json.Unmarshal(content, &data)
		if err != nil {
			return err
		}

		// filter out legacy content if needed.
		if noLegacyContent && data.IsLegacy() {
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
		return nil, err
	}

	return out, nil
}

func buildDataset(path string, contents []string, licenses []string, noLegacyContent bool) error {
	// fix paths with '~' start
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		path = strings.Replace(path, "~", homeDir, 1)
	}

	// ensure we always use the absolute path
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	// create <absPath>/packs/<content paths> to walk and walk them using there matching foundry type
	for _, c := range contents {
		for _, val := range contentsToDirs[c] {
			p := path + packs + val
			fmt.Printf("Loading content under %s\n", p)
			switch val {
			case "backgrounds":
				_, err := walkDir[models.Background](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
				//writeAll(bgs)
			case "ancestries":
				_, err := walkDir[models.Ancestry](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
				//writeAll(as)
			case "ancestryfeatures", "classfeatures", "feats":
				_, err := walkDir[models.Feature](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
				//writeAll(fs)
			case "classes":
				_, err := walkDir[models.Class](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
				//writeAll(cs)
			case "equipment":
				_, err := walkDir[models.EquipmentEnvelope](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
			case "equipment-effects":
				_, err := walkDir[models.EquipmentEffect](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
			case "feat-effects":
				_, err := walkDir[models.FeatEffect](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
			case "heritages":
				_, err := walkDir[models.Heritage](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
			case "other-effects":
				_, err := walkDir[models.OtherEffect](p, noLegacyContent, licenses)
				if err != nil {
					return err
				}
			default:
				fmt.Printf("%s is not a supported content type right now.", val)
			}
		}
	}
	return nil
}

func writeAll[T models.FoundryModel](toWrite []T) {
	for i, item := range toWrite {
		fmt.Printf("%d. %+v\n", i, item)
	}
}
