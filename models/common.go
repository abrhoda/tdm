package models

type foundryType interface {
	// types
	Ancestry |
		//ancestryFeature |
		Background |
		Class |
		//		classFeature |
		//EquipmentEffect |
		Equipment |
		Feature
		//FeatureEffect |
		//Heritage
}

type filterable interface {
	// funcs
	IsLegacy() bool
	HasProvidedLicense(string) bool
}

type FoundryModel interface {
	foundryType
	filterable
}

// common shared types
type commonSystem struct {
	Description description `json:"description"`
	Publication publication `json:"publication"`
	Traits      traits      `json:"traits"`
	Rules       []any       `json:"rules"`
}

type publication struct {
	Title    string `json:"title"`
	Remaster bool   `json:"remaster"`
	License  string `json:"license"`
}

type description struct {
	Value                 string `json:"value,omitempty"`
	GameMasterDescription string `json:"gm,omitempty"`
}

type traits struct {
	Rarity    string   `json:"rarity"`
	Value     []string `json:"value"`
	OtherTags []string `json:"otherTags,omitempty"`
}

// to pull out all those annoyingly nested objs with just a value key
type valueNode[T any] struct {
	Value T `json:"value"`
}

type boosts struct {
	First  valueNode[[]string] `json:"0"`
	Second valueNode[[]string] `json:"1"`
	Third  valueNode[[]string] `json:"2"`
}

type systemItem struct {
	Level int    `json:"level"`
	Name  string `json:"name"`
	UUID  string `json:"uuid"`
}
