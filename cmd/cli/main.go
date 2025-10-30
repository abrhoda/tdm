package main

import (
	"fmt"
	"os"
	"github.com/abrhoda/tdm/internal/handlers"
)

func main() {
	path := "~/repositories/pf2e"
	contents := []string{"ancestries", "backgrounds", "classes", "equipment", "feats", "heritages", "effects", "spells"}
	licenses := []string{"OGL", "ORC"}
	noLegacy := false

	err := handlers.Build(path, contents, licenses, noLegacy)
	if err != nil {
		fmt.Printf("Error when running build: %v\n", err)
		os.Exit(1)
	}
}
