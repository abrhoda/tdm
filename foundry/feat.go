package foundry

import (
	"github.com/abrhoda/tdm/pkg/storage"
)

// note on features
// ancestryfeature and classfeature as "level 0" things that an ancestry or class provide. These are NOT picked by the player
// ancestry and class are feats that are actually picked by the player.
type Feature struct {
	Name   string        `json:"name"`
	System featureSystem `json:"system"`
}

func (f Feature) IsLegacy() bool {
	return !f.System.Publication.Remaster
}

func (f Feature) HasProvidedLicense(license string) bool {
	return f.System.Publication.License == license
}

type featureSystem struct {
	commonSystem                    // description, publication, traits, and rules
	ActionType    valueNode[string] `json:"actionType"`
	Actions       valueNode[int]    `json:"actions"`
	Category      string            `json:"category"` // classfeature vs class
	Level         valueNode[int]    `json:"level"`
	Prerequisites prerequisites     `json:"prerequisites"`
	MaxTakable    int               `json:"maxTakable,omitempty"`
	Frequency     frequency         `json:"frequency"`
	SubFeatures   subFeatures       `json:"subfeatures"`
}

type frequency struct {
	Max   int    `json:"max"`
	Per   string `json:"per"`
	Value int    `json:"value,omitempty"`
}

type prerequisites struct {
	Value []valueNode[string] `json:"value"`
}


type subFeatures struct {
	Proficiencies      map[string]map[string]int `json:"proficiences"` // top map will have "attribute" that is effected and nested map should have 1 key of "rank" and an int to tell the rank. other keys can be ignored in nested map.
	Senses             map[string]sense          `json:"senses"`
	SuppressedFeatures []string                  `json:"suppressedFeatures"`
	Languages          subFeatureLanguages       `json:"languages"`
}

type sense struct {
	Acuity  string          `json:"acuity"`
	Range   int             `json:"range"`
	Special map[string]bool `json:"special"`
}

type subFeatureLanguages struct {
	Granted []string `json:"granted"`
	Slots   int      `json:"slots"`
}

func (f Feature) toDatabaseModel() (storage.AncestryFeature, error) {
	prereqs := make([]string, len(f.System.Prerequisites.Value))
	for _, vnode := range f.System.Prerequisites.Value {
		prereqs = append(prereqs, vnode.Value)
	}

	af := storage.AncestryFeature{
		Name: f.Name,
	  Description: f.System.Description.Value,
	  GameMasterDescription: f.System.Description.GameMasterDescription,
		Title: f.System.Publication.Title,
		Remaster:f.System.Publication.Remaster,
		License:f.System.Publication.License,
		Rarity: f.System.Traits.Rarity,
		Traits: f.System.Traits.Value,
		Rules: f.System.Rules,
		ActionType: f.System.ActionType.Value,
		Actions: f.System.Actions.Value,
		Category: f.System.Category,
		Level: f.System.Level.Value,
		Prerequisites: prereqs,
		GrantsLanguages: f.System.SubFeatures.Languages.Granted,
		GrantsLanguageCount: f.System.SubFeatures.Languages.Slots,

		// TODO make these not constants
		GrantsLowLightVision: false,
		GrantsDarkVisionIfAncestryHasLowLightVision: false,
	}
	
	return af, nil
}
