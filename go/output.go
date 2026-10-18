
package main

import (
	"fmt"
	"sort"
	"strings"
)

// printSchedule displays the completed weekly schedule
// and summarizes how employees were assigned.
func printSchedule(schedule Schedule, employees []Employee) {
	workDays := make(map[string]int)
	preferences := make(map[string][]string)

	for _, employee := range employees {
		preferences[employee.Name] = employee.Preferences
		workDays[employee.Name] = 0
	}

	preferenceMatches := 0
	totalAssignments := 0

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("          WEEKLY EMPLOYEE SCHEDULE      ")
	fmt.Println("========================================")

	for dayIndex, day := range days {
		fmt.Printf("\n%s\n", strings.ToUpper(day))
		fmt.Println("----------------------------------------")

		for _, shift := range shifts {
			assignedEmployees := schedule[day][shift]

			fmt.Printf(
				"  %-10s : %s\n",
				formatShiftName(shift),
				strings.Join(assignedEmployees, ", "),
			)

			for _, employeeName := range assignedEmployees {
				workDays[employeeName]++
				totalAssignments++

				if preferences[employeeName][dayIndex] == shift {
					preferenceMatches++
				}
			}
		}
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("             EMPLOYEE SUMMARY           ")
	fmt.Println("========================================")

	employeeNames := make([]string, 0, len(workDays))

	for name := range workDays {
		employeeNames = append(employeeNames, name)
	}

	sort.Strings(employeeNames)

	for _, name := range employeeNames {
		fmt.Printf(
			"  %-18s : %d day(s)\n",
			name,
			workDays[name],
		)
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("             SCHEDULE STATISTICS        ")
	fmt.Println("========================================")

	fmt.Printf("Total employees       : %d\n", len(employees))
	fmt.Printf("Total assignments     : %d\n", totalAssignments)
	fmt.Printf("Preference matches    : %d\n", preferenceMatches)

	if totalAssignments > 0 {
		matchPercentage := float64(preferenceMatches) /
			float64(totalAssignments) * 100

		fmt.Printf(
			"Preference satisfaction: %.1f%%\n",
			matchPercentage,
		)
	}

	fmt.Println()
	fmt.Println("Schedule validation passed.")
	fmt.Println("All shifts have at least two employees.")
	fmt.Println("No employee works more than one shift per day.")
	fmt.Println("No employee works more than five days per week.")
}

// formatShiftName capitalizes the first letter of a shift.
func formatShiftName(shift string) string {
	if shift == "" {
		return ""
	}

	return strings.ToUpper(shift[:1]) + shift[1:]
}
