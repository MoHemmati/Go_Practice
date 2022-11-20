package main

import "fmt"

func main() {
	fmt.Println("Hello World!! " + foo())
	for i := 0; i < 101; i++ {
		if i%2 == 0 {
			fmt.Print(i)
		}
	}
	fmt.Println()
	fmt.Println(bar())
}
func bar() string {
	return "Exit!!"
}
func foo() string {
	return "This is foo!!"
}
