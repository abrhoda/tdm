package models

type spell struct {
	Name   string      `json:"name"`
	System spellSystem `json:"system"`
}

type spellSystem struct {
	Description description `json:"description"`
	Publication publication `json:"publication"`
	Traits      spellTraits `json:"traits"`

	// TODO make this []rule and implement the rules parsing.
	Rules []any `json:"rules"`
}

type spellTraits struct {
	Rarity     string   `json:"rarity"`
	Traditions []string `json:"traditions"`
	Value      []string `json:"value"`
}

type area struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}
