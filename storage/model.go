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

type FeatLevel struct {
	ID    int
	Level int
	Type  string
}

type SkillIncreaseLevel struct {
	ID    int
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
	ID                       int
	Name                     string
	Description              string
	Title                    string
	Remaster                 bool
	License                  string
	Rarity                   string
	FeatLevels               []FeatLevel
	SkillIncreaseLevels      []SkillIncreaseLevel
	AttackProficiencies      []Proficiency
	DefenseProficiencies     []Proficiency
	HP                       int
	ClassProperties          []ClassProperty
	KeyAbilities             []KeyAbility
	Perception               string
	SavingThrowProficiencies []Proficiency
	Spellcasting             string
	AdditionalTrainedSkills  int
	TrainedSkills            []Proficiency
}

// TODO placeholder for now.
type FeatEffect struct{}

type AncestryFeat struct {
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
	MaxTakable            int
	FeatEffect            FeatEffect
}

type BonusFeat struct {
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
}

type ClassFeat struct {
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
	FeatEffect            FeatEffect
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
	FeatEffect            FeatEffect // only 1 skill feat has a `selfEffect` key
}

type CraftableAsItem struct {
	ID    int
	Value string
}

type Ammo struct {
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
	Level                 int
	BaseItem              string
	Bulk                  float64
	CraftableAs           []CraftableAsItem
	PricePer              int
	PriceInCopper         int
	Quantity              int
	Size                  string
	MaxUses               int
	AutoDestoryOnUse      bool
}

type ArmorRune struct {
	ID   int
	Name string
}

type Armor struct {
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
	Level                 int
	BaseItem              string
	Bulk                  float64
	HP                    int
	Hardness              int // TODO check if this is always 0
	PricePer              int // TODO check if this is always 0 (as in doesnt exist on data)
	PriceInCopper         int
	Quantity              int // TODO check if this is always 1
	Size                  string
	ACBonus               int
	Category              string
	Group                 string
	CheckPenalty          int
	SpeedPenalty          int
	Strength              int
	Potency               int
	Resilient             int
	Runes                 []ArmorRune
}

type Backpack struct{
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
	Level                 int
	BaseItem              string
	Bulk                  float64
	BulkIgnored float64
	Capacity int
	HeldOrStowed float64
	Hardness              int // TODO check if this is always 0
	PricePer              int // TODO check if this is always 0 (as in doesnt exist on data)
	PriceInCopper         int
	Quantity              int // TODO check if this is always 1
	Collapsed bool // TODO what is this field?
	Stowing bool // TODO rename this. Stowing == "can stow items in this container"?
	Usage string // is this always worn/wornpackback
}
type Consumable struct{
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
	Level                 int
	BaseItem string
	Bulk float64
	MaxHP int // TODO is this always 9 for cons?
	PriceInCopper int
	Hardness int // TODO is this always 0 for cons?
	Quantity int // TODO is this always 1?
	Size string
	Category string
	DamageDiceCount int
	DamageDiceType string // TODO make this an enum of d4, d6, d8, d10, d12, or "" (if "", put flat)
	DamageType string // TODO make this an enum
	PersistentDamageDiceCount int
	PersistentDamageDiceType string // TODO make this the same enum as DamageDiceType
	PersistentDamageType string // TODO make this an enum like damageType
	MaxUses               int
	AutoDestoryOnUse      bool
	StackGroup string
	Usage string
	CanBeAmmom bool
}
type Equipment struct{}
type Kit struct{}
type Shield struct{}
type Treasure struct{}
type Weapon struct{}
