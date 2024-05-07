package db

var MockData = []Employee{
	{ID: "abc", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abb", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abd", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abe", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abf", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abg", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abh", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abx", Name: "vikas", Salary: 100, Position: "abs"},
	{ID: "abz", Name: "vikas", Salary: 100, Position: "abs"},
}

func NewMockDB() Database {
	return DB{Employees: &MockData}
}
