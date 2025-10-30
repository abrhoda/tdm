package tdm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type ContentOption string
const (
	Ancestries ContentOption = "ancestries"
	BackGrounds ContentOption = "backgrounds"
	Classes ContentOption = "classes"
	Equipment ContentOption = "equipment"
	Feats ContentOption = "feats"
	Heritages ContentOption = "heritages"
	Effects ContentOption = "effects"
	Spells ContentOption = "spells"
)
const maxContentLength = 8

type LicenseOption string
const (
	OpenGamingLicense LicenseOption = "OGL"
	OpenRPGCreativeLicense LicenseOption = "ORC"
)

type OutputOption int
const (
	InMemory OutputOption = 1
	JSON OutputOption = 2
	PostgreSQL OutputOption = 3
)

type configuration struct {
	updateFoundry bool
	outputType OutputOption

	// outputType == InMemory has no configuration.

	// outputType == Json
	outputDirectory string

	// outputType == Postgresql
	username string
	password string
	host string
	port int

	// content path options 
	foundryDirectory string
	content []ContentOption

	// filter options
	includeLegacy bool
	licenses []LicenseOption
}

func NewInMemoryConfig(updateFoundry bool, foundryDirectory string, content []ContentOption, includeLegacy bool, licenses []LicenseOption) (*configuration, error) {
	return NewConfig(updateFoundry, InMemory, "", "", "", "", 0, foundryDirectory, content, includeLegacy, licenses)
}

func NewJsonConfig(updateFoundry bool, outputDirectory string, foundryDirectory string, content []ContentOption, includeLegacy bool, licenses []LicenseOption) (*configuration, error) {
	return NewConfig(updateFoundry, JSON, outputDirectory, "", "", "", 0, foundryDirectory, content, includeLegacy, licenses)
}

func NewPostgreSQLConfig(updateFoundry bool, username, string, password string, host string, port int, foundryDirectory string, content []ContentOption, includeLegacy bool, licenses []LicenseOption) (*configuration, error) {
	return NewConfig(updateFoundry, PostgreSQL, "", username, password, host, port, foundryDirectory, content, includeLegacy, licenses)
}

// providing options not needed for the OutputOption chosen will be ignored. Example: passing in a username, password, host, and port for a postgresql database while setting OutputOption == InMemory will cause the username, password, host, and password to not be set on the config struct.
func NewConfig(updateFoundry bool, outputType OutputOption, outputDirectory string, username string, password string, host string, port int, foundryDirectory string, content []ContentOption, includeLegacy bool, licenses []LicenseOption) (*configuration, error) {
	if strings.TrimSpace(foundryDirectory) == "" {
		return nil, fmt.Errorf("foundryDirectory cannot be blank or empty.")
	}

	foundryPath, err := normalizePath(foundryDirectory)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, fmt.Errorf("content list cannot be empty.")
	}

	// strip duplicates from content list
	contentSet :=  make(map[ContentOption]struct{}, maxContentLength)
	for _, c := range content {
		contentSet[c] = struct{}{}
	}
	// set back to content var
	content = make([]ContentOption, len(contentSet))
	for k := range contentSet {
		content = append(content, k)
	}

	if len(licenses) == 0 || len(licenses) > 2 {
		return nil, fmt.Errorf("licenses list only allows values OGL or ORC.")
	}

	if len(licenses) == 2 && licenses[0] == licenses[1] {
		return nil, fmt.Errorf("licenses list cannot contain duplicates.")
	}

	var config *configuration
	switch outputType {
		case InMemory:
			config = &configuration{updateFoundry: updateFoundry, outputType: outputType, foundryDirectory: foundryPath, content: content, includeLegacy: includeLegacy, licenses: licenses}
		case JSON:
			if strings.TrimSpace(outputDirectory) == "" {
				return nil, fmt.Errorf("outputDirectory cannot be blank or empty.")
			}
			outputPath, err := normalizePath(outputDirectory)
			if err != nil {
				return nil, err
			}
			config = &configuration{updateFoundry: updateFoundry, outputType: outputType, foundryDirectory: foundryPath, outputDirectory: outputPath, content: content, includeLegacy: includeLegacy, licenses: licenses}
		case PostgreSQL:
			if strings.TrimSpace(username) == "" {
				return nil, fmt.Errorf("username cannot be blank or empty.")
			}
			if strings.TrimSpace(password) == "" {
				return nil, fmt.Errorf("password cannot be blank or empty.")
			}
			if strings.TrimSpace(host) == "" {
				return nil, fmt.Errorf("host cannot be blank or empty.")
			}
			if port < 0 || port > 65536 {
				return nil, fmt.Errorf("port must be between 0-65536.")
			}
			config = &configuration{updateFoundry: updateFoundry, outputType: outputType, username: username, password: password, host: host, port: port, foundryDirectory: foundryPath, content: content, includeLegacy: includeLegacy, licenses: licenses}
	}

	return config, nil
}

func normalizePath(path string) (string, error) {
	// fix paths with '~' start
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = strings.Replace(path, "~", homeDir, 1)
	}

	// ensure we always use the absolute path
	path, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	return path, nil
}
