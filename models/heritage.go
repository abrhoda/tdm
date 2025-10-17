package models

type Heritage struct {
	Name string
	System heritageSystem
}

type heritageSystem struct {
	commonSystem
	Ancestry ancestryInHeritage `json:"ancestry"`
}

func (h Heritage) IsLegacy() bool {
	return !h.System.Publication.Remaster
}

func (h Heritage) HasProvidedLicense(license string) bool {
	return h.System.Publication.License == license
}

type ancestryInHeritage struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
	UUID string `json:"uuid"`
}
