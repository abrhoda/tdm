package models

import (
	"encoding/json"
	"fmt"
)

type equipmentEffect struct {
	Name string
}

type Equipment struct {
	payload any
}

type armor struct {
	Name   string
	System armorSystem
}

type backpack struct {
	Name   string
	System backpackSystem
}

type weapon struct {
	Name   string
	System weaponSystem
}

// implementing filterable on the equipment itself.
// TODO think about moving to the concrete types instead of equipment
func (e Equipment) IsLegacy() bool {
	switch target := e.payload.(type) {
	case armor:
		 return !target.System.Publication.Remaster
	case backpack:
		 return !target.System.Publication.Remaster
	case weapon:
		 return !target.System.Publication.Remaster
	default:
		return false
	}
}

func (e Equipment) HasProvidedLicense(license string) bool {
	switch target := e.payload.(type) {
	case armor:
		 return target.System.Publication.License == license
	case backpack:
		 return target.System.Publication.License == license
	case weapon:
		 return target.System.Publication.License == license
	default:
		return false
	}
}

func (e *Equipment) UnmarshalJSON(b []byte) error {
	var temp struct {
		Type string          `json:"type"`
		Name string          `json:"name"`
		Rest json.RawMessage `json:"system"`
	}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	switch temp.Type {
	case "armor":
		var armor armor
		err = json.Unmarshal(temp.Rest, &armor.System)
		if err != nil {
			return err
		}
		armor.Name = temp.Name
		e.payload = armor
	case "backpack":
		var backpack backpack
		err = json.Unmarshal(temp.Rest, &backpack.System)
		if err != nil {
			return err
		}
		backpack.Name = temp.Name
		e.payload = backpack
	case "weapon":
		var weapon weapon
		err = json.Unmarshal(temp.Rest, &weapon.System)
		if err != nil {
			return err
		}
		weapon.Name = temp.Name
		e.payload = weapon
	default:
		return fmt.Errorf("Unknown equipment type: %s", temp.Type)
	}

	return nil
}

type price struct {
	Per   int            `json:"per,omitempty"`
	Value map[string]int `json:"value"` // will only have cp, sp, gp, or pp keys
}

type physicalSystem struct {
	BaseItem string             `json:"baseItem"`
	Bulk     valueNode[float64] `json:"bulk"`
	HP       itemHP             `json:"hp"`
	Price    price              `json:"price"`
	Hardness int                `json:"hardness"`
	Level    valueNode[int]     `json:"level"`
	Quantity int                `json:"quantity,omitempty"`
	Size     string             `json:"size"`
}

type weaponSystem struct {
	commonSystem                     // description, publication, traits, and rules
	physicalSystem                   // from template.json physical category
	Bonus          valueNode[int]    `json:"bonus"`
	BonusDamage    valueNode[int]    `json:"bonusDamage"`
	Category       string            `json:"category"`
	Group          string            `json:"group"`
	Expend         int               `json:"expend,omitempty"`
	Material       material          `json:"material"`
	Usage          usage             `json:"usage"`
	SplashDamage   valueNode[int]    `json:"splashDamage"`
	Damage         damage            `json:"damage"`
	Reload         valueNode[string] `json:"reload"` // will be null, "-", or a string number like "1"
	Range          int               `json:"range,omitempty"`
	WeaponRunes    weaponRunes       `json:"runes"`
}

type weaponRunes struct {
	Potency  int      `json:"potency"`
	Striking int      `json:"striking"`
	Property []string `json:"property"`
}

type itemHP struct {
	Current int `json:"current"`
	Max     int `json:"max"`
}

type material struct {
	Grade string `json:"grade"`
	Type  string `json:"type"`
}

// if die == "", then it is flat damage of "Dice" amount
type damage struct {
	Type             string           `json:"damageType"`
	Dice             int              `json:"dice"`
	Die              string           `json:"die"`
	PersistentDamage persistentDamage `json:"persistent"`
}

// why. oh god why. are the keys here different from the above struct???
// faces == die, Number == dice, type == damageType.
type persistentDamage struct {
	Faces  int    `json:"faces"` // if null (might before to 0 for int), then its flat damage of "Number" amount.
	Number int    `json:"number"`
	Type   string `json:"type"`
}

type usage struct {
	CanBeAmmo bool   `json:"canBeAmmo,omitempty"`
	Value     string `json:"value"`
}

type armorSystem struct {
	commonSystem
	physicalSystem
	ACBonus      int        `json:"acBonus"`
	Category     string     `json:"category"`
	Group        string     `json:"group"`
	CheckPenalty int        `json:"checkPenalty"`
	DexCap       int        `json:"dexCap"`
	ArmorRunes   armorRunes `json:"runes"`
	SpeedPenalty int        `json:"speedPenalty"`
	Strength     int        `json:"strength"`
}

type armorRunes struct {
	Potency   int      `json:"potency"`
	Resilient int      `json:"resilient"`
	Property  []string `json:"property"`
}

type backpackSystem struct {
	commonSystem
	// cant use physicalSystem because `bulk` has additional keys
	BaseItem  string            `json:"baseItem"`
	Bulk      backpackBulk      `json:"bulk"`
	HP        itemHP            `json:"hp"`
	Price     price             `json:"price"`
	Hardness  int               `json:"hardness"`
	Level     valueNode[int]    `json:"level"`
	Quantity  int               `json:"quantity,omitempty"`
	Size      string            `json:"size"`
	Usage     valueNode[string] `json:"usage"`
	Collapsed bool              `json:"collapsed"`
	Stowing   bool              `json:"stowing"`
}

type backpackBulk struct {
	Value        float64 `json:"value"`
	HeldOrStowed int     `json:"heldOrStowed"`
	Ignored      int     `json:"ignored"`
	Capacity     int     `json:"capacity"`
}
