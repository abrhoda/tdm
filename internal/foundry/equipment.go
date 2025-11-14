package foundry

import (
	"encoding/json"
	"fmt"
)

type EquipmentEnvelope struct {
	payload any
}

type Ammo struct {
	Name   string
	System ammoSystem
}

type Armor struct {
	Name   string
	System armorSystem
}

type Backpack struct {
	Name   string
	System backpackSystem
}

type Consumable struct {
	Name   string
	System consumableSystem
}

type Equipment struct {
	Name   string
	System equipmentSystem
}

type Kit struct {
	Name   string
	System kitSystem
}

type Shield struct {
	Name   string
	System shieldSystem
}

type Treasure struct {
	Name   string
	System treasureSystem
}

type Weapon struct {
	Name   string
	System weaponSystem
}

// implementing filterable on the equipment itself.
// TODO think about moving to the concrete types instead of equipment
func (e EquipmentEnvelope) IsLegacy() bool {
	switch target := e.payload.(type) {
	case Ammo:
		return !target.System.Publication.Remaster
	case Armor:
		return !target.System.Publication.Remaster
	case Backpack:
		return !target.System.Publication.Remaster
	case Consumable:
		return !target.System.Publication.Remaster
	case Equipment:
		return !target.System.Publication.Remaster
	case Kit:
		return !target.System.Publication.Remaster
	case Shield:
		return !target.System.Publication.Remaster
	case Treasure:
		return !target.System.Publication.Remaster
	case Weapon:
		return !target.System.Publication.Remaster
	default:
		return false
	}
}

func (e EquipmentEnvelope) HasProvidedLicense(license string) bool {
	switch target := e.payload.(type) {
	case Ammo:
		return target.System.Publication.License == license
	case Armor:
		return target.System.Publication.License == license
	case Backpack:
		return target.System.Publication.License == license
	case Consumable:
		return target.System.Publication.License == license
	case Equipment:
		return target.System.Publication.License == license
	case Kit:
		return target.System.Publication.License == license
	case Shield:
		return target.System.Publication.License == license
	case Treasure:
		return target.System.Publication.License == license
	case Weapon:
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
	case "ammo":
		var ammo Ammo
		err = json.Unmarshal(temp.Rest, &ammo.System)
		if err != nil {
			return err
		}
		ammo.Name = temp.Name
		e.payload = ammo
	case "armor":
		var armor Armor
		err = json.Unmarshal(temp.Rest, &armor.System)
		if err != nil {
			return err
		}
		armor.Name = temp.Name
		e.payload = armor
	case "backpack":
		var backpack Backpack
		err = json.Unmarshal(temp.Rest, &backpack.System)
		if err != nil {
			return err
		}
		backpack.Name = temp.Name
		e.payload = backpack
	case "consumable":
		var consumable Consumable
		err = json.Unmarshal(temp.Rest, &consumable.System)
		if err != nil {
			return err
		}
		consumable.Name = temp.Name
		e.payload = consumable
	case "equipment":
		var equipment Equipment
		err = json.Unmarshal(temp.Rest, &equipment.System)
		if err != nil {
			return err
		}
		equipment.Name = temp.Name
		e.payload = equipment
	case "kit":
		var kit Kit
		err = json.Unmarshal(temp.Rest, &kit.System)
		if err != nil {
			return err
		}
		kit.Name = temp.Name
		e.payload = kit
	case "shield":
		var shield Shield
		err = json.Unmarshal(temp.Rest, &shield.System)
		if err != nil {
			return err
		}
		shield.Name = temp.Name
		e.payload = shield
	case "treasure":
		var treasure Treasure
		err = json.Unmarshal(temp.Rest, &treasure.System)
		if err != nil {
			return err
		}
		treasure.Name = temp.Name
		e.payload = treasure
	case "weapon":
		var weapon Weapon
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
	Per   int            `json:"per"`
	Value map[string]int `json:"value"` // will only have cp, sp, gp, or pp keys
}

type physicalSystem struct {
	BaseItem string             `json:"baseItem"`
	Bulk     valueNode[float64] `json:"bulk"`
	HP       itemHP             `json:"hp"`
	Price    price              `json:"price"`
	Hardness int                `json:"hardness"`
	Level    valueNode[int]     `json:"level"`
	Quantity int                `json:"quantity"`
	Size     string             `json:"size"`
}

type ammoSystem struct {
	commonSystem
	BaseItem    string             `json:"baseItem"`
	Bulk        valueNode[float64] `json:"bulk"`
	CraftableAs []string           `json:"craftableAs"`
	Price       price              `json:"price"`
	Level       valueNode[int]     `json:"level"`
	Quantity    int                `json:"quantity"`
	Size        string             `json:"size"`
	Uses        uses               `json:"uses"`
}

type weaponSystem struct {
	commonSystem                               // description, publication, traits, and rules
	physicalSystem                             // from template.json physical category
	Bonus          valueNode[int]              `json:"bonus"`
	BonusDamage    valueNode[int]              `json:"bonusDamage"`
	Category       string                      `json:"category"`
	Group          string                      `json:"group"`
	Expend         *int                        `json:"expend"`
	Material       material                    `json:"material"`
	Usage          usage                       `json:"usage"`
	SplashDamage   valueNode[maybeStringAsInt] `json:"splashDamage"` // NOTE this is once again because foundryvtt/pf2e has ABSOLUTELY NO STANDARDIZATION on their data. Juggling Club has a random empty string where which doesnt parse correctly into int. Ankylostar randomly has a null value.
	Damage         damage                      `json:"damage"`
	MeleeUsage     meleeUsage                  `json:"meleeUsage"`
	Reload         valueNode[*string]          `json:"reload"` // will be null, "-" (meaning its thrown and must be drawn with an interact action first), or a string number like "1"
	Range          *int                        `json:"range"`
	WeaponRunes    weaponRunes                 `json:"runes"`
	Ammo           *weaponAmmo                 `json:"ammo"`
}

type meleeUsage struct {
	Damage damage   `json:"damage"`
	Group  string   `json:"group"`
	Traits []string `json:"traits"`
}

type weaponAmmo struct {
	BaseType string `json:"baseType"`
	BuiltIn  bool   `json:"builtIn"`
	Capacity int    `json:"capacity"`
}

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
	CanBeAmmo bool   `json:"canBeAmmo"`
	Value     string `json:"value"`
}

type armorSystem struct {
	commonSystem
	physicalSystem
	Material     material   `json:"material"`
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
	physicalSystem
	Description  description     `json:"description"`
	Publication  publication     `json:"publication"`
	Rules        json.RawMessage `json:"rules"`
	Material     material        `json:"material"`
	ACBonus      int             `json:"acBonus"`
	SpeedPenalty int             `json:"speedPenalty"`
	SubItems     any             `json:"subItems"`
	Runes        shieldRunes     `json:"runes"`
	Specific     specific        `json:"specific"`
	Traits       shieldTraits    `json:"traits"`
}

type shieldTraits struct {
	Rarity     string     `json:"rarity"`
	Value      []string   `json:"value"`
	OtherTags  []string   `json:"otherTags"`
	Integrated integrated `json:"integrated"`
}

// TODO specific appears in armor + weapons but this only accounts for shields shields.
type specific struct {
	Material material    `json:"material"`
	Runes    shieldRunes `json:"runes"`
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
