package main

import "fmt"

func main() {

	var p, q int
	fmt.Scan(&p, &q)

	for i := 1; ; i++ {
		switch {
		case i > q:
			break
		case i%p == 0:
			for j := 0; j < i/p; j++ {
				fmt.Printf("Hope ")
			}
			fmt.Println("")
		default:
			fmt.Println(i)
		}
		if i > q {
			break
		}
	}

}
