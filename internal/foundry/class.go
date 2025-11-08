package foundry

type Class struct {
	Name   string      `json:"name"`
	System classSystem `json:"system"`
}

func (c Class) IsLegacy() bool {
	return !c.System.Publication.Remaster
}

func (c Class) HasProvidedLicense(license string) bool {
	return c.System.Publication.License == license
}

type classSystem struct {
	commonSystem                              // description, publication, traits, and rules
	AncestryFeatLevels  valueNode[[]int]      `json:"ancestryFeatLevels"`
	ClassFeatLevels     valueNode[[]int]      `json:"classFeatLevels"`
	GeneralFeatLevels   valueNode[[]int]      `json:"generalFeatLevels"`
	SkillFeatLevels     valueNode[[]int]      `json:"SkillFeatLevels"`
	SkillIncreaseLevels valueNode[[]int]      `json:"skillIncreaseLevels"`
	Attacks             attacks               `json:"attacks"`
	Defenses            defenses              `json:"defenses"`
	HP                  int                   `json:"hp"`
	Items               map[string]SystemItem `json:"items"`
	KeyAbility          valueNode[[]string]   `json:"keyAbility"`
	Perception          int                   `json:"perception"`
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

type otherAttacks struct {
	Name string `json:"name"`
	Rank int    `json:"rank"`
}

type attacks struct {
	Other    otherAttacks `json:"other"`
	Unarmed  int          `json:"unarmed"`
	Simple   int          `json:"simple"`
	Martial  int          `json:"martial"`
	Advanced int          `json:"advanced"`
}

type defenses struct {
	Heavy     int `json:"heavy"`
	Medium    int `json:"medium"`
	Light     int `json:"light"`
	Unarmored int `json:"unarmored"`
}
