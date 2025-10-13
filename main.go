package main

import (
	"fmt"
	"os"
)

func main() {
	path := "~/repositories/pf2e"
	contents := allContents
	licenses := allLicenses
	noLegacy := false

	err := buildDataset(path, contents, licenses, noLegacy)
	if err != nil {
		fmt.Printf("Error when running build: %v\n", err)
		os.Exit(1)
	}
}
