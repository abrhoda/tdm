package main

type class struct {
	Name   string      `json:"name"`
	System classSystem `json:"system"`
}

func (c class) IsLegacy() bool {
	return !c.System.Publication.Remaster
}

func (c class) HasProvidedLicense(license string) bool {
	return c.System.Publication.License == license
}

type classSystem struct {
	commonSystem // description, publication, traits, and rules
	AncestryFeatLevels  valueNode[[]int]      `json:"ancestryFeatLevels"`
	ClassFeatLevels     valueNode[[]int]      `json:"classFeatLevels"`
	GeneralFeatLevels   valueNode[[]int]      `json:"generalFeatLevels"`
	SkillFeatLevels     valueNode[[]int]      `json:"SkillFeatLevels"`
	SkillIncreaseLevels valueNode[[]int]      `json:"skillIncreaseLevels"`
	Attacks             attacks               `json:"attacks"`
	Defenses            defenses              `json:"defenses"`
	HP                  int                   `json:"hp"`
	Items               map[string]systemItem `json:"items"`
	KeyAbility          valueNode[[]string]   `json:"keyAbility"`
	Serception          int                   `json:"perception"`
	SavingThrows        savingThrows          `json:"savingThrows"`
	Spellcasting        int                   `json:"spellcasting"`
	ClassTrainedSkills  classTrainedSkills    `json:"trainedSkills"`
}

type classTrainedSkills struct {
	Additional int      `json:"additional"`
	Value      []string `json:"value"`
}

type savingThrows struct {
	Fortitude int `json:"fortitude"`
	Will      int `json:"will"`
	Reflex    int `json:"reflex"`
}

type additionalAttacks struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type attacks struct {
	additionalAttacks
	Unarmed  int `json:"unarmed"`
	Simple   int `json:"simple"`
	Martial  int `json:"martial"`
	Advanced int `json:"advanced"`
}

type defenses struct {
	Heavy     int `json:"heavy"`
	Medium    int `json:"medium"`
	Light     int `json:"light"`
	Unarmored int `json:"unarmored"`
}
