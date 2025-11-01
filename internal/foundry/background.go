package foundry

type Background struct {
	Name   string           `json:"name"`
	System backgroundSystem `json:"system"`
}

func (bg Background) IsLegacy() bool {
	return !bg.System.Publication.Remaster
}

func (bg Background) HasProvidedLicense(license string) bool {
	return bg.System.Publication.License == license
}

type backgroundTrainedSkills struct {
	Custom string   `json:"custom"`
	Lore   []string `json:"lore"`
	Value  []string `json:"value"`
}

type backgroundSystem struct {
	commonSystem                          // description, publication, traits, and rules
	Boosts        boosts                  `json:"boosts"`
	TrainedSkills backgroundTrainedSkills `json:"trainedSkills"`
	Items         map[string]SystemItem   `json:"items"`
}
