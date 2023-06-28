package main

import (
	"flag"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Employee struct {
	gorm.Model
	Name         string
	Email        string
	ManagerID    *uint
	Manager      *Employee
	DepartmentID uint
	Department   Department
	CompanyID    uint
}

type Department struct {
	gorm.Model
	Name      string
	CompanyID uint
	Employees []Employee
}

type Company struct {
	gorm.Model
	Name        string
	Employees   []Employee
	Departments []Department
}

func (e *Employee) AfterCreate() {
	fmt.Printf("Welcome mail sent to %s\n", e.Email)
}

func main() {
	host := flag.String("host", "localhost", "PostgreSQL host")
	port := flag.String("port", "5432", "PostgreSQL port")
	user := flag.String("user", "postgres", "PostgreSQL user")
	password := flag.String("password", "password", "PostgreSQL password")
	database := flag.String("database", "postgres", "PostgreSQL database name")
	flag.Parse()
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", *host, *port, *user, *password, *database)
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Debug().DropTableIfExists(&Company{}, &Department{}, &Employee{})
	db.Debug().CreateTable(&Employee{}, &Company{}, &Department{})
	fmt.Println("Connection Made")

	company := Company{
		Name: "ABC",
		Departments: []Department{
			{Name: "Dev"},
			{Name: "HR"},
		},
	}
	db.Create(&company)

	departmentHR := Department{}
	db.Where(&Department{Name: "HR"}).Find(&departmentHR)

	employee := Employee{
		Name:         "a",
		Email:        "a@abc",
		DepartmentID: departmentHR.ID,
		CompanyID:    company.ID,
	}
	db.Save(&employee)

	departmentDev := Department{Name: "Dev"}
	db.Where(&departmentDev).Find(&departmentDev)

	manager := Employee{
		Name:         "B",
		Email:        "bb@abc",
		DepartmentID: departmentDev.ID,
		CompanyID:    company.ID,
	}
	db.Save(&manager)

	employee2 := Employee{
		Name:         "b",
		Email:        "b@abc",
		Manager:      &manager,
		DepartmentID: departmentDev.ID,
		CompanyID:    company.ID,
	}
	db.Save(&employee2)

	Employees := []Employee{}
	db.Preload("Manager").Preload("Department").Find(&Employees)

	for _, v := range Employees {
		fmt.Println("Employee Name : " + v.Name)
		fmt.Println("Email : " + v.Email)
		fmt.Println("Department : " + v.Department.Name)
		if v.ManagerID != nil {
			fmt.Println("Manager : " + v.Manager.Name)
		} else {
			fmt.Println("Manager : none")
		}
		fmt.Println()
	}

}
