package question7

import "project-zoo-functions-go/getData"

var regions = []string{"NE", "NW", "SE", "SW"}

type animalsResult struct {
	NE interface{}
	NW interface{}
	SE interface{}
	SW interface{}
}

type Options struct {
	includeNames string
	sex          string
	sorted       bool
}

func defaultResult(param []getData.Specie, result animalsResult) {
	NE := make([]string, 0, len(param))
	NW := make([]string, 0, len(param))
	SE := make([]string, 0, len(param))
	SW := make([]string, 0, len(param))
	for _, v := range param {
		switch v.Location {
		case "NE":
			NE = append(NE, v.Name)
		case "NW":
			NW = append(NW, v.Name)
		case "SE":
			SE = append(SE, v.Name)
		case "SW":
			SW = append(SW, v.Name)
		}
	}
	result.NE = NE
	result.NW = NW
	result.SE = SE
	result.SW = SW
}

func sortNamesResult(param []getData.Specie) {
	NE := make([]string, 0, len(param))
	NW := make([]string, 0, len(param))
	SE := make([]string, 0, len(param))
	SW := make([]string, 0, len(param))
	for _, v := range param {
		switch v.Location {
		case "NE":
			NE = append(NE, v.Name)
		case "NW":
			NW = append(NW, v.Name)
		case "SE":
			SE = append(SE, v.Name)
		case "SW":
			SW = append(SW, v.Name)
		}
	}
	result.NE = NE
	result.NW = NW
	result.SE = SE
	result.SW = SW
}

func conditions(options Options, param []getData.Specie) animalsResult {
	if options.includeNames != "" {
		var result animalsResult
		defaultResult(param, result)
		return result
	}
	/*	if options.sex != "" {
			if options.sorted {
				return sortNamesResult()
			}
			return includeNamesResult()
		}
		const maleOrFemale = options.sex
		if options.sorted {
			return sexAndSortResult(maleOrFemale)
		}*/
	return nil
}
func GetAnimalMap(options Options) {
	data := getData.GetZooData()
	conditions(options, data.Species)

}
