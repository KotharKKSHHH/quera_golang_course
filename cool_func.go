package main

import (
	"strconv"
	"strings"
)

type FilterFunc func(int) bool
type MapperFunc func(int) int

func IsSquare(x int) bool {
	if x == 1 || x == 0 {
		return true
	}
	if x < 0 {
		x = -x
	}
	for i := 0; i < x; i++ {
		if i*i == x {
			return true
		}
	}

	return false
}

func IsPalindrome(x int) bool {
	x2 := strconv.Itoa(x)
	slice := strings.Split(x2, "")
	if len(slice)%2 == 0 {
		for i := 0; i < len(slice)/2; i++ {
			if slice[i] != slice[len(slice)-i-1] {
				return false
			}
		}
	} else {
		for i := 0; i < len(slice)/2; i++ {
			if i+1 == len(slice)/2+1 {
				continue
			}
			for i := 0; i < len(slice)/2; i++ {
				if slice[i] != slice[len(slice)-i-1] {
					return false
				}
			}
		}
	}
	return false
}

func Abs(num int) int {
	if num < 0 {
		return -num
	} else {
		return num
	}

}

func Cube(num int) int {

	return num * num * num
}

func Filter(input []int, f FilterFunc) []int {
	var output []int
	for i := 0; i < len(input); i++ {
		if f(input[i]) {
			output = append(output, input[i])
		}
	}
	return output
}

func Map(input []int, m MapperFunc) []int {
	var output []int
	for i := 0; i < len(input); i++ {
		output = append(output, m(input[i]))
	}
	return output
}
