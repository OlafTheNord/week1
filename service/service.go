package service

import (
	"encoding/json"
	"fmt"
	"main/models"
	"os"
)

func EmployeeReader(name string) (map[string]float64, error) {
	// зачитали файл
	file, err := os.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("cannot read file: %v", err)
	}

	// записали сотрудников в массив модельки
	var employees []models.Employee
	if err := json.Unmarshal(file, &employees); err != nil {
		return nil, fmt.Errorf("cannot unmarshal file: %v", err)
	}

	// валидируем что не просто так работают
	for _, employee := range employees {
		if employee.Salary < 0 {
			return nil, fmt.Errorf("%s employee salary is negative", employee.Name)
		}
		if employee.Department == "" {
			return nil, fmt.Errorf("%s employee department is empty", employee.Name)
		}
	}

	// записали в мапу все зарплаты
	salaryMid := make(map[string][]float64)
	for _, employee := range employees {
		salaryMid[employee.Department] = append(salaryMid[employee.Department], employee.Salary)
	}

	//перебираем все зарпалты для департаментов и считаем среднее
	salaryAvg := make(map[string]float64)
	for dept, salar := range salaryMid {
		sum := 0.0
		for _, s := range salar {
			sum += s
		}
		// пишем среднее
		salaryAvg[dept] = sum / float64(len(salar))
	}
	return salaryAvg, nil
}
