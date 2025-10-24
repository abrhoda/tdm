package foundry

import (
	"encoding/json"
	"fmt"
	"strconv"
)

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

type OtherEffect struct {
	Effect
}

func (e OtherEffect) IsLegacy() bool {
	return !e.System.Publication.Remaster
}

func (e OtherEffect) HasProvidedLicense(license string) bool {
	return e.System.Publication.License == license
}

type SpellEffect struct {
	Effect
}

func (e SpellEffect) IsLegacy() bool {
	return !e.System.Publication.Remaster
}

func (e SpellEffect) HasProvidedLicense(license string) bool {
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

// NOTE custom UnmarshalJSON due to below issue with `value` in badge
// use this to apply rules
// if type == counter, value is an int
// if type == formula, value is a dice expression which might have a syntax of "XdY" where X could be an int or a tag like `(@item.level)`
type badge struct {
	Max        int      `json:"max,omitempty"`
	Min        int      `json:"min,omitempty"`
	Evaluate   bool     `json:"evaluate,omitempty"`
	Labels     []string `json:"labels,omitempty"`
	Reevaluate string   `json:"reevaluate,omitempty"`
	Loop       bool     `json:"loop,omitempty"`
	Type       string   `json:"type"`
	Value      badgeValue
}

type badgeValue struct {
	value string
}

func (bv *badgeValue) UnmarshalJSON(b []byte) error {
	temp := map[string]any{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return nil
	}

	switch val := temp["value"].(type) {
	case string:
		bv.value = val
	case float64:
		bv.value = strconv.Itoa(int(val))
	case nil:
		bv.value = ""
	default:
		return fmt.Errorf("effect.system.badge.value has an unrecognized type of %T\n", val)
	}

	return nil
}

// I dont think this is actually ever used?
type start struct {
	Initiative any `json:"initiative"`
	Value      int `json:"value"`
}
