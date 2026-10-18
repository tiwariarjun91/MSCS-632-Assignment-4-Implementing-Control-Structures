
# Assignment 4: Implementing Control Structures

**Course:** MSCS-632 Advanced Programming Languages  
**Application:** Employee Shift Scheduling System  
**Language:** Go

## Overview

This project implements an employee scheduling system using Go. The application generates a weekly work schedule based on employee shift preferences while following the company's staffing requirements.

The implementation demonstrates Go control structures such as `if/else`, `for` loops, `break`, and `continue`, along with structs, maps, and slices.

## Scheduling Rules

- The company operates seven days a week.
- Each day has three shifts: morning, afternoon, and evening.
- Every shift requires at least two employees.
- An employee can work only one shift per day.
- An employee can work a maximum of five days per week.
- Employee shift preferences are prioritized when possible.
- Scheduling conflicts are resolved by assigning employees to other available shifts.

The application requires at least nine employees to cover the 42 weekly shift assignments.

## Project Structure

```text
go/
├── main.go
├── models.go
├── input.go
├── scheduler.go
├── flow.go
├── validation.go
├── output.go
├── employees.json
└── go.mod
```

## Implementation

The application is divided into several components:

- **Input:** Reads employee information from `employees.json` or accepts interactive input.
- **Scheduling:** Uses a min-cost flow algorithm to assign employees while considering their preferences.
- **Conflict Resolution:** Reassigns employees when necessary to satisfy scheduling constraints.
- **Validation:** Checks daily assignments, weekly limits, and minimum staffing.
- **Output:** Prints the weekly schedule, employee work totals, and preference-matching statistics.

## Running the Application

### Requirements

- Go installed

### Run with Sample Data

Navigate to the Go directory:

```bash
cd go
go run .
```

### Run with Interactive Input

```bash
go run . --interactive
```

Enter at least nine employees and their preferred shifts for all seven days.

## Schedule Output

The application prints employee assignments for every day and shift, followed by weekly employee totals and preference statistics.

### Go Implementation Screenshot


## Validation

The generated schedule is checked to ensure:

1. Every shift has at least two employees.
2. No employee works more than one shift per day.
3. No employee works more than five days per week.
4. All seven days and three shifts are covered.

If the application cannot generate a valid schedule, it reports an error.

## Conclusion

The Go implementation demonstrates how control structures and data structures can be used to solve a scheduling problem. It separates input, scheduling, validation, and output into individual files to keep the application organized and easier to maintain.
