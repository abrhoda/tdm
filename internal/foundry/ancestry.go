package foundry

type Ancestry struct {
	Name   string         `json:"name"`
	System ancestrySystem `json:"system"`
}

func (a Ancestry) IsLegacy() bool {
	return !a.System.Publication.Remaster
}

// TODO call this HasLicense instead
func (a Ancestry) HasProvidedLicense(license string) bool {
	return a.System.Publication.License == license
}

// ancestry specific
type additionalLanguages struct {
	Count  int      `json:"count"`
	Custom string   `json:"custom"`
	Value  []string `json:"value"`
}

type languages struct {
	Custom string   `json:"custom"`
	Value  []string `json:"value"`
}

type ancestrySystem struct {
	commonSystem                              // description, publication, traits, and rules
	AdditionalLanguages additionalLanguages   `json:"additionalLanguages"`
	Boosts              boosts                `json:"boosts"`
	Items               map[string]SystemItem `json:"items"`
	Flaws               boosts                `json:"flaws"`
	HP                  int                   `json:"hp"`
	Size                string                `json:"size"`
	Reach               int                   `json:"reach"`
	Hands               int                   `json:"hands"`
	Speed               int                   `json:"speed"`
	Languages           languages             `json:"languages"`
	Vision              string                `json:"vision"`
}
