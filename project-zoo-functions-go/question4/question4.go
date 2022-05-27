package question4

import (
	"project-zoo-functions-go/getdata"
)

//GetRelatedEmployees if it is a manager collaborating person, it must return an array containing the names of the collaborating people it is responsible for. if not a manager contributor, return the message "The id entered is not a manager contributor person!"
func GetRelatedEmployees(id string) []string {
	data := getdata.GetZooData()

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
