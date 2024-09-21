package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var count int
	fmt.Scanln(&count)
	var orders []string
	var books []string
	book := make(map[int]string)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < count; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			orders = append(orders, line)
		}
	}

	for i := 0; i < count; i++ {
		line := orders[i]
		words := strings.Split(line, " ")
		var name string
		for i := 2; i < len(words); i++ {
			name += words[i]
		}
		num, _ := strconv.Atoi(words[1])

		if words[0] == "ADD" {
			book[num] = name
		} else if words[0] == "REMOVE" {
			if _, ok := book[num]; ok {
				delete(book, num)
			}
		}
	}
	for _, v := range book {
		books = append(books, v)
	}

	if len(books) > 0 {
		sort.Strings(books)
		// fmt.Println(books)
		arr := make(map[string]int)
		// fmt.Println(arr)
		arr[books[0]] = 1
		//fine---------------------------------------
		for i := 1; i < len(books); i++ {
			if _, ok := arr[books[i]]; ok {
				arr[books[i]]++
			} else {
				arr[books[i]] = 1
			}
		}
		// fmt.Println(arr)

		var b2 []string
		for k, _ := range arr {
			b2 = append(b2, k)
		}
		for _, v := range b2 {
			for k, val := range book {
				if v == val && arr[v] == 1 {
					fmt.Println(k)
				} else if v == val && arr[v] > 1 {
					var a []int
					for key, value := range book {
						if value == v {
							a = append(a, key)
						}
					}
					sort.Ints(a)
					for _, val := range a {
						fmt.Println(val)

					}
					break
				}
			}
		}
	}
	// -------------------------------------------------
	// count = 101
	// for k, v := range arr {
	// 	if v > 1 {
	// 		for key, val := range book {
	// 			if val == k {
	// 				if key < count {
	// 					count = key
	// 				}
	// 			}
	// 		}
	// 		arr[k] = count
	// 	} else {
	// 		for key, val := range book {
	// 			if val == k {
	// 				arr[k] = key
	// 			}
	// 		}
	// 	}

	// }
	// fmt.Println(arr)
	// fine----------------------------------------

}
