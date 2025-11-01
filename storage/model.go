package storage

type databaseModel interface {
	Validate() error
}

// TODO DELETE
type commonFields struct {
	Name                  string
	Description           string
	GameMasterDescription string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	Traits                []string
	Rules                   string
}

type Trait struct {
	ID   int
	Name string
}

type Boost struct {
	ID   int
	Name string
}

type Proficiency struct {
	ID   int
	Name string
	Rank int
}

type Ancestry struct {
	ID                      int
	Name                    string
	Description             string
	GameMasterDescription   string
	Title                   string
	Remaster                bool
	License                 string
	Rarity                  string
	Traits                  []Trait
	Rules                   string
	Boosts                  map[string][]Boost
	Flaw                    []Boost
	Languages               []string
	AdditionalLanguageCount int
	AdditionalLanguages     []string
	HP                      int
	Reach                   int
	Size                    string
	Speed                   int
	Vision                  string
	AncestryFeatures        []AncestryProperty
}

type AncestryProperty struct {
	ID                    int
	Name                  string
	Description           string
	GameMasterDescription string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	Traits                []Trait
	Rules                   string
	ActionType            string   // is this always "passive"?
	Actions               int      // is this always null/0?
	Category              string   // is this always "ancestryfeature"?
	Level                 int      // is this always 0?
	Prerequisites         []string // is this always empty?
	GrantsLanguages       []string
	GrantsLanguageCount   int
	Senses                []Sense
	SuppressedFeatures    []string      // is this always empty?
	KeyAbilityOptions     []string      // is this always empty?
	Proficiencies         []Proficiency // is this always empty?
}

type Sense struct {
	ID                         int
	Name                       string
	Acuity                     string
	Range                      int
	ElevateIfHasLowLightVision bool // field used to "elevate llv to dv if applied to a entity with existing llv
}

type Skill struct {
	ID   int
	Name string
}

type Background struct {
	ID                    int
	Name                  string
	Description           string
	GameMasterDescription string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	Traits                []Trait
	Rules                   string
	Boosts                map[string]Boost
	Skills                []Skill
	Feat                  GeneralFeat
}

type Prerequisite struct {
	ID    string
	Value string
}

type GeneralFeat struct {
	ID                    int
	Name                  string
	Description           string
	GameMasterDescription string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	Traits                []Trait
	Rules                   string
	action_type           string
	actions               int
	category              string
	level                 int
	Prerequisites         []Prerequisite
	MaxTakable            int
	FrequencyMax          int           // only 1 general feat has this
	FrequencyPeriod       string        // only 1 general feat has this
	Proficiencies         []Proficiency // only 1 general feat grants proficiencies
}
