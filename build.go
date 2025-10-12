package main

import (
	"fmt"
	"os"
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

// try without worker pool first?
func walkDir[T foundryType](path string, licenses []string, noLegacy bool) ([]T, error) {
	// TODO should probably set a default capacity to avoid resizing.
	vals := make([]T, 0)

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
					
					fmt.Println(string(content))

					return nil
				})


	if err != nil {
		return nil, err
	}
	
	return vals, nil
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

	fmt.Printf("Root path to use is %s\n", path)

	// create <absPath>/packs/<content paths> to walk and walk them using there matching foundry type
	for _, c := range contents {
		for _, val := range contentsToDirs[c] {

			p := path + packs + val
			fmt.Printf("Loading content under %s\n", p)
			switch c {
				case "backgrounds":
					walkDir[background](p, licenses, noLegacy)
				default:
					return fmt.Errorf("%s is not a supported content type right now.", c)
			}
		}
	}
	return nil
}
