package storage

type commonFields struct {
	Name string
	Description string
	GameMasterDescription string
	Title string
	Remaster bool
	License string
	Rarity string
	Traits []string
	Rules []any
}

type Ancestry struct {
	Name string
	Description string
	GameMasterDescription string
	Title string
	Remaster bool
	License string
	Rarity string
	Traits []string
	Rules []any
	FirstBoost []string
	SecondBoost []string
	ThirdBoost []string
	Language []string
	AdditionalLanguageCount int
	AdditionalLanguageOptions []string
	Flaws []string
	HP int
	Reach int
	Size string
	Speed int
	Vision string
	AncestryFeatures []AncestryFeature
}

type AncestryFeature struct {
	Name string
	Description string
	GameMasterDescription string
	Title string
	Remaster bool
	License string
	Rarity string
	Traits []string
	Rules []any
	ActionType string
	Actions int
	Category string
	Level int
	Prerequisites []string
	GrantsLanguages []string
	GrantsLanguageCount int
	GrantsLowLightVision bool
	GrantsDarkVisionIfAncestryHasLowLightVision bool
}
