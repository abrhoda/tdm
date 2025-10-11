package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func walkFunc(path string, dirEntry os.DirEntry, err error) error {
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

	var pl map[string]interface{}
	err = json.Unmarshal([]byte(content), &pl)

	fmt.Printf("Json obj with name %s\n", pl["name"].(string))

	return nil
}

func main() {
	abspath := "/home/alexander/repositories/pf2e/packs"
	//relpath := "../pf2e/packs"

	filepath.WalkDir(abspath, walkFunc)

}
