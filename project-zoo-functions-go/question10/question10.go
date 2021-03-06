package question10

import (
	"project-zoo-functions-go/getdata"
)

//EmployeeCoverage represents coverage information of the employees.
type EmployeeCoverage struct {
	ID        string
	FullName  string
	Species   []string
	Locations []string
}

func makeEmployeeCoverage(employee getdata.Employee, allSpecies []getdata.Specie) EmployeeCoverage {
	fullName := employee.FirstName + " " + employee.LastName
	id := employee.ID
	species := make([]string, 0, 10)
	locations := make([]string, 0, 4)
	for _, vRF := range employee.ResponsibleFor {
		for _, vS := range allSpecies {
			if vS.ID == vRF {
				species = append(species, vS.Name)
				locations = append(locations, vS.Location)

			}
		}
	}
	var employeeCoverage = EmployeeCoverage{}
	employeeCoverage.ID = id
	employeeCoverage.FullName = fullName
	employeeCoverage.Species = species
	employeeCoverage.Locations = locations
	return employeeCoverage
}

//GetEmployeesCoverage This function will be responsible for associating coverage information of the employees.`
func GetEmployeesCoverage(args string) interface{} {
	data := getdata.GetZooData()
	var employeeCoverage = EmployeeCoverage{}
	var result interface{}
	allEmployeesCoverage := make([]EmployeeCoverage, 0, 20)
	for _, v := range data.Employees {
		employeeCoverage = makeEmployeeCoverage(v, data.Species)
		if args == "" {
			allEmployeesCoverage = append(allEmployeesCoverage, employeeCoverage)
		} else if args == v.ID || args == v.FirstName || args == v.LastName {
			allEmployeesCoverage = append(allEmployeesCoverage, employeeCoverage)
			break
		}

	}
	if len(allEmployeesCoverage) == 0 {
		return "Invalid information"
	} else if len(allEmployeesCoverage) == 1 {
		return employeeCoverage
	}
	result = allEmployeesCoverage
	return result
}
