package getdata

import (
	"encoding/json"
	"fmt"
	"os"
)

//Employee represents zoo employee
type Employee struct {
	ID             string   `json:"id"`
	FirstName      string   `json:"firstName"`
	LastName       string   `json:"lastName"`
	Managers       []string `json:"managers"`
	ResponsibleFor []string `json:"responsibleFor"`
}

//Specie represents a animal specie with your data
type Specie struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Popularity   int      `json:"popularity"`
	Location     string   `json:"location"`
	Availability []string `json:"availability"`
	Residents    []struct {
		Name string `json:"name"`
		Sex  string `json:"sex"`
		Age  int    `json:"age"`
	} `json:"residents"`
}

//Hours represents opener and closer hours of the zoo
type Hours struct {
	Tuesday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Tuesday"`
	Wednesday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Wednesday"`
	Thursday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Thursday"`
	Friday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Friday"`
	Saturday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Saturday"`
	Sunday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Sunday"`
	Monday struct {
		Open  int `json:"open"`
		Close int `json:"close"`
	} `json:"Monday"`
}

//Species represents the response from a mock ZooDATA
type Species struct {
	Species   []Specie   `json:"species"`
	Employees []Employee `json:"employees"`
	Hours     Hours      `json:"hours"`
	Prices    struct {
		Adult  float64 `json:"adult"`
		Senior float64 `json:"senior"`
		Child  float64 `json:"child"`
	} `json:"prices"`
}

//Entrant represent a costumer in the Zoo
type Entrant struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//Entrants represents a list of costumers in the zoo
type Entrants struct {
	Entrants []Entrant
}

//GetZooData is the function get all zoo data
func GetZooData() Species {
	jsonFile, err := os.ReadFile("./data/output.json")
	if err != nil {
		fmt.Print(err)
	}
	var jsonData Species
	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		fmt.Print(err)
	}
	return jsonData
}

//GetEntrants get all zoo costumers
func GetEntrants() Entrants {

	jsonFile, err := os.ReadFile("./data/entrants.json")
	if err != nil {
		fmt.Print(err)
	}
	var jsonData Entrants
	err = json.Unmarshal(jsonFile, &jsonData)
	if err != nil {
		fmt.Print(err)
	}

	return jsonData

}
