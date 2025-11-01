package storage

// INTERFACE DEFINTIONS
// NOTE could add additional `GetXXXByField1AndField2` type functions
type ProficiencyStore interface {
	CreateProficiency(proficiency Proficiency) error
	GetProficiencyByID(id int) (Proficiency, error)
	UpdateProficiency(proficiency Proficiency) error
	DeleteProficiency(id int) error
}

type SenseStore interface {
	CreateSense(sense Sense) error
	GetSenseByID(id int) (Sense, error)
	UpdateSense(sense Sense) error
	DeleteSense(id int) error
}

type AncestryStore interface {
	CreateAncestry(ancestry Ancestry) error
	GetAncestryByID(id int) (Ancestry, error)
	UpdateAncestry(ancestry Ancestry) error
	DeleteAncestry(id int) error
}

type AncestryPropertyStore interface {
	CreateAncestryProperty(ancestryProperty AncestryProperty) error
	GetAncestryPropertyByID(id int) (AncestryProperty, error)
	UpdateAncestryProperty(ancestryPropertys AncestryProperty) error
	DeleteAncestryProperty(id int) error
}

type InMemoryDatastore struct {
	Ancestries         []Ancestry
	AncestryProperties []AncestryProperty
}

type PostgresDatastore struct {
	// conn/dbpool using pgx
}

type JsonFileDatastore struct {
	//
}
