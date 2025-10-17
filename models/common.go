package models

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type foundryType interface {
	// types
	Ancestry |
		Background |
		Class |
		EquipmentEffect |
		EquipmentEnvelope |
		Feature |
		FeatEffect |
		OtherEffect |
		Heritage
}

type filterable interface {
	// funcs
	IsLegacy() bool
	HasProvidedLicense(string) bool
}

type FoundryModel interface {
	foundryType
	filterable
}

// common shared types
type commonSystem struct {
	Description description `json:"description"`
	Publication publication `json:"publication"`
	Traits      traits      `json:"traits"`
	Rules       []rule      `json:"rules"`
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
	Level maybeStringAsInt `json:"level"` // the foundryvtt/pf2e project has INSANE data choices. this could be string ("1") or int (1).
	Name  string           `json:"name"`
	UUID  string           `json:"uuid"`
}

// type to handle the cases where a value that's expected to be a string could also be an int
type maybeIntAsString struct {
	Value string
}

func (maybeIntAsString *maybeIntAsString) UnmarshalJSON(b []byte) error {
	var f float64
	err := json.Unmarshal(b, &f)
	if err == nil {
		maybeIntAsString.Value = strconv.Itoa(int(f))
		return nil
	}

	var s string
	err = json.Unmarshal(b, &s)
	if err == nil {
		maybeIntAsString.Value = s
		return nil
	}

	return fmt.Errorf("maybeIntAsString.value was not float64 or string: %s", b)
}

type maybeStringAsInt struct {
	Value int
}

func (maybeStringAsInt *maybeStringAsInt) UnmarshalJSON(b []byte) error {
	var f float64
	err := json.Unmarshal(b, &f)
	if err == nil {
		maybeStringAsInt.Value = int(f)
		return nil
	}

	var s string
	err = json.Unmarshal(b, &s)
	if err == nil {
		if s == "" {
			maybeStringAsInt.Value = 0
		} else {
			maybeStringAsInt.Value, err = strconv.Atoi(s)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return fmt.Errorf("maybeStringAsInt.value was not float64 or string: %s", b)
}

type maybeStringAsSlice struct {
	Value []string
}

func (maybeStringAsSlice *maybeStringAsSlice) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err == nil {
		maybeStringAsSlice.Value = []string{s}
		return nil
	}

	var ss []string
	err = json.Unmarshal(b, &ss)
	if err == nil {
		maybeStringAsSlice.Value = ss
		return nil
	}

	return fmt.Errorf("maybeStringAsSlice.value was not string or []string: %s", b)
}

type rule struct {
	Key   string `json:"key"`
	Value any
}

func (r *rule) UnmarshalJSON(b []byte) error {
	var temp struct {
		Key string `json:"key"`
	}

	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	switch temp.Key {
	// TODO implement to type and set as r.Value
	}

	return nil
}

// NOTE RULE TO ADD TO RULE UNMARSHAL
type actorTraitsRule struct {
	Add       []string        `json:"add,omitempty"`
	Remove    []string        `json:"remove,omitempty"`
	Predicate []predicateItem `json:"predicate,omitempty"`
}

type predicateItem struct {
	Value any
}

type predicateStringItem string

type predicateComplexItem struct {
	Or  []string `json:"or"`
	And []string `json:"and"`
	Gte []maybeIntAsString `json:"gte"`
}

func (p *predicateItem) UnmarshalJSON(b []byte) error {
	var s predicateStringItem
	err := json.Unmarshal(b, &s)
	if err == nil {
		p.Value = s
		return nil
	}

	var c predicateComplexItem
	err = json.Unmarshal(b, &c)
	if err == nil {
		p.Value = c
		return nil
	}

	return fmt.Errorf("Could not unmarshal precidate into string or complex: %s", b)
}

// NOTE RULE TO ADD TO RULE UNMARSHAL
type senseRule struct {
	Selector string `json:"selector"`
}

// NOTE RULE TO ADD TO RULE UNMARSHAL
type immunityRule struct {
	Type string `json:"type"`
}

// NOTE RULE TO ADD TO RULE UNMARSHAL
type flatModifierRule struct {
	Slug      string             `json:"slug"`
	Value     maybeIntAsString   `json:"value"`
	Selector  maybeStringAsSlice `json:"selector"`
	Predicate []predicateItem    `json:"predicate"`
	Type      string             `json:"type"`
	Label string `json:"label"`
}

type activeEffectLikeRule struct {
	Mode string `json:"mode"`
	Path string `json:"path"`
	Value int `json:"value"`
}

type rollOptionRule struct {
	Label string `json:"label"`
	Option string `json:"option"`	
	Toggleable bool `json:"toggleable"`
}

type grantItemRule struct {
	UUID string `json:"uuid"`
}

type strikeRule struct {
	BaseType string `json:"baseType"`
	Category string `json:"category"`
	Label string `json:"label"`
	Predicate []predicateItem    `json:"predicate"`
	Slug string `json:"slug"`
	Traits []string `json:"traits"`


}

type strikeRuleDamage struct {
	Base strikeRuleDamageBase `json:"base"`
}

type strikeRuleDamageBase struct {
	DamageType string `json:"damageType"`
	Dice int `json:"dice"`
	Die string `json:"die"`
}
