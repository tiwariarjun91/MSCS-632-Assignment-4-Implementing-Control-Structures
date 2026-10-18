package main


package main

// Employee represents a team member and their preferred
// shift for each day of the week.
type Employee struct {
	Name        string   `json:"name"`
	Preferences []string `json:"preferences"`
}

// Schedule stores the employees assigned to each shift.
// The first key represents the day, and the second the shift.
type Schedule map[string]map[string][]string

var days = []string{
	"Monday",
	"Tuesday",
	"Wednesday",
	"Thursday",
	"Friday",
	"Saturday",
	"Sunday",
}

var shifts = []string{
	"morning",
	"afternoon",
	"evening",
}
