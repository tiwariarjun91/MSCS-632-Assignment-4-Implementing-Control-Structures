package main

import (
	"errors"
	"math/rand"
	"sort"
)

// BuildSchedule assigns employees to shifts for the entire week.
// It prioritizes preferred shifts while respecting workload limits.
func BuildSchedule(employees []Employee, seed int64) (Schedule, error) {
	if err := validEmployees(employees); err != nil {
		return nil, err
	}

	employeeCount := len(employees)
	totalDays := len(days)
	shiftsPerDay := len(shifts)
	requiredPerShift := 2
	totalAssignments := totalDays * shiftsPerDay * requiredPerShift

	// Each section of the graph represents a scheduling constraint.
	source := 0
	employeeStart := 1
	dayStart := employeeStart + employeeCount
	shiftStart := dayStart + employeeCount*totalDays
	sink := shiftStart + totalDays*shiftsPerDay

	graph := make([][]Edge, sink+1)
	assignmentRefs := make([]EdgeRef, 0)
	randomizer := rand.New(rand.NewSource(seed))

	// Connect employees to the source.
	// Capacity five limits each employee to five working days.
	for employeeIndex, employee := range employees {
		employeeNode := employeeStart + employeeIndex
		addEdge(&graph, source, employeeNode, 5, 0)

		for dayIndex := 0; dayIndex < totalDays; dayIndex++ {
			employeeDayNode := dayStart + employeeIndex*totalDays + dayIndex

			// Only one shift is allowed per employee per day.
			addEdge(&graph, employeeNode, employeeDayNode, 1, 0)

			for shiftIndex, shiftName := range shifts {
				shiftNode := shiftStart + dayIndex*shiftsPerDay + shiftIndex

				// Preferred shifts have lower costs.
				// Random variation helps break ties between employees.
				assignmentCost := randomizer.Intn(10)

				if employee.Preferences[dayIndex] != shiftName {
					assignmentCost += 100
				}

				edgeIndex := addEdge(
					&graph,
					employeeDayNode,
					shiftNode,
					1,
					assignmentCost,
				)

				assignmentRefs = append(assignmentRefs, EdgeRef{
					i:   employeeIndex,
					d:   dayIndex,
					s:   shiftIndex,
					u:   employeeDayNode,
					idx: edgeIndex,
				})
			}
		}
	}

	// Every shift requires exactly two employees in this model.
	for dayIndex := 0; dayIndex < totalDays; dayIndex++ {
		for shiftIndex := 0; shiftIndex < shiftsPerDay; shiftIndex++ {
			shiftNode := shiftStart + dayIndex*shiftsPerDay + shiftIndex
			addEdge(&graph, shiftNode, sink, requiredPerShift, 0)
		}
	}

	// Find the lowest-cost available assignment path.
	// Reverse edges allow earlier choices to be rearranged.
	for assigned := 0; assigned < totalAssignments; assigned++ {
		const infinity = int(1 << 30)

		distances := make([]int, len(graph))
		previousNode := make([]int, len(graph))
		previousEdge := make([]int, len(graph))
		inQueue := make([]bool, len(graph))

		for node := range distances {
			distances[node] = infinity
			previousNode[node] = -1
		}

		distances[source] = 0
		queue := []int{source}
		inQueue[source] = true

		for len(queue) > 0 {
			currentNode := queue[0]
			queue = queue[1:]
			inQueue[currentNode] = false

			for edgeIndex, edge := range graph[currentNode] {
				if edge.cap <= 0 {
					continue
				}

				newDistance := distances[currentNode] + edge.cost

				if newDistance < distances[edge.to] {
					distances[edge.to] = newDistance
					previousNode[edge.to] = currentNode
					previousEdge[edge.to] = edgeIndex

					if !inQueue[edge.to] {
						queue = append(queue, edge.to)
						inQueue[edge.to] = true
					}
				}
			}
		}

		if previousNode[sink] == -1 {
			return nil, errors.New(
				"unable to create a schedule with the available employees",
			)
		}

		// Update capacities along the selected path.
		for node := sink; node != source; {
			parent := previousNode[node]
			edgeIndex := previousEdge[node]
			reverseIndex := graph[parent][edgeIndex].rev

			graph[parent][edgeIndex].cap--
			graph[node][reverseIndex].cap++

			node = parent
		}
	}

	// Initialize an empty weekly schedule.
	weeklySchedule := make(Schedule)

	for _, day := range days {
		weeklySchedule[day] = make(map[string][]string)

		for _, shift := range shifts {
			weeklySchedule[day][shift] = []string{}
		}
	}

	// Extract employee assignments from the completed flow graph.
	for _, reference := range assignmentRefs {
		if graph[reference.u][reference.idx].cap == 0 {
			day := days[reference.d]
			shift := shifts[reference.s]
			employeeName := employees[reference.i].Name

			weeklySchedule[day][shift] = append(
				weeklySchedule[day][shift],
				employeeName,
			)
		}
	}

	// Sort names to make the printed schedule easier to read.
	for _, day := range days {
		for _, shift := range shifts {
			sort.Strings(weeklySchedule[day][shift])
		}
	}

	// Perform a final check before returning the schedule.
	if err := Verify(weeklySchedule); err != nil {
		return nil, err
	}

	return weeklySchedule, nil
}
