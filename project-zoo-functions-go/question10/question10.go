package question10

import (
	"project-zoo-functions-go/getData"
	"sort"
)

type EmployeeCoverage struct {
	Id        string
	FullName  string
	Species   []string
	Locations []string
}

func makeEmployeeCoverage(employee getData.Employee, allSpecies []getData.Specie) EmployeeCoverage {
	fullName := employee.FirstName + " " + employee.LastName
	id := employee.ID
	species := make([]string, 0, 10)
	locations := make([]string, 0, 4)
	for _, vRF := range employee.ResponsibleFor {
		for _, vS := range allSpecies {
			if vS.ID == vRF {
				species = append(species, vS.Name)
				sort.Strings(locations)
				pos := sort.SearchStrings(locations, vS.Location)
				if pos == len(locations) {
					locations = append(locations, vS.Location)
				}
			}
		}
	}
	var employeeCoverage = EmployeeCoverage{}
	employeeCoverage.Id = id
	employeeCoverage.FullName = fullName
	employeeCoverage.Species = species
	employeeCoverage.Locations = locations
	return employeeCoverage
}

func GetEmployeesCoverage(args string) interface{} {
	data := getData.GetZooData()
	var employeeCoverage = EmployeeCoverage{}
	var result interface{}
	allEmployeesCoverage := make([]EmployeeCoverage, 0, 20)
	for _, v := range data.Employees {
		employeeCoverage = makeEmployeeCoverage(v, data.Species)
		if args == "" {
			allEmployeesCoverage = append(allEmployeesCoverage, employeeCoverage)
		} else {
			if args == v.ID || args == v.FirstName || args == v.LastName {
				allEmployeesCoverage = append(allEmployeesCoverage, employeeCoverage)
				break
			}
		}
	}
	if args == "" {
		result = allEmployeesCoverage
	} else {
		result = allEmployeesCoverage[0]
	}
	return result
}
