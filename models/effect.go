package models

type EquipmentEffect struct {
	Name   string       `json:"name"`
	Type   string       `json:"type"`
	System effectSystem `json:"system"`
}

func (ee EquipmentEffect) IsLegacy() bool {
	return !ee.System.Publication.Remaster
}

func (ee EquipmentEffect) HasProvidedLicense(license string) bool {
	return ee.System.Publication.License == license
}

type effectSystem struct {
	commonSystem
	Level    valueNode[int] `json:"level"`
	Duration duration       `json:"duration"`
}

type duration struct {
	Expiry    string `json:"expiry"`
	Sustained bool   `json:"sustained"`
	Unit      string `json:"unit"`
	Value     int    `json:"value"`
}
