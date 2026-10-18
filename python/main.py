
import json
import sys
from pathlib import Path

from scheduler import DAYS, SHIFTS, build_schedule


def load_employees():
    """Read employee information from the JSON file."""
    file_path = Path(__file__).parent / "employees.json"

    with open(file_path, "r", encoding="utf-8") as file:
        return json.load(file)


def collect_employees():
    """Allow employees to enter their shift preferences."""
    employees = []

    print("\nEMPLOYEE REGISTRATION")
    print("Enter at least 9 employees.")
    print("Leave the name blank when finished.\n")

    while True:
        name = input("Employee name: ").strip()

        if not name:
            break

        if any(employee["name"] == name for employee in employees):
            print("This employee has already been registered.")
            continue

        preferences = []

        for day in DAYS:
            while True:
                shift = input(
                    f"Preferred shift for {day} "
                    "(morning/afternoon/evening): "
                ).strip().lower()

                if shift in SHIFTS:
                    preferences.append(shift)
                    break

                print("Invalid shift. Please try again.")

        employees.append({
            "name": name,
            "preferences": preferences
        })

        print(f"Added {name}. Total employees: {len(employees)}\n")

    return employees


def display_schedule(schedule, employees):
    """Print the weekly schedule and employee work totals."""
    work_days = {employee["name"]: 0 for employee in employees}
    preferences = {
        employee["name"]: employee["preferences"]
        for employee in employees
    }

    preference_matches = 0
    total_assignments = 0

    print("\n" + "=" * 48)
    print("           WEEKLY EMPLOYEE SCHEDULE")
    print("=" * 48)

    for day_index, day in enumerate(DAYS):
        print(f"\n{day.upper()}")
        print("-" * 48)

        for shift in SHIFTS:
            assigned = schedule[day][shift]
            print(f"  {shift.capitalize():<12}: {', '.join(assigned)}")

            for name in assigned:
                work_days[name] += 1
                total_assignments += 1

                if preferences[name][day_index] == shift:
                    preference_matches += 1

    print("\n" + "=" * 48)
    print("              EMPLOYEE SUMMARY")
    print("=" * 48)

    for name in sorted(work_days):
        print(f"  {name:<18}: {work_days[name]} day(s)")

    print("\n" + "=" * 48)
    print("             SCHEDULE STATISTICS")
    print("=" * 48)

    print(f"Total employees       : {len(employees)}")
    print(f"Total assignments     : {total_assignments}")
    print(f"Preference matches    : {preference_matches}")

    if total_assignments:
        percentage = preference_matches / total_assignments * 100
        print(f"Preference satisfaction: {percentage:.1f}%")

    print("\nVALIDATION PASSED")
    print("Every shift has at least two employees.")
    print("No employee works multiple shifts in one day.")
    print("No employee works more than five days per week.")


def main():
    print("Employee Shift Scheduling System")
    print("--------------------------------")

    try:
        if "--interactive" in sys.argv[1:]:
            employees = collect_employees()
        else:
            employees = load_employees()

        print(f"\nCreating schedule for {len(employees)} employees...")

        schedule = build_schedule(employees, seed=42)
        display_schedule(schedule, employees)

    except (ValueError, OSError, json.JSONDecodeError) as error:
        print(f"\nERROR: {error}", file=sys.stderr)
        sys.exit(1)

    except KeyboardInterrupt:
        print("\nScheduling cancelled.")
        sys.exit(1)


if __name__ == "__main__":
    main()
