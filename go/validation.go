package main

import (
	"errors"
	"fmt"
	"strings"
)

// validEmployees checks employee details before scheduling begins.
func validEmployees(employees []Employee) error {
	// Each employee can work at most five days.
	// We need 42 assignments, so at least nine employees are required.
	if len(employees) < 9 {
		return errors.New("at least 9 employees are required")
	}

	registeredNames := make(map[string]bool)

	for _, employee := range employees {
		name := strings.TrimSpace(employee.Name)

		if name == "" {
			return errors.New("employee name cannot be empty")
		}

		if registeredNames[name] {
			return fmt.Errorf("duplicate employee name: %s", name)
		}
		registeredNames[name] = true

		if len(employee.Preferences) != len(days) {
			return fmt.Errorf(
				"%s must provide a shift preference for all seven days",
				name,
			)
		}

		for dayIndex, preferredShift := range employee.Preferences {
			shiftIsValid := false

			for _, availableShift := range shifts {
				if preferredShift == availableShift {
					shiftIsValid = true
					break
				}
			}

			if !shiftIsValid {
				return fmt.Errorf(
					"%s has an invalid shift preference for %s",
					name,
					days[dayIndex],
				)
			}
		}
	}

	return nil
}

// Verify checks whether the generated weekly schedule follows
// the company's staffing and employee workload rules.
func Verify(schedule Schedule) error {
	daysWorked := make(map[string]int)

	for _, day := range days {
		assignedToday := make(map[string]bool)

		dailySchedule, exists := schedule[day]
		if !exists {
			return fmt.Errorf("missing schedule for %s", day)
		}

		for _, shift := range shifts {
			assignedEmployees, exists := dailySchedule[shift]
			if !exists {
				return fmt.Errorf("missing %s shift on %s", shift, day)
			}

			// Every shift needs at least two employees.
			if len(assignedEmployees) < 2 {
				return fmt.Errorf(
					"not enough employees for %s on %s",
					shift,
					day,
				)
			}

			for _, employeeName := range assignedEmployees {
				// An employee cannot work multiple shifts in one day.
				if assignedToday[employeeName] {
					return fmt.Errorf(
						"%s has multiple shifts on %s",
						employeeName,
						day,
					)
				}

				assignedToday[employeeName] = true
				daysWorked[employeeName]++
			}
		}
	}

	// Employees cannot work more than five days in a week.
	for employeeName, totalDays := range daysWorked {
		if totalDays > 5 {
			return fmt.Errorf(
				"%s is scheduled for %d days (maximum is 5)",
				employeeName,
				totalDays,
			)
		}
	}

	return nil
}
