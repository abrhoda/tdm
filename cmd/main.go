package main

import (
	"fmt"
	"os"
	"github.com/abrhoda/tdm"
)

func main() {
	updateFoundry := false
	foundryDirectory := "~/repositories/pf2e"
	contents := []tdm.ContentOption{tdm.Ancestries, tdm.Backgrounds, tdm.Classes, tdm.Equipment, tdm.Feats, tdm.Heritages, tdm.Effects, tdm.Spells}
	includeLegacy := true
	licenses := []tdm.LicenseOption{tdm.OpenGamingLicense, tdm.OpenRPGCreativeLicense}

	
	_, err := tdm.NewInMemoryConfig(updateFoundry, foundryDirectory, contents, includeLegacy, licenses)
	if err != nil {
		fmt.Printf("Error when building config: %v\n", err)
		os.Exit(1)
	}

	

}
