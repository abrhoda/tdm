package main

import (
	"encoding/json"
	"fmt"
)

// main foundry types
type ancestry struct {
	Name   string         `json:"name"`
	System ancestrySystem `json:"system"`
}

func (a ancestry) IsLegacy() bool {
	return !a.System.Publication.Remaster
}

func (a ancestry) HasProvidedLicense(license string) bool {
	return a.System.Publication.License == license
}

type background struct {
	Name   string           `json:"name"`
	System backgroundSystem `json:"system"`
}

func (bg background) IsLegacy() bool {
	return !bg.System.Publication.Remaster
}

func (bg background) HasProvidedLicense(license string) bool {
	return bg.System.Publication.License == license
}

type class struct {
	Name   string      `json:"name"`
	System classSystem `json:"system"`
}

func (c class) IsLegacy() bool {
	return !c.System.Publication.Remaster
}

func (c class) HasProvidedLicense(license string) bool {
	return c.System.Publication.License == license
}

// not a class feat in the trdaitional sense but is instead a "thing" granted by the class.
//type classFeature struct {
//	Name string
//}
//
//func (cf classFeature) IsLegacy() bool {
//	return !cf.System.Publication.Remaster
//}
//
//func (cf classFeature) HasProvidedLicense(license string) bool {
//	return cf.System.Publication.License == license
//}

type equipmentEffect struct {
	Name string
}

// equipmentSubtype container that holds the actual subtype, all subtypes, and equipmentType constraint interface
type equipmentSubtype interface {
	weapon
}

type equipment struct {
	Subtype any // any could be the equipmentType constraint?
}

type weapon struct {
	Name string `json:"name"`
	System weaponSystem `json:"system"`
}

func (e *equipment) UnmarshalJSON(b []byte) error {
	var temp struct {
		Type string `json:"type"`
	}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}
	
	switch temp.Type {
	case "weapon":
		var weapon weapon
		err = json.Unmarshal(b, &weapon)
		if err != nil {
			return err
		}
		e.Subtype = weapon
	default:
		return fmt.Errorf("Unknown equipment type: %s", temp.Type)
	}

	return nil
}

type feature struct {
	Name   string        `json:"name"`
	System featureSystem `json:"system"`
}

func (f feature) IsLegacy() bool {
	return !f.System.Publication.Remaster
}

func (f feature) HasProvidedLicense(license string) bool {
	return f.System.Publication.License == license
}

type featureEffect struct {
	Name string
}

type heritage struct {
	Name string
}

type foundryType interface {
	// types
	ancestry |
		//ancestryFeature |
		background |
		class |
		//		classFeature |
		equipmentEffect |
		equipment |
		feature |
		featureEffect |
		heritage
}

type filterable interface {
	// funcs
	IsLegacy() bool
	HasProvidedLicense(string) bool
}

type foundryModel interface {
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

// feature specific
type featureSystem struct {
	commonSystem // description, publication, traits, and rules
	ActionType    valueNode[string] `json:"actionType"`
	Actions       valueNode[string] `json:"actions"`
	Category      string            `json:"category"`
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

// ancestry specific
type additionalLanguages struct {
	Count  int      `json:"count"`
	Custom string   `json:"custom"`
	Value  []string `json:"value"`
}

type languages struct {
	Custom string   `json:"custom"`
	Value  []string `json:"value"`
}

type ancestrySystem struct {
	commonSystem // description, publication, traits, and rules
	AdditionalLanguages additionalLanguages   `json:"additionalLanguages"`
	Boosts              boosts                `json:"boosts"`
	Items               map[string]systemItem `json:"items"`
	Flaws               boosts                `json:"flaws"`
	HP                  int                   `json:"hp"`
	Size                string                `json:"size"`
	Reach               int                   `json:"reach"`
	Hands               int                   `json:"hands"`
	Speed               int                   `json:"spped"`
	Languages           languages             `json:"languages"`
	Vision              string                `json:"vision"`
}

// background specific
type backgroundTrainedSkills struct {
	Custom string   `json:"custom,omitempty"`
	Lore   []string `json:"lore"`
	Value  []string `json:"value"`
}

type backgroundSystem struct {
	commonSystem // description, publication, traits, and rules
	Boosts        boosts                  `json:"boosts"`
	TrainedSkills backgroundTrainedSkills `json:"trainedSkills"`
	Items         map[string]systemItem   `json:"items"`
}

// class specific
type classSystem struct {
	commonSystem // description, publication, traits, and rules
	AncestryFeatLevels  valueNode[[]int]      `json:"ancestryFeatLevels"`
	ClassFeatLevels     valueNode[[]int]      `json:"classFeatLevels"`
	GeneralFeatLevels   valueNode[[]int]      `json:"generalFeatLevels"`
	SkillFeatLevels     valueNode[[]int]      `json:"SkillFeatLevels"`
	SkillIncreaseLevels valueNode[[]int]      `json:"skillIncreaseLevels"`
	Attacks             attacks               `json:"attacks"`
	Defenses            defenses              `json:"defenses"`
	HP                  int                   `json:"hp"`
	Items               map[string]systemItem `json:"items"`
	KeyAbility          valueNode[[]string]   `json:"keyAbility"`
	Serception          int                   `json:"perception"`
	SavingThrows        savingThrows          `json:"savingThrows"`
	Spellcasting        int                   `json:"spellcasting"`
	ClassTrainedSkills  classTrainedSkills    `json:"trainedSkills"`
}

type classTrainedSkills struct {
	Additional int      `json:"additional"`
	Value      []string `json:"value"`
}

type savingThrows struct {
	Fortitude int `json:"fortitude"`
	Will      int `json:"will"`
	Reflex    int `json:"reflex"`
}

type additionalAttacks struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type attacks struct {
	additionalAttacks
	Unarmed  int `json:"unarmed"`
	Simple   int `json:"simple"`
	Martial  int `json:"martial"`
	Advanced int `json:"advanced"`
}

type defenses struct {
	Heavy     int `json:"heavy"`
	Medium    int `json:"medium"`
	Light     int `json:"light"`
	Unarmored int `json:"unarmored"`
}

// equipment specific
type price struct {
	Per int `json:"per,omitempty"`
	Value map[string]int `json:"value"` // will only have cp, sp, gp, or pp keys
}

type weaponSystem struct {
	commonSystem // description, publication, traits, and rules
	Level  valueNode[int]    `json:"level"`
	BaseItem string `json:"baseItem"`
	Bonus valueNode[int] `json:"bonus"`
	BonusDamage valueNode[int] `json:"bonusDamage"`
	Bulk valueNode[float64] `json:"bulk"`
	Category string `json:"category"`
	Group string `json:"group"`
	Quantity int `json:"quantity,omitempty"`
	HP itemHP `json:"hp"`
	Hardness int `json:"hardness"`
	Price price `json:"price"`
	Expend int `json:"expend,omitempty"`
	Range int `json:"range,omitempty"`
	Material material `json:"material"`
	Usage usage `json:"usage"`
	SplashDamage valueNode[int] `json:"splashDamage"`
	Size string `json:"size"`
	Damage damage `json:"damage"`
	Reload valueNode[string] `json:"reload"` // will be null, "-", or a string number like "1"
}

type itemHP struct {
	Current int `json:"current"`
	Max int `json:"max"`
}

type material struct {
	Grade string `json:"grade"`
	Type string `json:"type"`
}

// if die == "", then it is flat damage of "Dice" amount
type damage struct {
	Type string `json:"damageType"`
	Dice int `json:"dice"`
	Die string `json:"die"`
	PersistentDamage persistentDamage `json:"persistent"`
}

// why. oh god why. are the keys here different from the above struct???
// faces == die, Number == dice, type == damageType.
type persistentDamage struct {
	Faces int `json:"faces"` // if null (might before to 0 for int), then its flat damage of "Number" amount.
	Number int `json:"number"`
	Type string `json:"type"`
}

type usage struct {
	CanBeAmmo bool `json:"canBeAmmo,omitempty"`
	Value string `json:"value"`
}
