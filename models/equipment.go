package models

import (
	"encoding/json"
	"fmt"
)

type EquipmentEnvelope struct {
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

type consumable struct {
	Name   string
	System consumableSystem
}

type equipment struct {
	Name   string
	System equipmentSystem
}

type kit struct {
	Name   string
	System kitSystem
}

type shield struct {
	Name   string
	System shieldSystem
}

type treasure struct {
	Name   string
	System treasureSystem
}

type weapon struct {
	Name   string
	System weaponSystem
}

// implementing filterable on the equipment itself.
// TODO think about moving to the concrete types instead of equipment
func (e EquipmentEnvelope) IsLegacy() bool {
	switch target := e.payload.(type) {
	case armor:
		return !target.System.Publication.Remaster
	case backpack:
		return !target.System.Publication.Remaster
	case consumable:
		return !target.System.Publication.Remaster
	case equipment:
		return !target.System.Publication.Remaster
	case kit:
		return !target.System.Publication.Remaster
	case shield:
		return !target.System.Publication.Remaster
	case treasure:
		return !target.System.Publication.Remaster
	case weapon:
		return !target.System.Publication.Remaster
	default:
		return false
	}
}

func (e EquipmentEnvelope) HasProvidedLicense(license string) bool {
	switch target := e.payload.(type) {
	case armor:
		return target.System.Publication.License == license
	case backpack:
		return target.System.Publication.License == license
	case consumable:
		return target.System.Publication.License == license
	case equipment:
		return target.System.Publication.License == license
	case kit:
		return target.System.Publication.License == license
	case shield:
		return target.System.Publication.License == license
	case treasure:
		return target.System.Publication.License == license
	case weapon:
		return target.System.Publication.License == license
	default:
		return false
	}
}

func (e *EquipmentEnvelope) UnmarshalJSON(b []byte) error {
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
	case "consumable":
		var consumable consumable
		err = json.Unmarshal(temp.Rest, &consumable.System)
		if err != nil {
			return err
		}
		consumable.Name = temp.Name
		e.payload = consumable
	case "equipment":
		var equipment equipment
		err = json.Unmarshal(temp.Rest, &equipment.System)
		if err != nil {
			return err
		}
		equipment.Name = temp.Name
		e.payload = equipment
	case "kit":
		var kit kit
		err = json.Unmarshal(temp.Rest, &kit.System)
		if err != nil {
			return err
		}
		kit.Name = temp.Name
		e.payload = kit
	case "shield":
		var shield shield
		err = json.Unmarshal(temp.Rest, &shield.System)
		if err != nil {
			return err
		}
		shield.Name = temp.Name
		e.payload = shield
	case "treasure":
		var treasure treasure
		err = json.Unmarshal(temp.Rest, &treasure.System)
		if err != nil {
			return err
		}
		treasure.Name = temp.Name
		e.payload = treasure
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
	commonSystem                               // description, publication, traits, and rules
	physicalSystem                             // from template.json physical category
	Bonus          valueNode[int]              `json:"bonus"`
	BonusDamage    valueNode[int]              `json:"bonusDamage"`
	Category       string                      `json:"category"`
	Group          string                      `json:"group"`
	Expend         int                         `json:"expend,omitempty"`
	Material       material                    `json:"material"`
	Usage          usage                       `json:"usage"`
	SplashDamage   valueNode[maybeStringAsInt] `json:"splashDamage"` // NOTE this is once again because foundryvtt/pf2e has ABSOLUTELY NO STANDARDIZATION on their data. Juggling Club has a random empty string where which doesnt parse correctly into int. Ankylostar randomly has a null value.
	Damage         damage                      `json:"damage"`
	Reload         valueNode[string]           `json:"reload"` // will be null, "-", or a string number like "1"
	Range          int                         `json:"range,omitempty"`
	WeaponRunes    weaponRunes                 `json:"runes"`
}

//type splashDamage struct {
//	Value int
//}
//
//func (sd *splashDamage) UnmarshalJSON(b []byte) error {
//	temp := map[string]any{}
//	err := json.Unmarshal(b, &temp)
//	if err != nil {
//		return err
//	}
//
//	switch val := temp["value"].(type) {
//	case float64:
//		sd.Value = int(val)
//	case string:
//		if val == "" {
//			sd.Value = 0
//		} else {
//			sd.Value, err = strconv.Atoi(val)
//			if err != nil {
//				return err
//			}
//		}
//	case nil:
//		sd.Value = 0
//	default:
//		return fmt.Errorf("weapon.system.splashDamage.value has an unrecognized type of %T\n", val)
//	}
//
//	return nil
//}

type weaponRunes struct {
	Potency  int
	Striking int
	Property []string
}

// NOTE doing this because handwraps-of-mighty-blows RANDOMLY has an obj instead of a list of strings.
func (wr *weaponRunes) UnmarshalJSON(b []byte) error {
	temp := map[string]any{}
	err := json.Unmarshal(b, &temp)
	if err != nil {
		return err
	}

	// need to cast because encoding/json treats all numbers as float64 by default.
	wr.Potency = int(temp["potency"].(float64))
	wr.Striking = int(temp["striking"].(float64))

	switch val := temp["property"].(type) {
	case map[string]any:
		var props []string
		for _, v := range val {
			props = append(props, v.(string))
		}
		wr.Property = props
	case []any:
		var props []string
		for _, v := range val {
			props = append(props, v.(string))
		}
		wr.Property = props
	default:
		return fmt.Errorf("weapon.system.runes.property has an unrecognized type of %T\n", val)
	}
	return nil
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
	HeldOrStowed float64 `json:"heldOrStowed"`
	Ignored      int     `json:"ignored"`
	Capacity     int     `json:"capacity"`
}

type consumableSystem struct {
	commonSystem
	physicalSystem
	Category   string `json:"category"`
	Damage     damage `json:"damage"`
	Usage      usage  `json:"usage"`
	Uses       uses   `json:"uses"`
	StackGroup string `json:"stackGroup"`
	Spell      any    `json:"spell"` // TODO should grab the _id, name, and system.location.heightenedLevel here?
}

type uses struct {
	Value       int  `json:"value"`
	Max         int  `json:"max"`
	AutoDestroy bool `json:"autoDestroy"`
}

type equipmentSystem struct {
	commonSystem
	physicalSystem
	Usage valueNode[string] `json:"usage"`
}

type shieldSystem struct {
	commonSystem
	physicalSystem
	ACBonus      int         `json:"acBonus"`
	SpeedPenalty int         `json:"speedPenalty"`
	SubItems     any         `json:"subItems"`
	Runes        shieldRunes `json:"runes"`
	Specific     specific    `json:"specific"`
}

type specific struct {
	Material   material    `json:"material"`
	Runes      shieldRunes `json:"runes"`
	Integrated integrated  `json:"integrated"`
}

type integrated struct {
	Runes integratedShieldRunes `json:"runes"`
}

type integratedShieldRunes struct {
	Potency  int      `json:"potency"`
	Striking int      `json:"striking"`
	Property []string `json:"property"`
}

type shieldRunes struct {
	Reinforcing int `json:"reinforcing"`
}

type treasureSystem struct {
	commonSystem
	physicalSystem
	StackGroup string `json:"stackGroup"`
}

type kitSystem struct {
	commonSystem
	Price price     `json:"price"`
	Items []kitItem `json:"item"`
}

type kitItem struct {
	IsContainer bool      `json:"isContainer"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	UUID        string    `json:"uuid"` // might need this to link to the contained item.
	Items       []kitItem `json:"items"`
}
