package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	var (
		mid          int
		n            int
		teachersName string
		//score        int
		scores   = []int{}
		teachers = []string{}
		results  = []string{}
	)

	fmt.Scan(&n)
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		teachersName, _ = reader.ReadString('\n')
		teachersName = strings.TrimSpace(teachersName) // Remove the newline character
		teachers = append(teachers, teachersName)
		scoreStrs, _ := reader.ReadString('\n')
		scoreStrs = strings.TrimSpace(scoreStrs) // Remove the newline character
		for _, scoreStr := range strings.Split(scoreStrs, " ") {
			score, _ := strconv.Atoi(scoreStr)
			scores = append(scores, score)
		}
		for _, score := range scores {
			mid = score + mid
		}

		mid = mid / len(scores)
		//fmt.Println(mid)
		switch {
		case mid >= 80:
			results = append(results, "Excellent")
		case mid >= 60:
			results = append(results, "Very Good")
		case mid >= 40:
			results = append(results, "Good")
		default:
			results = append(results, "Fair")
		}
		scores = nil
		mid = 0
	}

	for i := 0; i < n; i++ {
		fmt.Println(teachers[i] + " " + results[i])

	}
}
