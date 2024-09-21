package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := 0
	var names []string
	var orders []int
	var speed []int

	// fmt.Println("Enter the count of lines:")
	fmt.Scanln(&count)
	scanner := bufio.NewScanner(os.Stdin)

	// fmt.Println("Enter the lines of data:")
	for i := 0; i < count; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			// fmt.Printf("Processing line: %s\n", line)
			words := strings.Split(line, " ")
			names = append(names, words[0])
			// fmt.Printf("Name: %s\n", words[0])
			words = words[1:]
			for _, el := range words {
				if num, err := strconv.Atoi(el); err == nil {
					speed = append(speed, num)
					// fmt.Printf("Appending to speed: %d\n", num)
				} else {
					// fmt.Printf("Error converting %s to integer\n", el)
				}
			}
			speed = append(speed, -1)
			// fmt.Println("Speed after appending -1:", speed)
		}
	}

	speed = append(speed, -2)
	// fmt.Println("Final speed after appending -2:", speed)

	i := 0
	for i < len(speed) {
		// fmt.Printf("Current index i: %d, speed[i]: %d\n", i, speed[i])
		if speed[i] == -2 {
			// fmt.Println("Encountered -2, breaking the loop")
			break
		}
		if speed[i] == -1 {
			// fmt.Println("Encountered -1, processing rocket")
			rocket := speed[:i]
			// fmt.Printf("Rocket: %v\n", rocket)
			speed = speed[i+1:]
			// fmt.Printf("Updated speed: %v\n", speed)
			i = 0 // Reset i to 0 after modifying speed slice
			for j := 3; j <= len(rocket); j++ {
				for k := 0; k <= len(rocket)-j; k++ {
					// fmt.Printf("Adding rocket slice %v to orders\n", rocket[k:k+j])
					orders = append(orders, rocket[k:k+j]...)
					orders = append(orders, -1)
					// fmt.Printf("Updated orders: %v\n", orders)
				}
			}
			orders = append(orders, -2)
			// fmt.Printf("Orders after appending -2: %v\n", orders)
		} else {
			i++
		}
	}
	orders = append(orders, -3)
	speed = nil
	// fmt.Println("Final orders before processing:", orders)
	a := 0
	for a < len(orders) {
		// fmt.Printf("Current index a: %d, orders[a]: %d\n", a, orders[a])
		if orders[a] == -2 && orders[a+1] != -3 {
			// fmt.Println("Encountered -2, processing remaining orders")
			speed = append(speed, -2)
			orders = orders[a+1:]
			// fmt.Printf("Updated speed after -2: %v\n", speed)

		}

		if orders[a+1] == -3 {
			// fmt.Println("Encountered -3, breaking the loop")
			break
		} else if orders[a] == -1 {

			// fmt.Println("Encountered -1, checking subsequences")
			s := orders[:a]
			// fmt.Printf("Current subsequence s: %v\n", s)
			nu := true
			orders = orders[a+1:]

			sub := s[1] - s[0]
			for j := 1; j < len(s)-1; j++ {
				if s[j+1]-s[j] != sub {
					nu = false
					break
				}
			}
			if nu {
				// fmt.Println("Found arithmetic sequence, adding to speed :")
				// fmt.Println(s)

				speed = append(speed, s...)
				speed = append(speed, -1)
			}

			if orders[1] == -3 {
				break
			}

			a = 0
		} else {
			a++
		}
	}
	speed = append(speed, -2)
	i = 0
	orders = nil
	for _, el := range speed {
		if el == -1 {
			i++
		} else if el == -2 {
			orders = append(orders, i)
			i = 0
		}
	}
	// fmt.Println("Final speed:", speed)
	// fmt.Println("Final orders:", orders)

	for index, el := range orders {
		fmt.Println(names[index], el)
	}

}
