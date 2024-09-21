package main

import "slices"

func AddElement(numbers *[]int, element int) {
	*numbers = append(*numbers, element)
}

func FindMin(numbers *[]int) int {
	a := *numbers
	n := 0
	if len(a) > 0 {
		n = a[0]
		for i := 0; i < len(*numbers); i++ {
			if a[i] < n {
				n = a[i]
			}
		}
	}

	return n
}

func ReverseSlice(numbers *[]int) {
	slice := *numbers
	slices.Reverse(slice)

}

func SwapElements(numbers *[]int, i, j int) {
	slice := *numbers
	if i < len(slice) && j < len(slice) && i > -1 && j > -1 {
		temp := slice[i]
		slice[i] = slice[j]
		slice[j] = temp
	}

}
