package main

// main foundry types
type ancestry struct {
	name string
}

type ancestryFeature struct {
	name string
}

type background struct {
	name string
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

// other types
type publication struct {
	Title string `json:"title"`
	Remaster bool  `json:"remaster"`
	License string `json:"license"`
}

type Description struct {
	Value string `json:"value,omitempty"`
	GameMasterDescription string `json:"gm,omitempty"`
}
