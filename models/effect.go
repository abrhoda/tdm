package models

type Effect struct {
	Name   string       `json:"name"`
	System effectSystem `json:"system"`
}

type EquipmentEffect struct {
	Effect
}

func (e EquipmentEffect) IsLegacy() bool {
	return !e.System.Publication.Remaster
}

func (e EquipmentEffect) HasProvidedLicense(license string) bool {
	return e.System.Publication.License == license
}

type FeatEffect struct {
	Effect
}

func (e FeatEffect) IsLegacy() bool {
	return !e.System.Publication.Remaster
}

func (e FeatEffect) HasProvidedLicense(license string) bool {
	return e.System.Publication.License == license
}

type effectSystem struct {
	commonSystem
	Level    valueNode[int] `json:"level"`
	Duration duration       `json:"duration"`
	Start    start          `json:"start"`
	Badge    badge          `json:"badge"`
}

type duration struct {
	Expiry    string `json:"expiry"`
	Sustained bool   `json:"sustained"`
	Unit      string `json:"unit"`
	Value     int    `json:"value"`
}

// use this to apply rules
// if type == counter, value is an int
// if type == formula, value is a dice expression which might have a syntax of "XdY" where X could be an int or a tag like `(@item.level)`
type badge struct {
	Max        int      `json:"max,omitempty"`
	Min        int      `json:"min,omitempty"`
	Evaluate   bool     `json:"evaluate,omitempty"`
	Labels     []string `json:"labels,omitempty"`
	Reevaluate bool     `json:"reevaluate,omitempty"`
	Loop       bool     `json:"loop,omitempty"`
	Type       string   `json:"type"`
	Value      string   `json:"value"`
}

// I dont think this is actually ever used?
type start struct {
	Initiative any `json:"initiative"`
	Value      int `json:"value"`
}
