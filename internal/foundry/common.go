package foundry

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
		Heritage |
		SpellEffect |
		Spell
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

	// TODO make this []rule and implement the rules parsing.
	Rules json.RawMessage `json:"rules"`
}

type publication struct {
	Title    string `json:"title"`
	Remaster bool   `json:"remaster"`
	License  string `json:"license"`
}

type description struct {
	Value                 string `json:"value"`
	GameMasterDescription string `json:"gm"`
}

type traits struct {
	Rarity    string   `json:"rarity"`
	Value     []string `json:"value"`
	OtherTags []string `json:"otherTags"`
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

type SystemItem struct {
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
