package db

import (
	"fmt"
	"sync"

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

var mu sync.RWMutex

type DB struct {
	Employees *[]Employee
}

func NewDB() Database {
	return DB{Employees: &[]Employee{}}
}

func (db DB) GetAllEmployee(skip, limit int) ([]Employee, error) {
	l := len(*db.Employees)
	if skip > l {
		skip = l
	}

	upper := skip + limit
	if upper > l {
		upper = l
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
	mu.Lock()
	defer mu.Unlock()

	for i, emp := range *db.Employees {
		if emp.ID == employee.ID {
			(*db.Employees)[i] = employee

			return (*db.Employees)[i], nil
		}
	}

	return Employee{}, ErrNotFound
}
func (db DB) DeleteEmployee(employeeID string) error {
	mu.Lock()
	defer mu.Unlock()

	for i, emp := range *db.Employees {
		if emp.ID == employeeID {
			*db.Employees = append((*db.Employees)[:i], (*db.Employees)[i+1:]...)

			return nil
		}
	}

	return ErrNotFound
}
func (db DB) CreateEmployee(employee Employee) (Employee, error) {
	mu.Lock()
	defer mu.Unlock()

	id := ulid.Make().String()
	employee.ID = id
	*db.Employees = append(*db.Employees, employee)

	return employee, nil
}
