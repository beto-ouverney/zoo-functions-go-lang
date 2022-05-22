package question4

import (
	"project-zoo-functions-go/getData"
)

func GetRelatedEmployees(id string) []string {
	data := getData.GetZooData()

	result := make([]string, 0, len(data.Employees))

	for _, vEmployee := range data.Employees {
		for _, vEmployeeManager := range vEmployee.Managers {
			if vEmployeeManager == id {
				result = append(result, vEmployee.FirstName+" "+vEmployee.LastName)
			}

		}
	}

	if len(result) == 0 {
		result = append(result, "O id inserido não é de uma pessoa colaboradora gerente!")
	}
	return result
}
