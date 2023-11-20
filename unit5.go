package main

import "fmt"

func pyramid(n int) {
	for i := 1; i < n; i++ {
		for j := 1; j <= i-1; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= 2*i-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()

	}
}
func main() {
	n := 9
	pyramid(n)
}
