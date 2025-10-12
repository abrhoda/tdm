package main

import (
	"fmt"
	"os"
	"encoding/json"
  "path/filepath"
	"strings"
)

const packs = "/packs/"

var contentsToDirs = map[string][]string {
	//"actions": {"actions"},
	//"ancestries": {"ancestries", "ancestryfeatures"},
	"backgrounds": {"backgrounds"},
	//"bestiaries": {.. list them all ..},
	//"classes": {"classes", "classfeatures"},
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
	//"ancestries", 
	"backgrounds",
	//"bestiaries",
	//"classes",
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
func walkDir[T foundryType](path string) ([]T, error) {
	// TODO should probably set a default capacity to avoid resizing.
	out := make([]T, 0)

	err := filepath.WalkDir(path, func(path string, dirEntry os.DirEntry, err error) error {
					if err != nil {
						fmt.Printf("Error for entry %s. Error: %v", path, err)
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
//func removeLegacyContent[T foundryType](items []T) []T {
//	out := make([]T, 0)
//	for _, item := range items {
//		if item.System.Publication.Remaster {
//			out = append(out, item)
//		}
//	}
//
//	return out
//}

// TODO out slice should have a capacity to avoid reallocations when adding elements.
//func removeByLicense[T foundryType](items []T, license string) []T {
//	out := make([]T, 0)
//	for _, item := range items {
//		if item.System.Publication.License != license {
//			out = append(out, item)
//		}
//	}
//	return out
//}

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
			switch c {
				case "backgrounds":
					bgs, err := walkDir[background](p)
					if err != nil {
						return err
					}

					for i, bg := range bgs {
						fmt.Printf("%d. %s\n", i, bg.Name)
					}
				default:
					return fmt.Errorf("%s is not a supported content type right now.", c)
			}
		}
	}
	return nil
}
