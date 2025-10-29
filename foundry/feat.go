package foundry

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
	KeyOptions []string `json:"keyOptions"`
	Proficiencies      map[string]map[string]int `json:"proficiences"` // top map will have "attribute" that is effected and nested map should have 1 key of "rank" and an int to tell the rank. other keys can be ignored in nested map.
	Senses             map[string]Sense          `json:"senses"`
	SuppressedFeatures []string                  `json:"suppressedFeatures"`
	Languages          subFeatureLanguages       `json:"languages"`
}

type Sense struct {
	Acuity  string          `json:"acuity"`
	Range   int             `json:"range"`
	Special map[string]bool `json:"special"`
}

type subFeatureLanguages struct {
	Granted []string `json:"granted"`
	Slots   int      `json:"slots"`
}
