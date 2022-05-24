package question1

import "project-zoo-functions-go/getData"

func GetSpeciesByIds(args []string) []getData.Specie {

	result := getData.GetZooData()

	var ids []string
	ids = append(ids, args...)

	var results = make([]getData.Specie, 0, len(result.Species))

	for _, vId := range ids {
		for _, vS := range result.Species {
			if vId == vS.ID {
				results = append(results, vS)
			}
		}
	}
	return results
}
