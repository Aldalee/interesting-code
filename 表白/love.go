package main

import "fmt"

func main() {
	a := 0.0
	for y := 1.5; y > -1.5; y -= 0.09 {
		for x := -1.5; x < 1.5; x += 0.02 {
			a = x*x + y*y - 1.0
			if a*a*a < x*x*y*y*y {
				fmt.Print("u")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
