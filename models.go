package main

// main foundry types
type ancestry struct {
	name string
}

type ancestryFeature struct {
	name string
}

type background struct {
	Name string `json:"name"`
	System backgroundSystem `json:"system"`
}

type class struct {
	name string
}

type classFeature struct {
	name string
}

type equipmentEffect struct {
	name string
}

type equipment struct {
	name string
}

type feature struct {
	name string
}

type featureEffect struct {
	name string
}

type heritage struct {
	name string
}

type foundryType interface {
	ancestry | ancestryFeature | background | class | classFeature | equipmentEffect | equipment | feature | featureEffect | heritage
}

// common shared types
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
}

// background specific
type boosts struct {
	First individualBoost `json:"0"`
	Second individualBoost `json:"1"`
}

type individualBoost struct {
	Value []string `json:"value"`
}

type trainedSkills struct {
	Custom string `json:"custom,omitempty"`
	Lore []string `json:"lore"`
	Value []string `json:"value"`
}

// NOTE this might be more shared?
type backgroundItem struct {
	Level int `json:"level"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}

type backgroundSystem struct {
	Boosts boosts `json:"boosts"`
	Description description `json:"description"`
	Publication publication `json:"publication"`
	TrainedSkills trainedSkills `json:"trainedSkills"`
	Traits traits `json:"traits"`
	Rules []any `json:"rules"`
	Items map[string]backgroundItem `json:"items"`
}
