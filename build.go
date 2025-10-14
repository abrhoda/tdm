package main

import (
	"encoding/json"
	"fmt"
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
	//"effects": {"other-effects"}
	//"equipment": {"equipment", "equipment-effects"},
	//"feats": {"feats", "feat-effects"},
	//"heritages": {"heritages"},
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
	//"equipment",
	//"feats",
	//"hazards",
	//"heritages",
	//"effects",
	//"spells",
}
var allLicenses = []string{"ogl", "orc"}

// TODO items:
// 1. make vals start with a capacity to reduce the amount times append(vals, T) has to reallocate underlying memory
// 2. func could take `licenses` and `noLegacy` params to filter content BEFORE unmarshalling into T
func walkDir[T foundryModel](path string) ([]T, error) {
	// TODO should probably set a default capacity to avoid resizing.
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
			return nil
		}

		out = append(out, data)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return out, nil
}

// TODO out slice should have a capacity to avoid reallocations when adding elements.
func removeLegacyContent[T foundryModel](items []T) []T {
	out := make([]T, 0)
	for _, item := range items {
		if !item.IsLegacy() {
			out = append(out, item)
		}
	}

	return out
}

// TODO out slice should have a capacity to avoid reallocations when adding elements.
func removeForLicense[T foundryModel](items []T, license string) []T {
	out := make([]T, 0)
	for _, item := range items {
		if !item.HasProvidedLicense(license) {
			out = append(out, item)
		}
	}
	return out
}

func buildDataset(path string, contents []string, licenses []string, noLegacy bool) error {
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
				_, err := walkDir[background](p)
				if err != nil {
					return err
				}
				//writeAll(bgs)
			case "ancestries":
				_, err := walkDir[ancestry](p)
				if err != nil {
					return err
				}
				//writeAll(as)
			case "ancestryfeatures", "classfeatures":
				_, err := walkDir[feature](p)
				if err != nil {
					return err
				}
				//writeAll(fs)
			case "classes":
				_, err := walkDir[class](p)
				if err != nil {
					return err
				}
				//writeAll(cs)
			default:
				return fmt.Errorf("%s is not a supported content type right now.", c)
			}
		}
	}
	return nil
}

func writeAll[T foundryModel](toWrite []T) {
	for i, item := range toWrite {
		fmt.Printf("%d. %+v\n", i, item)
	}
}
