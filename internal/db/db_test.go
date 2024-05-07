package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
