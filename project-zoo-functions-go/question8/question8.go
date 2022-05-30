package question8

import (
	"project-zoo-functions-go/getdata"
	"strconv"
)

//ScheduleDay represents Zoo day schedule with open and close hours and the species exihibition
type ScheduleDay struct {
	OfficeHour  string
	Exihibition []string
}

func makeExihibition(paramSpecies []getdata.Specie, day string) []string {
	exihibition := make([]string, 0, 10)
	for _, vS := range paramSpecies {
		for _, vD := range vS.Availability {
			if day == vD {
				exihibition = append(exihibition, vS.Name)
			}
		}
	}
	return exihibition
}

func maheSchedule(hours getdata.Hours, paramSpecies []getdata.Specie, dayNamesParam []string) map[string]interface{} {
	schedule := make(map[string]interface{})
	for _, v := range dayNamesParam {
		var scheduleDay ScheduleDay
		switch v {
		case "Tuesday":
			scheduleDay.OfficeHour = "Open from " + strconv.Itoa(hours.Tuesday.Open) + "am until " + strconv.Itoa(hours.Tuesday.Close) + "pm"
			scheduleDay.Exihibition = makeExihibition(paramSpecies, v)
			schedule["Tuesday"] = scheduleDay
		case "Wednesday":
			scheduleDay.OfficeHour = "Open from " + strconv.Itoa(hours.Wednesday.Open) + "am until " + strconv.Itoa(hours.Wednesday.Close) + "pm"
			scheduleDay.Exihibition = makeExihibition(paramSpecies, v)
			schedule["Wednesday"] = scheduleDay
		case "Thursday":
			scheduleDay.OfficeHour = "Open from " + strconv.Itoa(hours.Thursday.Open) + "am until " + strconv.Itoa(hours.Thursday.Close) + "pm"
			scheduleDay.Exihibition = makeExihibition(paramSpecies, v)
			schedule["Thursday"] = scheduleDay
		case "Friday":
			scheduleDay.OfficeHour = "Open from " + strconv.Itoa(hours.Friday.Open) + "am until " + strconv.Itoa(hours.Friday.Close) + "pm"
			scheduleDay.Exihibition = makeExihibition(paramSpecies, v)
			schedule["Friday"] = scheduleDay
		case "Saturday":
			scheduleDay.OfficeHour = "Open from " + strconv.Itoa(hours.Saturday.Open) + "am until " + strconv.Itoa(hours.Saturday.Close) + "pm"
			scheduleDay.Exihibition = makeExihibition(paramSpecies, v)
			schedule["Saturday"] = scheduleDay
		case "Sunday":
			scheduleDay.OfficeHour = "Open from " + strconv.Itoa(hours.Sunday.Open) + "am until " + strconv.Itoa(hours.Sunday.Close) + "pm"
			scheduleDay.Exihibition = makeExihibition(paramSpecies, v)
			schedule["Sunday"] = scheduleDay
		case "Monday":
			monday := struct {
				OfficeHour  string
				Exihibition string
			}{
				OfficeHour:  "CLOSED",
				Exihibition: "The zoo will be closed!",
			}
			schedule["Monday"] = monday
		}

	}
	return schedule
}

//GetSchedule  is responsible for making the time information of the animals available in a query to the user, who may want to have access to the schedule of the week, a day or a specific animal.
func GetSchedule(args string) interface{} {
	dayNames := []string{"Friday", "Monday", "Saturday", "Sunday", "Thursday", "Tuesday", "Wednesday"}
	data := getdata.GetZooData()
	if args == "" {
		return maheSchedule(data.Hours, data.Species, dayNames)
	}
	for _, v := range dayNames {
		if v == args {
			param := make([]string, 0, 1)
			param = append(param, args)
			return maheSchedule(data.Hours, data.Species, param)
		}
	}
	for _, v := range data.Species {
		if v.Name == args {
			return v.Availability
		}
	}

	return maheSchedule(data.Hours, data.Species, dayNames)
}
