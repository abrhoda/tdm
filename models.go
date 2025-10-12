package main

// main foundry types
type ancestry struct {
	Name string `json:"name"`
	System ancestrySystem `json:"system"`
}

func (a ancestry) IsLegacy() bool {
	return !a.System.Publication.Remaster
}

func (a ancestry) HasProvidedLicense(license string) bool {
	return a.System.Publication.License == license
}

type ancestryFeature struct {
	Name string `json:"name"`
	System ancestryFeatureSystem `json:"system"`
}

func (af ancestryFeature) IsLegacy() bool {
	return !af.System.Publication.Remaster
}

func (af ancestryFeature) HasProvidedLicense(license string) bool {
	return af.System.Publication.License == license
}

type background struct {
	Name string `json:"name"`
	System backgroundSystem `json:"system"`
}

func (bg background) IsLegacy() bool {
	return !bg.System.Publication.Remaster
}

func (bg background) HasProvidedLicense(license string) bool {
	return bg.System.Publication.License == license
}

type class struct {
	Name string
}

type classFeature struct {
	Name string
}

type equipmentEffect struct {
	Name string
}

type equipment struct {
	Name string
}

type feature struct {
	Name string
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
	ancestryFeature | 
	background | 
	class | 
	classFeature |
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
	Traits traits `json:"traits"`
	Rules []any `json:"rules"`
}

type publication struct {
	Title string `json:"title"`
	Remaster bool  `json:"remaster"`
	License string `json:"license"`
}

type description struct {
	Value string `json:"value,omitempty"`
	GameMasterDescription string `json:"gm,omitempty"`
}

type traits struct {
	Rarity string `json:"rarity"`
	Value []string `json:"value"`
	OtherTags []string `json:"otherTags,omitempty"`
}

type boosts struct {
	First individualBoost `json:"0"`
	Second individualBoost `json:"1"`
	Third individualBoost `json:"2"`
}

type individualBoost struct {
	Value []string `json:"value"`
}

type systemItem struct {
	Level int `json:"level"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

// ancestry specific
type additionalLanguages struct {
	Count int `json:"count"`
	Custom string `json:"custom"`
	Value []string `json:"value"`
}

type languages struct {
	Custom string `json:"custom"`
	Value []string `json:"value"`
}

type ancestrySystem struct {
	commonSystem
	AdditionalLanguages additionalLanguages `json:"additionalLanguages"`
	Boosts boosts `json:"boosts"`
	Items map[string]systemItem `json:"items"`
	Flaws boosts `json:"flaws"`
	HP int `json:"hp"`
	Size string `json:"size"`
	Reach int `json:"reach"`
	Hands int `json:"hands"`
	Speed int `json:"spped"`
	Languages languages `json:"languages"`
	Vision string `json:"vision"`
}

// NOTE these might be shared across all feats?
// ancestryFeature specific
type ancestryFeatureSystem struct {
	commonSystem
	ActionType map[string]string `json:"actionType"`
	Actions map[string]string `json:"actions"`
	Category string `json:"category"`
	Level map[string]int `json:"level"`
	Prerequisites map[string][]string `json:"prerequisites"`
}

// background specific
type trainedSkills struct {
	Custom string `json:"custom,omitempty"`
	Lore []string `json:"lore"`
	Value []string `json:"value"`
}

type backgroundSystem struct {
	commonSystem
	Boosts boosts `json:"boosts"`
	TrainedSkills trainedSkills `json:"trainedSkills"`
	Items map[string]systemItem `json:"items"`
}
