package foundry

type ancestry struct {
	name string
}

type ancestryFeature struct {
	name string
}

type background struct {
	name string
}

type class struct {
	name string
}

type classFeature struct {
	name string
}

type equipmentEffect struct {
	name string
}

// this should have subtypes
type equipment struct {
	name string
}

type feature struct {
	name string
}

type featureEffect struct {
	name string
}

type heritage struct {
	name string
}

type ModelType interface {
	ancestry | ancestryFeature | background | class | classFeature | equipmentEffect | equipment | feature | featureEffect | heritage
}
