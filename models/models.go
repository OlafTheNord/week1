package models

type Employee struct {
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	HireDate   string  `json:"hire_date"`
}
