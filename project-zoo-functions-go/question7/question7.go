package question7

import (
	"project-zoo-functions-go/getdata"
	"sort"
)

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
	data := getdata.GetZooData()
	var result AnimalsResult
	for _, v := range data.Species {
		temp := make([]string, 0, 10)
		temp2 := make(map[string]interface{})
		var value interface{}

		if options.IncludeNames {
			for _, vR := range v.Residents {
				if options.Sex != "" {
					if options.Sex == vR.Sex {
						temp = append(temp, vR.Name)
					}
				} else {
					temp = append(temp, vR.Name)
				}
			}
			if options.Sorted {
				sort.Strings(temp)
			}
			temp2[v.Name] = temp
			value = temp2
		} else {
			value = v.Name
		}
		switch v.Location {
		case "NE":
			result.NE = append(result.NE, value)
		case "NW":
			result.NW = append(result.NW, value)
		case "SE":
			result.SE = append(result.SE, value)
		case "SW":
			result.SW = append(result.SW, value)
		}
	}
	return result
}
