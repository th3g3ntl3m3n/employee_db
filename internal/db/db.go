package db

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type Database interface {
	GetAllEmployee(offset, limit int) ([]Employee, error)
	GetEmployeeByID(employeeID string) (Employee, error)
	UpdateEmployee(employee Employee) (Employee, error)
	DeleteEmployee(employeeID string) error
	CreateEmployee(employee Employee) (Employee, error)
}

const DefaultLimit = 5

var (
	ErrNotFound = fmt.Errorf("employee not found")
)

type Employee struct {
	ID       string
	Name     string
	Salary   int32
	Position string
}

type DB struct {
	Employees *[]Employee
}

func NewDB() Database {
	return DB{Employees: &[]Employee{}}
}

func (db DB) GetAllEmployee(skip, limit int) ([]Employee, error) {

	if skip > len(*db.Employees) {
		skip = len(*db.Employees)
	}

	upper := skip + limit
	if upper > len(*db.Employees) {
		upper = len(*db.Employees)
	}

	return (*db.Employees)[skip:upper], nil
}

func (db DB) GetEmployeeByID(employeeID string) (Employee, error) {
	for _, emp := range *db.Employees {
		if emp.ID == employeeID {
			return emp, nil
		}
	}

	return Employee{}, ErrNotFound
}

func (db DB) UpdateEmployee(employee Employee) (Employee, error) {
	for i := range *db.Employees {
		emp := (*db.Employees)[i]
		if emp.ID == employee.ID {
			emp.Name = employee.Name
			emp.Position = employee.Position
			emp.Salary = employee.Salary
			(*db.Employees)[i] = emp

			return emp, nil
		}
	}

	return Employee{}, ErrNotFound
}
func (db DB) DeleteEmployee(employeeID string) error {
	for i := 0; i < len(*db.Employees); i++ {
		if (*db.Employees)[i].ID == employeeID {
			*db.Employees = append((*db.Employees)[:i], (*db.Employees)[:i]...)

			return nil
		}
	}

	return ErrNotFound
}
func (db DB) CreateEmployee(employee Employee) (Employee, error) {
	id := ulid.Make().String()
	employee.ID = id
	*db.Employees = append(*db.Employees, employee)

	return employee, nil
}
