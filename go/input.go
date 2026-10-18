
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// loadEmployees reads employee information from a JSON file.
func loadEmployees(filePath string) ([]Employee, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("unable to read employee file: %w", err)
	}

	var employees []Employee

	if err := json.Unmarshal(fileContent, &employees); err != nil {
		return nil, fmt.Errorf("invalid employee JSON: %w", err)
	}

	if err := validEmployees(employees); err != nil {
		return nil, err
	}

	return employees, nil
}

// collectEmployees accepts employee details through the terminal.
func collectEmployees() ([]Employee, error) {
	reader := bufio.NewReader(os.Stdin)
	employees := make([]Employee, 0)

	fmt.Println("EMPLOYEE SHIFT REGISTRATION")
	fmt.Println("---------------------------")
	fmt.Println("Enter at least 9 employees.")
	fmt.Println("Leave the employee name blank when finished.")

	for {
		fmt.Print("\nEmployee name: ")

		nameInput, err := reader.ReadString('\n')
		if err != nil {
			return nil, fmt.Errorf("unable to read employee name: %w", err)
		}

		employeeName := strings.TrimSpace(nameInput)

		if employeeName == "" {
			break
		}

		duplicateName := false

		for _, existingEmployee := range employees {
			if existingEmployee.Name == employeeName {
				duplicateName = true
				break
			}
		}

		if duplicateName {
			fmt.Println("This employee has already been registered.")
			continue
		}

		newEmployee := Employee{
			Name:        employeeName,
			Preferences: make([]string, 0, len(days)),
		}

		for _, day := range days {
			for {
				fmt.Printf(
					"Preferred shift for %s (morning/afternoon/evening): ",
					day,
				)

				shiftInput, err := reader.ReadString('\n')
				if err != nil {
					return nil, fmt.Errorf("unable to read shift: %w", err)
				}

				preferredShift := strings.ToLower(
					strings.TrimSpace(shiftInput),
				)

				validShift := false

				for _, availableShift := range shifts {
					if preferredShift == availableShift {
						validShift = true
						break
					}
				}

				if !validShift {
					fmt.Println("Invalid shift. Please try again.")
					continue
				}

				newEmployee.Preferences = append(
					newEmployee.Preferences,
					preferredShift,
				)

				break
			}
		}

		employees = append(employees, newEmployee)

		fmt.Printf(
			"Added %s successfully. Total employees: %d\n",
			employeeName,
			len(employees),
		)
	}

	if err := validEmployees(employees); err != nil {
		return nil, err
	}

	return employees, nil
}
