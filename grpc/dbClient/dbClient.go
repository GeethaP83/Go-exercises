package dbClient

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"home/geetha/Desktop/practice/grpc/dataModels"
)

type Database interface {
	CreateEmployee(dataModels.Employee) (dataModels.Employee, error)
	GetAllEmployees() ([]dataModels.Employee, error)
	UpdateEmailOfEmployee(string, string) (dataModels.Employee, error)
	DeleteEmployee(string) (bool, error)
}

type DBClient struct {
	Db *gorm.DB
}
