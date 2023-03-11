package models

type EmployeeModel struct {
	Name       string  `json:"name"`
	Age        int     `json:"age"`
	Department string  `json:"department"`
	Email      string  `json:"email"`
	BaseSalary float32 `json:"baseSalary"`
}
