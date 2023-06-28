package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"home/geetha/Desktop/practice/grpc/dataModels"
	"home/geetha/Desktop/practice/grpc/dbClient"
	routeGuide "home/geetha/Desktop/practice/grpc/proto"
)

type employeeDBServer struct {
	routeGuide.UnimplementedEmployeeDBServer
	database dbClient.Database
	db       *gorm.DB
}

const (
	port = ":50030"
)

func (server *employeeDBServer) AddEmployee(ctx context.Context, input *routeGuide.EmployeeToBeAdded) (*routeGuide.Employee, error) {
	log.Printf("AddEmployee %v %v %v %v\n", input.GetName(), input.GetEmail(), input.GetDepartmentId(), input.GetCompanyId())
	newEmployee := dataModels.Employee{
		Name:         input.GetName(),
		Email:        input.GetEmail(),
		DepartmentID: uint(input.GetDepartmentId()),
		CompanyID:    uint(input.GetCompanyId()),
	}
	if newEmployee.Name == "" || newEmployee.Email == "" || newEmployee.DepartmentID == 0 || newEmployee.CompanyID == 0 {
		return nil, fmt.Errorf("Provide valid details!")
	}
	employee, err := server.database.CreateEmployee(newEmployee)
	result := routeGuide.Employee{
		Id:    int64(employee.ID),
		Name:  employee.Name,
		Email: employee.Email,
	}
	return &result, err
}

func (server *employeeDBServer) GetEmployees(ctx context.Context, input *routeGuide.Empty) (*routeGuide.AllEmployees, error) {
	log.Println("Getting All employees.")
	result := []*routeGuide.Employee{}
	allEmployees, err := server.database.GetAllEmployees()
	for _, employee := range allEmployees {
		result = append(result, &routeGuide.Employee{Name: employee.Name, Email: employee.Email, DepartmentId: int64(employee.DepartmentID), CompanyId: int64(employee.CompanyID), Id: int64(employee.ID)})
	}
	return &routeGuide.AllEmployees{Employees: result}, err
}

func (server *employeeDBServer) UpdateEmailOfEmployee(ctx context.Context, input *routeGuide.Employee) (*routeGuide.Employee, error) {
	log.Println("Updating employee.")
	if input.GetName() == "" || input.GetEmail() == "" {
		return nil, fmt.Errorf("Provide valid details!")
	}
	employee, err := server.database.UpdateEmailOfEmployee(input.GetName(), input.GetEmail())
	result := routeGuide.Employee{
		Name:  employee.Name,
		Email: employee.Email,
	}
	return &result, err
}

func (server *employeeDBServer) DeleteEmployee(ctx context.Context, input *routeGuide.Employee) (*routeGuide.Empty, error) {
	log.Println("Deleting employee.")
	if input.GetName() == "" {
		return nil, fmt.Errorf("Provide valid details!")
	}
	_, err := server.database.DeleteEmployee(input.GetName())
	return &routeGuide.Empty{}, err
}

func main() {
	dataModels.PopulateDB()
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}
	connectionString := fmt.Sprintf("user=postgres password=password dbname=postgres sslmode=disable")
	dbConnection, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer dbConnection.Close()
	grpcServer := grpc.NewServer()
	routeGuide.RegisterEmployeeDBServer(grpcServer, &employeeDBServer{
		database: &dbClient.DBClient{Db: dbConnection},
		db:       dbConnection,
	})
	log.Printf("Server listening at %v\n", listen.Addr())
	err = grpcServer.Serve(listen)
	if err != nil {
		log.Fatal("Not able to serve", err.Error())
	}
}
