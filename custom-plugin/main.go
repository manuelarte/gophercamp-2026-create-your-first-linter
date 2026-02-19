package main

import "fmt"

const adultAge = 18

func main() {
	fmt.Println("Insert your age:")
	var i int
	_, err := fmt.Scanf("%d", &i)
	if err != nil {
		panic(err)
	}

	if i >= adultAge {
		fmt.Println("You are allowed")
	} else {
		fmt.Println("You aren't allowed")
	}
}
