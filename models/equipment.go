package main

import (
	"encoding/json"
	"fmt"
)

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
