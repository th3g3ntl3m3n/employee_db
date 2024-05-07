package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewDB(t *testing.T) {
	db := NewDB()
	assert.NotNil(t, db)
}

func TestGetAllEmployees(t *testing.T) {
	tests := []struct {
		name   string
		offset int
		limit  int
		expect int
	}{
		{
			name:   "db don't have enough data",
			offset: 0,
			limit:  25,
			expect: len(MockData),
		},
		{
			name:   "data all Okay",
			offset: 0,
			limit:  2,
			expect: 2,
		},
		{
			name:   "get data from inbetween",
			offset: 2,
			limit:  2,
			expect: 2,
		},
		{
			name:   "get limit 5 page 1",
			offset: 0,
			limit:  5,
			expect: 5,
		},
		{
			name:   "get limit 5 page 2",
			offset: 5,
			limit:  5,
			expect: 4,
		},
		{
			name:   "get limit 5 page 3",
			offset: 10,
			limit:  5,
			expect: 0,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(tx *testing.T) {
			tx.Parallel()
			db := NewMockDB()
			got, err := db.GetAllEmployee(test.offset, test.limit)
			assert.NoError(t, err)
			assert.Equal(tx, len(got), test.expect)
		})
	}
}

func TestCreateEmployees(t *testing.T) {
	tests := []struct {
		name string
		data Employee
	}{
		{
			name: "add new employee",
			data: Employee{Name: "Vikas", Salary: 1000, Position: "Dev"},
		},
		{
			name: "add new employee 2",
			data: Employee{Name: "Vikas 2", Salary: 1000, Position: "Dev"},
		},
		{
			name: "add new employee 3",
			data: Employee{Name: "Vikas 3", Salary: 1000, Position: "Dev"},
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(tx *testing.T) {
			tx.Parallel()
			db := NewMockDB()
			got, err := db.CreateEmployee(test.data)
			assert.NoError(t, err)
			assert.NotEmpty(t, got.ID)
		})
	}
}

func TestGetEmployeeByID(t *testing.T) {
	tests := []struct {
		name string
		data string
		err  error
	}{
		{
			name: "get employee with id abx",
			data: "abx",
			err:  nil,
		},
		{
			name: "get employee with id abz",
			data: "abz",
			err:  nil,
		},
		{
			name: "get employee with id ab1",
			data: "ab1",
			err:  ErrNotFound,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(tx *testing.T) {
			tx.Parallel()
			db := NewMockDB()
			got, err := db.GetEmployeeByID(test.data)
			assert.Equal(t, err, test.err)
			if test.err == nil {
				assert.Equal(t, got.ID, test.data)
			}
		})
	}
}

func TestUpdateEmployee(t *testing.T) {
	tests := []struct {
		name string
		data Employee
		err  error
	}{
		{
			name: "update employee with id abx",
			data: Employee{ID: "abx", Name: "NEWNAME", Salary: 10000, Position: "Something"},
			err:  nil,
		},
		{
			name: "update employee with id non-existent",
			data: Employee{ID: "ab1", Name: "Sim", Salary: 100, Position: "Something"},
			err:  ErrNotFound,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(tx *testing.T) {
			tx.Parallel()
			db := NewMockDB()
			_, err := db.UpdateEmployee(test.data)
			assert.Equal(t, err, test.err)
			if test.err == nil {
				got, err := db.GetEmployeeByID(test.data.ID)
				assert.Nil(t, err)
				assert.Equal(t, got.ID, test.data.ID)
				assert.Equal(t, got.Name, test.data.Name)
				assert.Equal(t, got.Salary, test.data.Salary)
				assert.Equal(t, got.Position, test.data.Position)
			}
		})
	}
}

func TestDeleteEmployee(t *testing.T) {
	tests := []struct {
		name string
		data string
		err  error
	}{
		{
			name: "delete employee with id abx",
			data: "abx",
			err:  nil,
		},
		{
			name: "delete employee with id non-existent",
			data: "ab1",
			err:  ErrNotFound,
		},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(tx *testing.T) {
			tx.Parallel()
			db := NewMockDB()
			err := db.DeleteEmployee(test.data)
			assert.Equal(t, test.err, err)
			if test.err == nil {
				_, err = db.GetEmployeeByID(test.data)
				assert.Equal(t, ErrNotFound, err)
			}
		})
	}
}

/*

curl -X POST http://localhost:8080/employees -H 'Content-Type: application/json' -d '{"Name": "Test", "Position": "Dev", "Salary": 10000}'
curl -X GET http://localhost:8080/employees

*/
