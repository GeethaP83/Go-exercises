package dbClient

import (
	"testing"
	_ "time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"home/geetha/Desktop/practice/grpc/dataModels"
)

func newMockDBClient() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, _ := sqlmock.New()
	gormDB, err := gorm.Open("postgres", db)
	if err != nil {
		panic("Unexpected error while creating mock sql connection.")
	}
	return gormDB, mock
}

func TestCreateEmployee(t *testing.T) {
	db, mock := newMockDBClient()
	defer db.Close()
	defer mock.ExpectClose()

	DBClient := DBClient{
		Db: db,
	}

	t.Run("Adding Employee", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "employees"(.+)`).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
		mock.ExpectCommit()

		newEmployee, err := DBClient.CreateEmployee(dataModels.Employee{
			Name:         "test",
			Email:        "test@test",
			DepartmentID: 1,
			CompanyID:    1,
		})
		assert.Nil(t, err)

		if newEmployee.ID != 1 {
			t.Errorf("Unable to add course %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("There were unfulfilled expectations: %s", err)
		}

	})
}
