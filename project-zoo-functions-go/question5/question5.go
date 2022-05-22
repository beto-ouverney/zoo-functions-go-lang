package question5

import "project-zoo-functions-go/getData"

type ParamTypeQ5 struct {
	Specie string
	Sex    string
}

func CountAnimals(param ParamTypeQ5) int {
	data := getData.ReturnStruct()
	result := 0
	if param.Specie == "" {
		for _, v := range data.Species {
			for i := 0; i < len(v.Residents); i++ {
				result++
			}
		}
	} else if param.Sex == "" {
		for _, v := range data.Species {
			if v.Name == param.Specie {
				for i := 0; i < len(v.Residents); i++ {
					result++
				}
			}
		}
	} else {
		for _, v := range data.Species {
			if v.Name == param.Specie {
				for _, vResidents := range v.Residents {
					if vResidents.Sex == param.Specie && vResidents.Sex == param.Sex {
						result++
					}
				}
			}
		}
	}
	return result
}
