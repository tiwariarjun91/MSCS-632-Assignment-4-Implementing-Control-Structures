
import random

DAYS = [
    "Monday", "Tuesday", "Wednesday", "Thursday",
    "Friday", "Saturday", "Sunday"
]

SHIFTS = ["morning", "afternoon", "evening"]

MAX_WORK_DAYS = 5
MIN_EMPLOYEES = 2


def validate_employees(employees):
    """Check employee names and weekly shift preferences."""
    if len(employees) < 9:
        raise ValueError("At least 9 employees are required.")

    names = set()

    for employee in employees:
        name = employee.get("name", "").strip()
        preferences = employee.get("preferences", [])

        if not name or name in names:
            raise ValueError("Employee names must be unique and nonempty.")

        names.add(name)

        if len(preferences) != len(DAYS):
            raise ValueError(f"{name} must provide seven preferences.")

        for shift in preferences:
            if shift not in SHIFTS:
                raise ValueError(f"Invalid shift preference for {name}.")


def build_schedule(employees, seed=42):
    """Create a weekly schedule while respecting staffing constraints."""
    validate_employees(employees)

    rng = random.Random(seed)
    employee_count = len(employees)
    total_slots = len(DAYS) * len(SHIFTS) * MIN_EMPLOYEES

    # Each employee can work five days, with one shift per day.
    source = 0
    employee_start = 1
    day_start = employee_start + employee_count
    shift_start = day_start + employee_count * len(DAYS)
    sink = shift_start + len(DAYS) * len(SHIFTS)

    graph = [[] for _ in range(sink + 1)]
    assignment_edges = []

    def add_edge(start, end, capacity, cost):
        forward = [end, len(graph[end]), capacity, cost]
        backward = [start, len(graph[start]), 0, -cost]

        graph[start].append(forward)
        graph[end].append(backward)

        return len(graph[start]) - 1

    for index, employee in enumerate(employees):
        employee_node = employee_start + index
        add_edge(source, employee_node, MAX_WORK_DAYS, 0)

        for day_index in range(len(DAYS)):
            employee_day = day_start + index * len(DAYS) + day_index
            add_edge(employee_node, employee_day, 1, 0)

            for shift_index, shift in enumerate(SHIFTS):
                shift_node = shift_start + day_index * len(SHIFTS) + shift_index

                cost = rng.randrange(10)
                if employee["preferences"][day_index] != shift:
                    cost += 100

                edge_index = add_edge(employee_day, shift_node, 1, cost)
                assignment_edges.append(
                    (index, day_index, shift_index, employee_day, edge_index)
                )

    for day_index in range(len(DAYS)):
        for shift_index in range(len(SHIFTS)):
            shift_node = shift_start + day_index * len(SHIFTS) + shift_index
            add_edge(shift_node, sink, MIN_EMPLOYEES, 0)

    # Find an available assignment path for each required position.
    for _ in range(total_slots):
        infinity = float("inf")
        distances = [infinity] * len(graph)
        previous = [None] * len(graph)
        in_queue = [False] * len(graph)

        distances[source] = 0
        queue = [source]
        in_queue[source] = True

        while queue:
            current = queue.pop(0)
            in_queue[current] = False

            for edge_index, edge in enumerate(graph[current]):
                neighbor, _, capacity, cost = edge

                if capacity <= 0:
                    continue

                new_cost = distances[current] + cost

                if new_cost < distances[neighbor]:
                    distances[neighbor] = new_cost
                    previous[neighbor] = (current, edge_index)

                    if not in_queue[neighbor]:
                        queue.append(neighbor)
                        in_queue[neighbor] = True

        if previous[sink] is None:
            raise ValueError("Unable to create a valid weekly schedule.")

        # Update the graph after choosing an assignment path.
        node = sink

        while node != source:
            parent, edge_index = previous[node]
            edge = graph[parent][edge_index]
            reverse_index = edge[1]

            edge[2] -= 1
            graph[node][reverse_index][2] += 1
            node = parent

    schedule = {
        day: {shift: [] for shift in SHIFTS}
        for day in DAYS
    }

    for employee_index, day_index, shift_index, node, edge_index in assignment_edges:
        if graph[node][edge_index][2] == 0:
            day = DAYS[day_index]
            shift = SHIFTS[shift_index]
            schedule[day][shift].append(employees[employee_index]["name"])

    for day in DAYS:
        for shift in SHIFTS:
            schedule[day][shift].sort()

    verify_schedule(schedule)
    return schedule


def verify_schedule(schedule):
    """Confirm that the generated schedule follows company rules."""
    weekly_count = {}

    for day in DAYS:
        assigned_today = set()

        if day not in schedule:
            raise ValueError(f"Missing schedule for {day}.")

        for shift in SHIFTS:
            employees = schedule[day].get(shift, [])

            if len(employees) < MIN_EMPLOYEES:
                raise ValueError(f"Understaffed shift: {day} {shift}")

            for name in employees:
                if name in assigned_today:
                    raise ValueError(f"{name} is double-booked on {day}.")

                assigned_today.add(name)
                weekly_count[name] = weekly_count.get(name, 0) + 1

    for name, count in weekly_count.items():
        if count > MAX_WORK_DAYS:
            raise ValueError(f"{name} exceeds the five-day weekly limit.")

    return True
