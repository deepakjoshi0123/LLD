package main

import "fmt"

// ============================================================
// WITHOUT SRP
// ============================================================

// Employee is responsible for:
// 1. Employee data
// 2. Tax calculation
// 3. Saving to database
// 4. Printing details
//
// So Employee has multiple responsibilities.

type Employee struct {
	Name   string
	Salary float64
}

func (e Employee) PrintDetails() {
	fmt.Println("Name:", e.Name)
	fmt.Println("Salary:", e.Salary)
}

func (e Employee) CalculateTax() float64 {
	return e.Salary * 0.10
}

func (e Employee) SaveToDatabase() {
	fmt.Println("Saving employee to database...")
}

// ============================================================
// WITH SRP
// ============================================================

// Employee is now responsible only for employee data.

type EmployeeData struct {
	Name   string
	Salary float64
}

// TaxCalculator is responsible only for tax calculation.

type TaxCalculator struct{}

func (t TaxCalculator) CalculateTax(employee EmployeeData) float64 {
	return employee.Salary * 0.10
}

// EmployeeRepository is responsible only for database operations.

type EmployeeRepository struct{}

func (r EmployeeRepository) Save(employee EmployeeData) {
	fmt.Println("Saving employee to database...")
}

// EmployeePrinter is responsible only for displaying employee details.

type EmployeePrinter struct{}

func (p EmployeePrinter) Print(employee EmployeeData) {
	fmt.Println("Name:", employee.Name)
	fmt.Println("Salary:", employee.Salary)
}

// ============================================================
// MAIN
// ============================================================

func main() {

	// if the tax calculation logic changes, I need to modify Employee.
	// If the database logic changes, I again need to modify Employee
	fmt.Println("========== WITHOUT SRP ==========")

	employee := Employee{
		Name:   "DX",
		Salary: 20000,
	}

	employee.PrintDetails()

	tax := employee.CalculateTax()
	fmt.Println("Tax:", tax)

	employee.SaveToDatabase()

	fmt.Println()

	fmt.Println("========== WITH SRP ==========")

	employeeData := EmployeeData{
		Name:   "Rahul",
		Salary: 50000,
	}

	// Display
	printer := EmployeePrinter{}
	printer.Print(employeeData)

	// Tax calculation
	taxCalculator := TaxCalculator{}
	tax = taxCalculator.CalculateTax(employeeData)
	fmt.Println("Tax:", tax)

	// Database
	repository := EmployeeRepository{}
	repository.Save(employeeData)
}
