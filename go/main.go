
package main

import (
	"fmt"
	"os"
)

// main coordinates the employee scheduling process.
func main() {
	var employees []Employee
	var err error

	fmt.Println("Employee Shift Scheduling System")
	fmt.Println("--------------------------------")

	// Use interactive mode when requested.
	// Otherwise, load the sample employee data from JSON.
	if len(os.Args) > 1 && os.Args[1] == "--interactive" {
		employees, err = collectEmployees()
	} else {
		employees, err = loadEmployees("employees.json")
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Input error:", err)
		os.Exit(1)
	}

	fmt.Printf(
		"\nCreating weekly schedule for %d employees...\n",
		len(employees),
	)

	// Use a fixed seed so sample schedules are reproducible.
	const randomSeed int64 = 42

	weeklySchedule, err := BuildSchedule(employees, randomSeed)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Scheduling error:", err)
		os.Exit(1)
	}

	// Verify the generated schedule before displaying it.
	if err := Verify(weeklySchedule); err != nil {
		fmt.Fprintln(os.Stderr, "Validation error:", err)
		os.Exit(1)
	}

	printSchedule(weeklySchedule, employees)
}
