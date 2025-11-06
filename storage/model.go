package storage

type Trait struct {
	ID   int
	Name string
}

// a "tag" is not a "trait"! Example: tag = "oracle mastery" whereas trait = "oracle"
type Tag struct {
	ID   int
	Name string
}

type Boost struct {
	ID   int
	Name string
}

// TODO Proficiency struct is doing a lot of heavy lifting.
type Proficiency struct {
	ID   int
	Name string
	Rank string
	Type string // can be attack, defense, or skill.
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
	Rules                 string
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
	Rules                 string
	Boosts                map[string]Boost
	TrainedSkills         []Proficiency
	Feat                  GeneralFeat
}

type Prerequisite struct {
	ID    string
	Value string
}

type SkillFeat struct {
	ID                    int
	Name                  string
	Description           string
	GameMasterDescription string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	Traits                []Trait
	Rules                 string
	ActionType            string
	Actions               int
	Category              string
	Level                 int
	Prerequisites         []Prerequisite
	MaxTakable            int
	FrequencyMax          int
	FrequencyPeriod       string
	Proficiencies         []Proficiency
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
	Rules                 string
	ActionType            string
	Actions               int
	Category              string
	Level                 int
	Prerequisites         []Prerequisite
	MaxTakable            int
	FrequencyMax          int           // only 1 general feat has this
	FrequencyPeriod       string        // only 1 general feat has this
	Proficiencies         []Proficiency // only 1 general feat grants proficiencies
}

type FeatLevel struct {
	ID    int
	Level int
	Type  string
}

type SkillIncreaseLevel struct {
	ID int
	Level int
}

type KeyAbility struct {
	ID    int
	Value string
}

type ClassProperty struct {
	ID                    int
	Name                  string
	Description           string
	GameMasterDescription string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	Traits                []Trait
	Rules                 string
	ActionType            string
	Actions               int
	Category              string
	Level                 int
	Tags                  []Tag
	KeyAbilities          []KeyAbility
	Proficiencies         []Proficiency
}

type Class struct {
	ID                    int
	Name                  string
	Description           string
	Title                 string
	Remaster              bool
	License               string
	Rarity                string
	FeatLevels            []FeatLevel
	SkillIncreaseLevels   []SkillIncreaseLevel
	AttackProficiencies   []Proficiency
	DefenseProficiencies  []Proficiency
	HP                    int
	ClassProperties []ClassProperty
	KeyAbilities             []KeyAbility
	Perception               string
	SavingThrowProficiencies []Proficiency
	Spellcasting             string
	AdditionalTrainedSkills  int
	TrainedSkills            []Proficiency
}
