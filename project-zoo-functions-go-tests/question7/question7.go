package question7

//AnimalsResult represents geographic mapping of the species and their animals,
type AnimalsResult struct {
	NE []interface{}
	NW []interface{}
	SE []interface{}
	SW []interface{}
}

//Options represents GetAnimalMap functions input parameters
type Options struct {
	IncludeNames bool
	Sex          string
	Sorted       bool
}

//GetAnimalMap is responsible for the geographic mapping of the species and their animals, and can also filter them by alphabetical order and sex.
func GetAnimalMap(options Options) AnimalsResult {
	return AnimalsResult{}
}
