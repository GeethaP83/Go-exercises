package main

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

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

// func TestGetEmployees(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := Mocks.NewMockDatabase(controller)
// 	mockServer := &employeeDBServer{database: mockDb}
// 	mockCtx := context.Background()

// }

// func TestUpdateEmailOfEmployee(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := Mocks.NewMockDatabase(controller)
// 	mockServer := &employeeDBServer{database: mockDb}
// 	mockCtx := context.Background()

// }

// func TestDeleteEmployee(t *testing.T) {
// 	controller := gomock.NewController(t)
// 	defer controller.Finish()
// 	mockDb := Mocks.NewMockDatabase(controller)
// 	mockServer := &employeeDBServer{database: mockDb}
// 	mockCtx := context.Background()

// }
