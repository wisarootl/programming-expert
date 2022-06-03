package main

import "fmt"

func FilterStrings(strings []string, filter func(string) bool) []string {
	filteredStrings := []string{}

	for _, str := range strings {
		if filter(str) {
			filteredStrings = append(filteredStrings, str)
		}
	}

	return filteredStrings
}

type Employee struct {
	salary float64
	bonus  float64
}

func GetEmployeeWages(employees []Employee) (salaries float64, bonuses float64) {
	for _, employee := range employees {
		salaries += employee.salary
		bonuses += employee.bonus
	}

	return
}

func main() {
	//  === Functions (cont.) ===
	fmt.Println("\n=== Functions ===")
	fmt.Println("\n=Q2")
	filterFunc := func (str string) bool {
  	return len(str) < 5
	}
	strings := []string{"hello", "hi", "tim", "ball", "fruit", "cat", "dog", "fish"}
	filteredStrings := FilterStrings(strings, filterFunc) // returns {"hi", "tim", "ball", "cat", "dog", "fish"}
	fmt.Println(filteredStrings)


	fmt.Println("\n=Q3")
	employees := []Employee{Employee{salary: 50000, bonus:10000}, Employee{salary: 45000, bonus:5000}, Employee{salary: 70000, bonus:20000}}
	salaries, bonuses := GetEmployeeWages(employees) // returns 165000 and 35000
	fmt.Println("Salaries : ", salaries)
	fmt.Println("bonuses : ", bonuses)
}