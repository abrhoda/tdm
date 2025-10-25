package storage


// TODO DELETE
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

type Proficiency struct {
	ID int
	Name string
	Rank int
}

type Ancestry struct {
	ID int
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
	ID int
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
	Senses []Sense
	SuppressedFeatures []string
	KeyAbilityOptions []string
	Proficiencies []Proficiency
}

type Sense struct {
	ID int
	Name string
	Acuity string
	Range int
	ElevateIfHasLowLightVision bool // field used to "elevate llv to dv if applied to a entity with existing llv
}
