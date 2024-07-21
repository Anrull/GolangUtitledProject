package lexicon

import (
	"encoding/json"
	"log"
	"os"
)

var (
	OlimpListStep  = 10
	OlimpListLeft  = "⬅️"
	OlimpListRight = "➡️"
)

var HelpMessage string

var InformationCaption string

var TimetableTime string

var SubjectsForButton []string

var ListDays []string
var Stages []string

var Week map[string]string
var DayTextToInt map[string]string
var Day map[string]string

var StagesTracker []string

var TeacherTracker []string

var TrackerOlimps []string


func init() {
	data, err := os.ReadFile("data/lexicon.json")
	if err != nil {
		log.Fatal(err)
	}

	var values struct {
		HelpMessage         string `json:"help_message"`
		InformationCaption  string `json:"information_caption"`
		TimetableTime        string    `json:"timetable_time"`
		SubjectsForButton    []string  `json:"subjects"`
		ListDays               []string `json:"list_days"`
		Stages               []string `json:"classes"`
		Week                 map[string]string `json:"week"`
		DayTextToInt         map[string]string `json:"days_to_int"`
		Days                  map[string]string `json:"int_to_days"`
		StagesTracker        []string `json:"stages_tracker"`
		TeacherTracker       []string `json:"teachers_tracker"`
		TrackerOlimps        []string `json:"tracker_olimps"`
	}

	if err := json.Unmarshal(data, &values); err != nil {
		log.Fatal(err)
	}

	HelpMessage = values.HelpMessage
	InformationCaption = values.InformationCaption
	TimetableTime = values.TimetableTime
	SubjectsForButton = values.SubjectsForButton
	ListDays = values.ListDays
	Stages = values.Stages
	Week = values.Week
	DayTextToInt = values.DayTextToInt
	Day = values.Days
	StagesTracker = values.StagesTracker
	TeacherTracker = values.TeacherTracker
	TrackerOlimps = values.TrackerOlimps
}