package main

import (
	"project-zoo-functions-go/getdata"
	"project-zoo-functions-go/question1"
	"project-zoo-functions-go/question2"
	"project-zoo-functions-go/question3"
	"project-zoo-functions-go/question4"
	"project-zoo-functions-go/question5"
	"project-zoo-functions-go/question6"
	"project-zoo-functions-go/question7"
	"project-zoo-functions-go/question9"

	"project-zoo-functions-go/question8"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSpeciesByIDs(t *testing.T) {
	t.Log("Testing Question 1!")
	type test struct {
		data   []string
		answer []getdata.Specie
	}
	tests := []test{
		{
			data:   []string{""},
			answer: []getdata.Specie{},
		},
		{
			data: []string{
				"0938aa23-f153-4937-9f88-4858b24d6bce",
			},
			answer: []getdata.Specie{
				{
					ID:           "0938aa23-f153-4937-9f88-4858b24d6bce",
					Name:         "lions",
					Popularity:   4,
					Location:     "NE",
					Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
					Residents: []getdata.Resident{
						{Name: "Zena", Sex: "female", Age: 12},
						{Name: "Maxwell", Sex: "male", Age: 15},
						{Name: "Faustino", Sex: "male", Age: 7},
						{Name: "Dee", Sex: "female", Age: 14},
					},
				},
			},
		}, {
			data: []string{"0938aa23-f153-4937-9f88-4858b24d6bce",
				"e8481c1d-42ea-4610-8e11-1752cfc05a46"},
			answer: []getdata.Specie{
				{
					ID:           "0938aa23-f153-4937-9f88-4858b24d6bce",
					Name:         "lions",
					Popularity:   4,
					Location:     "NE",
					Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
					Residents: []getdata.Resident{
						{Name: "Zena", Sex: "female", Age: 12},
						{Name: "Maxwell", Sex: "male", Age: 15},
						{Name: "Faustino", Sex: "male", Age: 7},
						{Name: "Dee", Sex: "female", Age: 14},
					},
				},
				{
					ID:           "e8481c1d-42ea-4610-8e11-1752cfc05a46",
					Name:         "tigers",
					Popularity:   5,
					Location:     "NW",
					Availability: []string{"Wednesday", "Friday", "Saturday", "Tuesday"},
					Residents: []getdata.Resident{
						{Name: "Shu", Sex: "female", Age: 19},
						{Name: "Esther", Sex: "female", Age: 17},
					},
				},
			},
		},
	}

	for _, v := range tests {
		for i := 0; i < len(v.data); i++ {
			actual := question1.GetSpeciesByIds(v.data)
			assert.Equal(t, actual, v.answer)
		}
	}
}

func TestGetAnimalsOlderThan(t *testing.T) {
	t.Log("Testing Question 2!")
	actual := question2.GetAnimalsOlderThan("lions", 10)
	assert.Equal(t, actual, true)

	actual = question2.GetAnimalsOlderThan("penguins", 10)
	assert.Equal(t, actual, true)
}

func TestGetEmployeeByName(t *testing.T) {
	t.Log("Testing Question 3!")
	lionID := "0938aa23-f153-4937-9f88-4858b24d6bce"
	elephantsID := "bb2a76d8-5fe3-4d03-84b7-dba9cfc048b5"
	bearsID := "baa6e93a-f295-44e7-8f70-2bcdc6f6948d"
	snakesID := "78460a91-f4da-4dea-a469-86fd2b8ccc84"

	actual := question3.GetEmployeeByName("")
	assert.Equal(t, actual, getdata.Employee{})

	test := []struct {
		data   string
		answer getdata.Employee
	}{
		{
			data: "Emery",
			answer: getdata.Employee{
				ID:        "b0dc644a-5335-489b-8a2c-4e086c7819a2",
				FirstName: "Emery",
				LastName:  "Elser",
				Managers:  []string{"9e7d4524-363c-416a-8759-8aa7e50c0992"},
				ResponsibleFor: []string{
					lionID,
					bearsID,
					elephantsID,
				},
			},
		},
		{
			data: "Wishart",
			answer: getdata.Employee{
				ID:        "56d43ba3-a5a7-40f6-8dd7-cbb05082383f",
				FirstName: "Wilburn",
				LastName:  "Wishart",
				Managers: []string{
					"0e7b460e-acf4-4e17-bcb3-ee472265db83",
					"fdb2543b-5662-46a7-badc-93d960fdc0a8",
				},
				ResponsibleFor: []string{
					snakesID,
					elephantsID,
				},
			},
		},
	}

	for _, v := range test {
		actual := question3.GetEmployeeByName(v.data)
		t.Log(actual, v.answer)
		assert.Equal(t, actual, v.answer)
	}
}

func TestGetRelatedEmployees(t *testing.T) {
	t.Log("Testing Question 4!")
	const stephanieID = "9e7d4524-363c-416a-8759-8aa7e50c0992"

	actual := question4.GetRelatedEmployees("4b40a139-d4dc-4f09-822d-ec25e819a5ad")
	assert.Equal(t, actual, "The id entered is not that of a collaborating person manager!")

	actual = question4.GetRelatedEmployees(stephanieID)
	assert.Equal(t, actual, []string{"Burl Bethea", "Ola Orloff", "Emery Elser"})
}

func TestCountAnimals(t *testing.T) {
	t.Log("Testing Question 5!")
	param := question5.ParamTypeQ5{}
	actual := question5.CountAnimals(param)
	assert.Equal(t, actual, 31)

	test := struct {
		data   []question5.ParamTypeQ5
		answer []int
	}{
		data: []question5.ParamTypeQ5{
			{Specie: "penguins", Sex: ""},
			{Specie: "giraffes", Sex: ""},
			{Specie: "bears", Sex: "female"},
			{Specie: "elephants", Sex: "male"},
		},
		answer: []int{4, 6, 0, 2},
	}

	for i := 0; i < len(test.data); i++ {
		actual := question5.CountAnimals(test.data[i])
		assert.Equal(t, actual, test.answer[i])
	}
}

func TestCalculateEntry(t *testing.T) {
	t.Log("Testing Question 6!")
	data := getdata.GetEntrants()

	t.Log("Testing Question 6 - Function CountEntrants !")
	actual := question6.CountEntrants(data.Entrants)
	expected := question6.EntrantsType{Adult: 2, Child: 3, Senior: 1}
	assert.Equal(t, actual, expected)

	t.Log("Testing Question 6 - Function CalculateEntry")
	emptyData := []getdata.Entrant{}
	actual2 := question6.CalculateEntry(emptyData)
	assert.Equal(t, actual2, 0.0)

	actual3 := question6.CalculateEntry(data.Entrants)
	assert.Equal(t, actual3, 187.94)

	newTestData := []getdata.Entrant{data.Entrants[3]}
	actual3 = question6.CalculateEntry(newTestData)
	assert.Equal(t, actual3, 49.99)

	newTestData = []getdata.Entrant{data.Entrants[5]}
	actual3 = question6.CalculateEntry(newTestData)
	assert.Equal(t, actual3, 24.99)

	newTestData = []getdata.Entrant{data.Entrants[0]}
	actual3 = question6.CalculateEntry(newTestData)
	assert.Equal(t, actual3, 20.99)

	newTestData = []getdata.Entrant{data.Entrants[0], data.Entrants[5]}
	actual3 = question6.CalculateEntry(newTestData)
	assert.Equal(t, actual3, 45.98)

}

func TestGetAnimalMap(t *testing.T) {
	t.Log("Testing Question 7!")
	expected := question7.AnimalsResult{
		NE: []interface{}{"lions", "giraffes"},
		NW: []interface{}{"tigers", "bears", "elephants"},
		SE: []interface{}{"penguins", "otters"},
		SW: []interface{}{"frogs", "snakes"},
	}

	input := question7.Options{}
	actual := question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected)
	input = question7.Options{Sex: "female"}
	actual = question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected)

	input = question7.Options{Sex: "female", Sorted: true}
	actual = question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected)

	expected2 := question7.AnimalsResult{
		NE: []interface{}{
			map[string]interface{}{"lions": []string{"Zena", "Maxwell", "Faustino", "Dee"}},
			map[string]interface{}{"giraffes": []string{"Gracia", "Antone", "Vicky", "Clay", "Arron", "Bernard"}},
		},
		NW: []interface{}{
			map[string]interface{}{"tigers": []string{"Shu", "Esther"}},
			map[string]interface{}{"bears": []string{"Hiram", "Edwardo", "Milan"}},
			map[string]interface{}{"elephants": []string{"Ilana", "Orval", "Bea", "Jefferson"}},
		},
		SE: []interface{}{
			map[string]interface{}{"penguins": []string{"Joe", "Tad", "Keri", "Nicholas"}},
			map[string]interface{}{"otters": []string{"Neville", "Lloyd", "Mercedes", "Margherita"}},
		},
		SW: []interface{}{
			map[string]interface{}{"frogs": []string{"Cathey", "Annice"}},
			map[string]interface{}{"snakes": []string{"Paulette", "Bill"}},
		},
	}
	input = question7.Options{IncludeNames: true}
	actual = question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected2)

	expected2 = question7.AnimalsResult{
		NE: []interface{}{
			map[string]interface{}{"lions": []string{"Dee", "Faustino", "Maxwell", "Zena"}},
			map[string]interface{}{"giraffes": []string{"Antone", "Arron", "Bernard", "Clay", "Gracia", "Vicky"}},
		},
		NW: []interface{}{
			map[string]interface{}{"tigers": []string{"Esther", "Shu"}},
			map[string]interface{}{"bears": []string{"Edwardo", "Hiram", "Milan"}},
			map[string]interface{}{"elephants": []string{"Bea", "Ilana", "Jefferson", "Orval"}},
		},
		SE: []interface{}{
			map[string]interface{}{"penguins": []string{"Joe", "Keri", "Nicholas", "Tad"}},
			map[string]interface{}{"otters": []string{"Lloyd", "Margherita", "Mercedes", "Neville"}},
		},
		SW: []interface{}{
			map[string]interface{}{"frogs": []string{"Annice", "Cathey"}},
			map[string]interface{}{"snakes": []string{"Bill", "Paulette"}},
		},
	}
	input = question7.Options{IncludeNames: true, Sorted: true}
	actual = question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected2)

	expected2 = question7.AnimalsResult{
		NE: []interface{}{
			map[string]interface{}{"lions": []string{"Zena", "Dee"}},
			map[string]interface{}{"giraffes": []string{"Gracia", "Vicky"}},
		},
		NW: []interface{}{
			map[string]interface{}{"tigers": []string{"Shu", "Esther"}},
			map[string]interface{}{"bears": []string{}},
			map[string]interface{}{"elephants": []string{"Ilana", "Bea"}},
		},
		SE: []interface{}{
			map[string]interface{}{"penguins": []string{"Keri"}},
			map[string]interface{}{"otters": []string{"Mercedes", "Margherita"}},
		},
		SW: []interface{}{
			map[string]interface{}{"frogs": []string{"Cathey", "Annice"}},
			map[string]interface{}{"snakes": []string{"Paulette"}},
		},
	}
	input = question7.Options{IncludeNames: true, Sex: "female"}
	actual = question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected2)

	expected2 = question7.AnimalsResult{
		NE: []interface{}{
			map[string]interface{}{"lions": []string{"Dee", "Zena"}},
			map[string]interface{}{"giraffes": []string{"Gracia", "Vicky"}},
		},
		NW: []interface{}{
			map[string]interface{}{"tigers": []string{"Esther", "Shu"}},
			map[string]interface{}{"bears": []string{}},
			map[string]interface{}{"elephants": []string{"Bea", "Ilana"}},
		},
		SE: []interface{}{
			map[string]interface{}{"penguins": []string{"Keri"}},
			map[string]interface{}{"otters": []string{"Margherita", "Mercedes"}},
		},
		SW: []interface{}{
			map[string]interface{}{"frogs": []string{"Annice", "Cathey"}},
			map[string]interface{}{"snakes": []string{"Paulette"}},
		},
	}
	input = question7.Options{IncludeNames: true, Sex: "female", Sorted: true}
	actual = question7.GetAnimalMap(input)
	assert.Equal(t, actual, expected2)
}

func TestGetSchedule(t *testing.T) {

	monday := struct {
		OfficeHour  string
		Exihibition string
	}{
		OfficeHour:  "CLOSED",
		Exihibition: "The zoo will be closed!",
	}
	expected := map[string]interface{}{
		"Tuesday": question8.ScheduleDay{
			OfficeHour:  "Open from 8am until 6pm",
			Exihibition: []string{"lions", "tigers", "bears", "penguins", "elephants", "giraffes"},
		},
		"Wednesday": question8.ScheduleDay{
			OfficeHour:  "Open from 8am until 6pm",
			Exihibition: []string{"tigers", "bears", "penguins", "otters", "frogs", "giraffes"},
		},
		"Thursday": question8.ScheduleDay{
			OfficeHour:  "Open from 10am until 8pm",
			Exihibition: []string{"lions", "otters", "frogs", "snakes", "giraffes"},
		},
		"Friday": question8.ScheduleDay{
			OfficeHour:  "Open from 10am until 8pm",
			Exihibition: []string{"tigers", "otters", "frogs", "snakes", "elephants", "giraffes"},
		},
		"Saturday": question8.ScheduleDay{
			OfficeHour: "Open from 8am until 10pm",
			Exihibition: []string{
				"lions", "tigers",
				"bears", "penguins",
				"otters", "frogs",
				"snakes", "elephants",
			},
		},
		"Sunday": question8.ScheduleDay{
			OfficeHour:  "Open from 8am until 8pm",
			Exihibition: []string{"lions", "bears", "penguins", "snakes", "elephants"},
		},
		"Monday": monday,
	}
	actual := question8.GetSchedule("")
	assert.Equal(t, actual, expected)

	actual = question8.GetSchedule("adjkdjakjdkajkankc")
	assert.Equal(t, actual, expected)

	expected = map[string]interface{}{
		"Monday": monday,
	}
	actual = question8.GetSchedule("Monday")
	assert.Equal(t, actual, expected)

	expected = map[string]interface{}{
		"Tuesday": question8.ScheduleDay{
			OfficeHour:  "Open from 8am until 6pm",
			Exihibition: []string{"lions", "tigers", "bears", "penguins", "elephants", "giraffes"},
		},
	}

	actual = question8.GetSchedule("Tuesday")
	assert.Equal(t, actual, expected)

	expected = map[string]interface{}{
		"Wednesday": question8.ScheduleDay{
			OfficeHour:  "Open from 8am until 6pm",
			Exihibition: []string{"tigers", "bears", "penguins", "otters", "frogs", "giraffes"},
		},
	}

	actual = question8.GetSchedule("Wednesday")
	assert.Equal(t, actual, expected)

	expected2 := []string{"Tuesday", "Thursday", "Saturday", "Sunday"}
	actual = question8.GetSchedule("lions")
	assert.Equal(t, actual, expected2)

	expected2 = []string{"Tuesday", "Wednesday", "Sunday", "Saturday"}
	actual = question8.GetSchedule("penguins")
	assert.Equal(t, actual, expected2)

}

func TestGetOldestFromFirstSpecies(t *testing.T) {
	test := struct {
		data   []string
		answer []question9.Response
	}{
		data: []string{
			"c5b83cb3-a451-49e2-ac45-ff3f54fbe7e1",
			"0e7b460e-acf4-4e17-bcb3-ee472265db83",
			"fdb2543b-5662-46a7-badc-93d960fdc0a8",
			"56d43ba3-a5a7-40f6-8dd7-cbb05082383f",
			"9e7d4524-363c-416a-8759-8aa7e50c0992",
			"4b40a139-d4dc-4f09-822d-ec25e819a5ad",
			"c1f50212-35a6-4ecd-8223-f835538526c2",
			"b0dc644a-5335-489b-8a2c-4e086c7819a2",
		},
		answer: []question9.Response{
			{Name: "Maxwell", Sex: "male", Age: 15},
			{Name: "Maxwell", Sex: "male", Age: 15},
			{Name: "Margherita", Sex: "female", Age: 10},
			{Name: "Bill", Sex: "male", Age: 6},
			{Name: "Margherita", Sex: "female", Age: 10},
			{Name: "Margherita", Sex: "female", Age: 10},
			{Name: "Shu", Sex: "female", Age: 19},
			{Name: "Maxwell", Sex: "male", Age: 15},
		},
	}
	for i := 0; i < len(test.data); i++ {
		actual := question9.GetOldestFromFirstSpecies(test.data[i])
		assert.Equal(t, actual, test.answer[i])
	}
}
