package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	routeGuide "home/geetha/Desktop/practice/grpc/proto"
)

const (
	address = "localhost:50030"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Could not connect", err.Error())
	}
	defer conn.Close()
	employeeDBClient := routeGuide.NewEmployeeDBClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	employeeCreated, err := employeeDBClient.AddEmployee(ctx, &routeGuide.EmployeeToBeAdded{Name: "Checkabc", Email: "check@abc", DepartmentId: 2, CompanyId: 1})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("CREATE : ")
	fmt.Printf("id : %v, name : %v and email : %v\n", employeeCreated.GetId(), employeeCreated.GetName(), employeeCreated.GetEmail())

	allEmployees, err := employeeDBClient.GetEmployees(ctx, &routeGuide.Empty{})
	fmt.Println(" ")
	fmt.Println("READ : ")

	if err != nil {
		log.Printf("Error, could'nt get all employees")
	}
	for _, employee := range allEmployees.Employees {
		fmt.Println(employee.GetId(), employee.GetName(), employee.GetEmail())
	}
	fmt.Println(" ")
	fmt.Println("UPDATE : ")
	employeeUpdated, err := employeeDBClient.UpdateEmailOfEmployee(ctx, &routeGuide.Employee{Name: "Checkabc", Email: "update@abc"})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Updated %v's email to %v\n", employeeUpdated.GetName(), employeeUpdated.GetEmail())
	fmt.Println(" ")
	fmt.Println("DELETE : ")
	employeeDeleted, err := employeeDBClient.DeleteEmployee(ctx, &routeGuide.Employee{Name: "Checkabc"})
	if err != nil {
		fmt.Println(employeeDeleted)
		panic(err.Error())
	}
	fmt.Println("Deletion done.")
}
