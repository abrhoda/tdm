package models

type Spell struct {
	Name   string      `json:"name"`
	System spellSystem `json:"system"`
}

func (s Spell) IsLegacy() bool {
	return !s.System.Publication.Remaster
}

func (s Spell) HasProvidedLicense(license string) bool {
	return s.System.Publication.License == license
}

type spellSystem struct {
	Description   description             `json:"description"`
	Publication   publication             `json:"publication"`
	Traits        spellTraits             `json:"traits"`
	Level         valueNode[int]          `json:"level"`
	Target        valueNode[string]       `json:"target"`
	Range         valueNode[string]       `json:"range"`
	Time          valueNode[string]       `json:"time"`
	Cost          valueNode[string]       `json:"cost"`
	Area          *area                   `json:"area"`
	Counteraction bool                    `json:"counteraction"`
	Requirements  string                  `json:"requirements"`
	Defense       *spellDefense           `json:"defense"`
	Damage        *map[string]spellDamage `json:"damage"`
	Duration      spellDuration           `json:"duration"`
	Heightening   *spellHeightening       `json:"heigthening"`
	Overlays      *map[string]overlay     `json:"overlays"`

	// TODO make this []rule and implement the rules parsing.
	Rules []any `json:"rules"`
}

type spellTraits struct {
	Rarity     string   `json:"rarity"`
	Traditions []string `json:"traditions"`
	Value      []string `json:"value"`
}

// AreaType only used when area struct is in the heightening struct.
type area struct {
	AreaType string           `json:"areaType"`
	Type     string           `json:"type"`
	Value    maybeStringAsInt `json:"value"`
}

type spellDamage struct {
	ApplyMod  bool     `json:"applyMod"`
	Category  string   `json:"category"` // null, persistent, splash. not making a pointer, just check for ""
	Formula   string   `json:"formula"`  // this could be "1" or "2d4" strings
	Kinds     []string `json:"kinds"`
	Materials []string `json:"materials"`
	Type      string   `json:"type"`
}

type spellDefense struct {
	Save spellDefenseSave `json:"save"`
}

type spellDefenseSave struct {
	Basic     bool   `json:"basic"`
	Statistic string `json:"statistic"`
}

type spellDuration struct {
	Sustained bool   `json:"sustained"`
	Value     string `json:"value"`
}

// type == "interval" means that for every interval, you use the damage of (interval + base level) [if a 2nd level spell is highethened +2, you would use level 4 damage] and multiply (area * interval) + base area value.
type spellHeightening struct {
	Area     int                              `json:"area"`
	Type     string                           `json:"type"`
	Interval int                              `json:"interval"` // this means the value on "heightened +VALUE" syntax if type field == "interval"
	Levels   *map[string]spellHeightenedLevel `json:"levels"`
	Damage   *map[string]string               `json:"damage"` // seems to be a map of "int":"diceFormula" or "uuid that match a damage key from the top level damage map": "diceFormula"
}

// ANY of these are liable to be nil (not present in the json payload)
type spellHeightenedLevel struct {
	Range  *valueNode[string]      `json:"range"`
	Target *valueNode[string]      `json:"target"`
	Area   *area                   `json:"area"`
	Damage *map[string]spellDamage `json:"damage"` // multiple keys in this map mean both apply!
}

// this is where the 1-3 action information is held. Each overlay has a 1, 2, or 3 action info
type overlay struct {
	Name        string      `json:"name"`
	OverlayType string      `json:"overlayType"`
	Sort        int         `json:"sort"`
	System      spellSystem `json:"system"`
}
