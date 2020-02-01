package main

import "fmt"

func main() {
	fmt.Println("This is a world class tutorial and the students are going to pass. Every one of you.")

	foo()

	fmt.Println("Hello the third paragraph is here we are testing")

	z := add(3, 7)
	fmt.Println(z)

	for i := 5; i < 20; i++ {
		fmt.Printf("The value %v\n", i)
	}
}

func foo() {
	fmt.Println("This is a second paragraph about Foo Fighters")
}

func add(x, y int) int {
	return (x + y)
}
