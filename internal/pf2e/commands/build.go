package commands

import (
//	"fmt"
	"os"
  "path/filepath"
	"strings"
	"github.com/abrhoda/tdm/internal/pf2e/foundry"
)

const packs = "/packs/"
const ancestries = "ancestries"
const ancestryFeatures = "ancestryfeatures"
const backgrounds = "backgrounds"
const classes = "classes"
const classFeatures = "classfeatures"
const equipment = "equipment"
const equipmentEffects = "equipment-effects"
const feats = "feats"
const featEffects = "feat-effects"
const heritages = "heritages"

var contentsToDirs = map[string][]string {
	ancestries: {ancestries, ancestryFeatures},
	backgrounds: {backgrounds},
	classes: {classes, classFeatures},
	equipment: {equipment, equipmentEffects},
	feats: {feats, featEffects},
	heritages: {heritages},
}

var Contents = []string{ancestries, backgrounds, classes, equipment, feats, heritages}
var Licenses = []string{"ogl", "orc"}

// try without worker pool first?
func walkDir[T foundry.ModelType](path string) []T {
	
}

func BuildDataset(path string, contents []string, license []string, noLegacy bool) error {
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

	// create <absPath>/packs/content paths to walk
	dirsToWalk := make([]string, 0, len(contents)*2)
	for _, c := range contents {
		for _, val := range contentsToDirs[c] {
			p := path + packs + val
			dirsToWalk = append(dirsToWalk, p)
		}
	}
	
	numWorkers := len(dirsToWalk)

	// TODO iterate dirs.
	// thinking to make slices to hold the results and use a channel per dirsToWalk value.
	return nil
}
