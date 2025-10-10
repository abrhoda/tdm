package main

import (
	"fmt"
	"os"

	"github.com/abrhoda/tdm/internal/pf2e"
)

// type stringList []string
// 
// func (sl stringList) String() string {
// 	return strings.Join(sl, ",")
// }
// 
// // splits on space if there is one, otherwise splits on comma, else returns slice of original string
// func (sl *stringList) Set(value string) error {
// 	if strings.Contains(value, " ") {
// 		*sl = strings.Split(value, " ")
// 	} else {
// 		*sl = strings.Split(value, ",")
// 	}
// 	return nil
// }

func main() {
	// var contents stringList
	// var path, license string
	// var isCommercial, noLegacy, shouldUpdate bool

	// fs := flag.NewFlagSet("build", flag.ExitOnError)
	// fs.Var(&contents, "contents", "A list of specific content to build. Options are all, backgrounds, classes, equipment, feats, heritages")
	// fs.BoolVar(&noLegacy, "no-legacy", true, "A flag to denote if content before the remaster should be included in the dataset.")
	// fs.BoolVar(&isCommercial, "commercial", false, "A flag to denote if dataset being used will be used for commercial purposes (applies a filter to remove Paizo content).")
	// fs.BoolVar(&shouldUpdate, "update", false, "A flag to denote if the foundryvtt/pf2e repo locally should be updated (git pull) before building the dataset.")
	// fs.StringVar(&license, "license", "all", "Which license content used to build the dataset should have. Possible values are: OGL, ORC, or all")


	// if len(os.Args) < 4 {
	// 	// TODO should look to see if theres a help command in these flags and print a usage message
	// 	fmt.Println("tdm pf2e build [options] path_to_pf2e_repo")
	// 	os.Exit(1)
	// }

	// if os.Args[1] != "pf2e" {
	// 	fmt.Printf("Only command that is supports is \n", len(os.Args), os.Args[1:])
	// 	os.Exit(1)
	// }

	// fs.Parse(os.Args[3:])

	// if strings.ToLower(license) != "all" || !slices.Contains(pf2e.Licenses, strings.ToLower(license)) {
	// 	fmt.Printf("%s is not a valid license\n", license)
	// 	os.Exit(1)
	// }

	// if contents == nil {
	// 	contents = pf2e.Contents
	// } else {
	// 	all := false
	// 	for _, s := range contents {
	// 		if strings.EqualFold(s, "all") {
	// 			fmt.Println("found all")
	// 			all = true
	// 			continue
	// 		}
	// 		if !slices.Contains(pf2e.Contents, strings.ToLower(s)) {
	// 			fmt.Printf("%s is not a valid content type\n", s)
	// 			os.Exit(1)
	// 		}
	// 		
	// 	}
	// 	if all {
	// 		contents = pf2e.Contents 
	// 	}
	// }

	// path = os.Args[len(os.Args)-1]

	path := "~/repository/pf2e"
	contents := pf2e.Contents
	licenses := pf2e.Licenses
	noLegacy := false
	
	err := pf2e.BuildDataset(path, contents, licenses, noLegacy)
	if err != nil {
		fmt.Printf("Error when running build: %v\n", err)
		os.Exit(1)
	}
}
