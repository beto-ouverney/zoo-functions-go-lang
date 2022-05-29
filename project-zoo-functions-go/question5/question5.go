package question5

import "project-zoo-functions-go/getdata"

//ParamTypeQ5 represents CountAnimals input parameters, with animal specie and animal sex
type ParamTypeQ5 struct {
	Specie string
	Sex    string
}

//CountAnimals  is responsible for counting the number of animals of each species.
func CountAnimals(param ParamTypeQ5) int {
	data := getdata.GetZooData()
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
					if vResidents.Sex == param.Sex {
						result++
					}
				}
			}
		}
	}
	return result
}
