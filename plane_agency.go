package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var count int
	fmt.Scanln(&count)
	country := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)
	var nums []string
	for i := 0; i < count; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			words := strings.Split(line, " ")
			words[1] = words[1][1:]
			country[words[1]] = words[0]
			// fmt.Println(words)
		}
	}
	// fmt.Println(country)
	fmt.Scanln(&count)
	for i := 0; i < count; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			line = line[1:3]
			nums = append(nums, line)
		}
	}
	// fmt.Println(nums)
	for _, v := range nums {
		if _, ok := country[v]; ok {
			fmt.Println(country[v])
		} else {
			fmt.Println("Invalid Number")
		}
	}
}
