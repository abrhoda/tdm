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

type AncestryFeatureStore interface {
	CreateAncestryFeature(ancestryFeature AncestryFeature) error
	GetAncestryFeatureByID(id int) (AncestryFeature, error)
	UpdateAncestryFeature(ancestryFeatures AncestryFeature) error
	DeleteAncestryFeature(id int) error
}

type InMemoryDatastore struct {
	// storage.InMemoryDataset
}

// TODO use pgx
type PostgresDatastore struct {
	// conn/dbpool 
}
