package main

import (
	"fmt"
	"log"
	"main/service"
)

func main() {
	avg, err := service.EmployeeReader("employees.json")
	if err != nil {
		log.Fatalf("error in handling %s", err)
	}
	fmt.Println("Avg of salaries of dept: ")
	for k, v := range avg {
		fmt.Printf("%s: %.2f\n", k, v)
	}
}
