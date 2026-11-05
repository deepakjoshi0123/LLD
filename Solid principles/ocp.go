package main

import "fmt"

// ============================================================
// WITHOUT OCP
// ============================================================

type EmployeeWithoutOCP struct {
	Name   string
	Salary float64
}

func CalculateBonusWithoutOCP(employee EmployeeWithoutOCP, employeeType string) float64 {

	if employeeType == "Developer" {
		return employee.Salary * 0.10
	}

	if employeeType == "Manager" {
		return employee.Salary * 0.20
	}

	// for new employee again we will be doing if and modifying existing code ...
	return 0
}

// ============================================================
// WITH OCP
// ============================================================

type EmployeeWithOCP struct {
	Name   string
	Salary float64
}

type BonusCalculator interface {
	Calculate(employee EmployeeWithOCP) float64
}

type DeveloperBonus struct{}

func (d DeveloperBonus) Calculate(employee EmployeeWithOCP) float64 {
	return employee.Salary * 0.10
}

type ManagerBonus struct{}

func (m ManagerBonus) Calculate(employee EmployeeWithOCP) float64 {
	return employee.Salary * 0.20
}

type HRBonus struct{}

func (h HRBonus) Calculate(employee EmployeeWithOCP) float64 {
	return employee.Salary * 0.15
}

// ============================================================
// MAIN
// ============================================================

func main() {

	// --------------------------------------------------------
	// WITHOUT OCP
	// --------------------------------------------------------

	fmt.Println("========== WITHOUT OCP ==========")

	employee1 := EmployeeWithoutOCP{
		Name:   "DX",
		Salary: 20000,
	}

	bonus := CalculateBonusWithoutOCP(employee1, "Developer")
	fmt.Println("Developer Bonus:", bonus)

	bonus = CalculateBonusWithoutOCP(employee1, "Manager")
	fmt.Println("Manager Bonus:", bonus)

	// If we add HR, we must MODIFY CalculateBonusWithoutOCP()
	//
	// if employeeType == "HR" {
	//     return employee.Salary * 0.15
	// }

	fmt.Println()

	// --------------------------------------------------------
	// WITH OCP
	// --------------------------------------------------------

	fmt.Println("========== WITH OCP ==========")

	employee2 := EmployeeWithOCP{
		Name:   "DX",
		Salary: 20000,
	}

	var calculator BonusCalculator

	calculator = DeveloperBonus{}
	fmt.Println("Developer Bonus:", calculator.Calculate(employee2))

	calculator = ManagerBonus{}
	fmt.Println("Manager Bonus:", calculator.Calculate(employee2))

	calculator = HRBonus{}
	fmt.Println("HR Bonus:", calculator.Calculate(employee2))

	// If tomorrow we need Tester bonus,
	// we simply create a new TesterBonus type.
	//
	// We DON'T modify DeveloperBonus or ManagerBonus.
}
