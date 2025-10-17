package models

type heritage struct {
	Name string
	System heritageSystem
}

type heritageSystem struct {
	commonSystem
}

func (h heritage) IsLegacy() bool {
	return !h.System.Publication.Remaster
}

func (h heritage) HasProvidedLicense(license string) bool {
	return h.System.Publication.License == license
}


