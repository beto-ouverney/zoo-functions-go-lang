package question1

import "project-zoo-functions-go/getdata"

//GetSpeciesByIds is responsible for searching animal species by id
func GetSpeciesByIds(args []string) []getdata.Specie {

	result := getdata.GetZooData()

	var ids []string
	ids = append(ids, args...)

	var results = make([]getdata.Specie, 0, len(result.Species))

	for _, vID := range ids {
		for _, vS := range result.Species {
			if vID == vS.ID {
				results = append(results, vS)
			}
		}
	}
	return results
}
