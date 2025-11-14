package models

type Employee struct {
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
	HireDt     string  `json:"hire_dt"`
}

type Employees struct {
	Data []Employee `json:"data"`
}
