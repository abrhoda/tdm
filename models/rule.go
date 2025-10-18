package models

import (
	"encoding/json"
	"fmt"
)

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

// predicate parsing
type predicateItem struct {
	Value any
}

type predicateComplexItem struct {
	Or  []predicateItem    `json:"or"`
	And []predicateItem    `json:"and"`
	Not []predicateItem    `json:"not"`
	Gte []maybeIntAsString `json:"gte"`
	Lte []maybeIntAsString `json:"lte"`
	Gt  []maybeIntAsString `json:"gt"`
	Lt  []maybeIntAsString `json:"lt"`
	Eq  []string           `json:"eq"`
}

func (p *predicateItem) UnmarshalJSON(b []byte) error {
	var s string
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

// rules
type rule struct {
	Key   string
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

	r.Key = temp.Key
	switch temp.Key {
	// TODO implement to type and set as r.Value
	case "sense":
	default:
		return fmt.Errorf("Unexpected key in rule: %s", temp.Key)
	}

	return nil
}

// NOTE RULE TO ADD TO RULE UNMARSHAL
type actorTraitsRule struct {
	Add       []string        `json:"add,omitempty"`
	Remove    []string        `json:"remove,omitempty"`
	Predicate []predicateItem `json:"predicate,omitempty"`
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
	Label     string             `json:"label"`
}

type activeEffectLikeRule struct {
	Mode  string `json:"mode"`
	Path  string `json:"path"`
	Value int    `json:"value"`
}

type rollOptionRule struct {
	Label      string `json:"label"`
	Option     string `json:"option"`
	Toggleable bool   `json:"toggleable"`
}

type grantItemRule struct {
	UUID string `json:"uuid"`
}

type strikeRule struct {
	BaseType  string           `json:"baseType"`
	Category  string           `json:"category"`
	Label     string           `json:"label"`
	Predicate []predicateItem  `json:"predicate"`
	Slug      string           `json:"slug"`
	Traits    []string         `json:"traits"`
	Damage    strikeRuleDamage `json:"damage"`
}

type strikeRuleDamage struct {
	Base strikeRuleDamageBase `json:"base"`
}

type strikeRuleDamageBase struct {
	DamageType string `json:"damageType"`
	Dice       int    `json:"dice"`
	Die        string `json:"die"`
}
