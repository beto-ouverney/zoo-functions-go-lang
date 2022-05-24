package question8

import (
	"project-zoo-functions-go/getData"
	"sort"
	"strconv"
)

type scheduleDay struct {
	officeHour  string
	exihibition []string
}

func makeExihibition(paramSpecies []getData.Specie, day string) []string {
	exihibition := make([]string, 0, 10)
	for _, vS := range paramSpecies {
		sort.Strings(vS.Availability)
		pos := sort.SearchStrings(vS.Availability, day)
		if pos < len(vS.Availability) {
			exihibition = append(exihibition, vS.Name)
		}

	}
	return exihibition
}

func maheSchedule(hours getData.Hours, paramSpecies []getData.Specie, dayNames []string) map[string]interface{} {
	schedule := make(map[string]interface{})
	for _, v := range dayNames {
		var scheduleDay scheduleDay
		if v == "Tuesday" {
			scheduleDay.officeHour = "Open from " + strconv.Itoa(hours.Tuesday.Open) + "am until " + strconv.Itoa(hours.Tuesday.Close)
			scheduleDay.exihibition = makeExihibition(paramSpecies, v)
			schedule["Tuesday"] = scheduleDay
		} else if v == "Wednesday" {
			scheduleDay.officeHour = "Open from " + strconv.Itoa(hours.Wednesday.Open) + "am until " + strconv.Itoa(hours.Wednesday.Close)
			scheduleDay.exihibition = makeExihibition(paramSpecies, v)
			schedule["Wednesday"] = scheduleDay
		} else if v == "Thursday" {
			scheduleDay.officeHour = "Open from " + strconv.Itoa(hours.Thursday.Open) + "am until " + strconv.Itoa(hours.Thursday.Close)
			scheduleDay.exihibition = makeExihibition(paramSpecies, v)
			schedule["Tursday"] = scheduleDay
		} else if v == "Friday" {
			scheduleDay.officeHour = "Open from " + strconv.Itoa(hours.Friday.Open) + "am until " + strconv.Itoa(hours.Friday.Close)
			scheduleDay.exihibition = makeExihibition(paramSpecies, v)
			schedule["Friday"] = scheduleDay
		} else if v == "Saturday" {
			scheduleDay.officeHour = "Open from " + strconv.Itoa(hours.Saturday.Open) + "am until " + strconv.Itoa(hours.Saturday.Close)
			scheduleDay.exihibition = makeExihibition(paramSpecies, v)
			schedule["Saturday"] = scheduleDay
		} else if v == "Sunday" {
			scheduleDay.officeHour = "Open from " + strconv.Itoa(hours.Sunday.Open) + "am until " + strconv.Itoa(hours.Sunday.Close)
			scheduleDay.exihibition = makeExihibition(paramSpecies, v)
			schedule["Sunday"] = scheduleDay
		} else {
			monday := struct {
				officeHour  string
				exihibition string
			}{
				officeHour:  "CLOSED",
				exihibition: "The zoo will be closed!",
			}
			schedule["Monday"] = monday
		}

	}
	return schedule
}

func GetSchedule(args string) interface{} {
	dayNames := []string{"Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday", "Monday"}
	sort.Strings(dayNames)
	data := getData.GetZooData()
	if args == "" {
		return maheSchedule(data.Hours, data.Species, dayNames)
	}
	pos := sort.SearchStrings(dayNames, args)
	if pos < len(dayNames) {
		param := make([]string, 0, 1)
		param = append(param, args)
		return maheSchedule(data.Hours, data.Species, param)
	}

	for _, v := range data.Species {
		if v.Name == args {
			return v.Availability
		}
	}

	return maheSchedule(data.Hours, data.Species, dayNames)
}
