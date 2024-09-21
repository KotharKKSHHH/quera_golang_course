package main

import "fmt"

func main() {

	var income, pay int
	fmt.Scan(&income)

	switch {
	case income >= 100000:
		income -= 100000
		pay = 12000 + income/5
	case income >= 50000:
		income -= 50000
		pay = 4500 + income*15/100
	case income <= 10000:
		pay = income / 20
	case income > 10000:
		income -= 10000
		pay = 500 + income/10
	}

	fmt.Println(pay)
}
