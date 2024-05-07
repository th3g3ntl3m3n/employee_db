package db

import "github.com/oklog/ulid/v2"

type Database interface {
	GetAllEmployee(offset, limit int) ([]Employee, error)
	GetEmployeeByID(employeeID string) (Employee, error)
	UpdateEmployee(employee Employee) (Employee, error)
	DeleteEmployee(employeeID string) error
	CreateEmployee(employee Employee) (Employee, error)
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

func (db DB) GetAllEmployee(skip, limit int) ([]Employee, error) {

	if skip > len(db.Employees) {
		skip = len(db.Employees)
	}

	upper := skip + limit
	if upper > len(db.Employees) {
		upper = len(db.Employees)
	}

	return db.Employees[skip:upper], nil
}

func (db DB) GetEmployeeByID(employeeID string) (Employee, error) {
	return Employee{}, nil
}

func (db DB) UpdateEmployee(employee Employee) (Employee, error) {
	return Employee{}, nil
}
func (db DB) DeleteEmployee(employeeID string) error {
	return nil
}
func (db DB) CreateEmployee(employee Employee) (Employee, error) {
	id := ulid.Make().String()
	employee.ID = id
	return employee, nil
}
