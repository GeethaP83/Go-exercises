package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jinzhu/gorm"

	"home/geetha/Desktop/practice/grpc/Mocks"
	"home/geetha/Desktop/practice/grpc/dataModels"
	p "home/geetha/Desktop/practice/grpc/proto"
)

func TestAddEmployee(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	mockServer := &employeeDBServer{database: mockDb}
	mockCtx := context.Background()

	tests := []struct {
		name           string
		input          *p.EmployeeToBeAdded
		mockFunc       func()
		expectedOutput *p.Employee
		expectedError  error
	}{
		{
			name: "Valid input",
			input: &p.EmployeeToBeAdded{
				Name:         "ABC",
				Email:        "abc@abc",
				DepartmentId: 1,
				CompanyId:    1,
			},
			mockFunc: func() {
				mockDb.EXPECT().CreateEmployee(gomock.Any()).Return(dataModels.Employee{
					Model: gorm.Model{ID: 1},
					Name:  "ABC",
					Email: "abc@abc",
				}, nil)
			},
			expectedOutput: &p.Employee{
				Id:    1,
				Name:  "ABC",
				Email: "abc@abc",
			},
			expectedError: nil,
		},
		{
			name: "Missing name",
			input: &p.EmployeeToBeAdded{
				Email:        "abc@abc",
				DepartmentId: 1,
				CompanyId:    1,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
		{
			name: "Missing email",
			input: &p.EmployeeToBeAdded{
				Name:         "ABC",
				DepartmentId: 1,
				CompanyId:    1,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
		{
			name: "Missing departmentId",
			input: &p.EmployeeToBeAdded{
				Name:      "ABC",
				Email:     "abc@abc",
				CompanyId: 1,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
		{
			name: "Missing companyId",
			input: &p.EmployeeToBeAdded{
				Name:         "ABC",
				Email:        "abc@abc",
				DepartmentId: 1,
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			output, err := mockServer.AddEmployee(mockCtx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("wrong error, Got: %v, Expected: %v", err, test.expectedError)
			}

		})
	}
}

func TestGetEmployees(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	mockServer := &employeeDBServer{database: mockDb}
	mockCtx := context.Background()

	input := &p.Empty{}
	employee := dataModels.Employee{
		Model:        gorm.Model{ID: 1},
		Name:         "ABC",
		Email:        "abc@abc",
		DepartmentID: 1,
		CompanyID:    1,
	}
	employeeList := []dataModels.Employee{
		employee,
	}
	mockDb.EXPECT().GetAllEmployees().Return(employeeList, nil)
	expectedOutput := &p.AllEmployees{}
	e := &p.Employee{
		Id:           1,
		Name:         "ABC",
		Email:        "abc@abc",
		DepartmentId: 1,
		CompanyId:    1,
	}
	expectedOutput.Employees = append(expectedOutput.Employees, e)

	output, err := mockServer.GetEmployees(mockCtx, input)
	if err != nil {
		t.Errorf("GetAllCourses returned an error: %v", err)
	}
	if !reflect.DeepEqual(output, expectedOutput) {
		t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", expectedOutput, output)
	}
}

func TestUpdateEmailOfEmployee(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	mockServer := &employeeDBServer{database: mockDb}
	mockCtx := context.Background()

	tests := []struct {
		name           string
		input          *p.Employee
		mockFunc       func()
		expectedOutput *p.Employee
		expectedError  error
	}{
		{
			name: "Valid input",
			input: &p.Employee{
				Name:  "ABC",
				Email: "abc@abc",
			},
			mockFunc: func() {
				mockDb.EXPECT().UpdateEmailOfEmployee(gomock.Any(), gomock.Any()).Return(
					dataModels.Employee{
						Name:  "ABC",
						Email: "abc@abc",
					}, nil)
			},
			expectedOutput: &p.Employee{
				Name:  "ABC",
				Email: "abc@abc",
			},
			expectedError: nil,
		},
		{
			name: "Empty Email",
			input: &p.Employee{
				Name:  "ABC",
				Email: "",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
		{
			name: "Empty Name",
			input: &p.Employee{
				Name:  "",
				Email: "abc@abc",
			},
			mockFunc:       func() {},
			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			output, err := mockServer.UpdateEmailOfEmployee(mockCtx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}

}

func TestDeleteEmployee(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	mockDb := Mocks.NewMockDatabase(controller)
	mockServer := &employeeDBServer{database: mockDb}
	mockCtx := context.Background()
	tests := []struct {
		name           string
		input          *p.Employee
		mockFunc       func()
		expectedOutput *p.Empty
		expectedError  error
	}{
		{
			name: "Valid input",
			input: &p.Employee{
				Name: "ABC",
			},
			mockFunc: func() {
				mockDb.EXPECT().DeleteEmployee(gomock.Any()).Return(true, nil)
			},
			expectedOutput: &p.Empty{},
			expectedError:  nil,
		},
		{
			name: "Empty Input",
			input: &p.Employee{
				Name: "",
			},
			mockFunc: func() {},

			expectedOutput: nil,
			expectedError:  fmt.Errorf("Provide valid details!"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mockFunc()
			output, err := mockServer.DeleteEmployee(mockCtx, test.input)

			// check expected output
			if !reflect.DeepEqual(output, test.expectedOutput) {
				t.Errorf("Got wrong results:( ,Expected: '%v',Got: '%v'", test.expectedOutput, output)
			}

			// check expected error
			if !reflect.DeepEqual(err, test.expectedError) {
				t.Errorf("Got wrong error: %v, expected: %v", err, test.expectedError)
			}

		})
	}

}
