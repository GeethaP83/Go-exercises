package dbClient

import (
	"fmt"

	"home/geetha/Desktop/practice/grpc/dataModels"
)

func (d *DBClient) CreateEmployee(newEmployee dataModels.Employee) (dataModels.Employee, error) {
	err := d.Db.Create(&newEmployee).Error
	if err != nil {
		return dataModels.Employee{}, fmt.Errorf(err.Error())
	}
	return newEmployee, nil
}

func (d *DBClient) GetAllEmployees() ([]dataModels.Employee, error) {
	allEmployees := []dataModels.Employee{}
	err := d.Db.Find(&allEmployees).Error
	if err != nil {
		return []dataModels.Employee{}, fmt.Errorf(err.Error())
	}
	return allEmployees, nil
}

func (d *DBClient) UpdateEmailOfEmployee(name string, email string) (dataModels.Employee, error) {
	err := d.Db.Model(&dataModels.Employee{}).Where("name=?", name).Update("email", email).Error
	employee := dataModels.Employee{}
	d.Db.Find(&employee).Where("name=?", name)
	if err != nil {
		return employee, fmt.Errorf(err.Error())
	}
	return employee, nil
}

func (d *DBClient) DeleteEmployee(name string) (bool, error) {
	err := d.Db.Where(&dataModels.Employee{Name: name}).Delete(&dataModels.Employee{}).Error
	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	return true, nil
}
