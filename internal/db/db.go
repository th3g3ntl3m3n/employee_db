package db

import "github.com/oklog/ulid/v2"

type Database interface {
	GetAllEmployee(offset, limit int) []Employee
	GetEmployeeByID(employeeID string) Employee
	UpdateEmployee(employee Employee) Employee
	DeleteEmployee(employeeID string) bool
	CreateEmployee(employee Employee) Employee
}

const DefaultLimit = 5

type Employee struct {
	ID       string
	Name     string
	Salary   int32
	Position string
}

type DB struct {
	Employees []Employee
}

func NewDB() Database {
	return DB{Employees: make([]Employee, 0)}
}

func (db DB) GetAllEmployee(skip, limit int) []Employee {

	if skip > len(db.Employees) {
		skip = len(db.Employees)
	}

	upper := skip + limit
	if upper > len(db.Employees) {
		upper = len(db.Employees)
	}

	return db.Employees[skip:upper]
}

func (db DB) GetEmployeeByID(employeeID string) Employee {
	return Employee{}
}

func (db DB) UpdateEmployee(employee Employee) Employee {
	return Employee{}
}
func (db DB) DeleteEmployee(employeeID string) bool {
	return true
}
func (db DB) CreateEmployee(employee Employee) Employee {
	id := ulid.Make().String()
	employee.ID = id
	return employee
}
