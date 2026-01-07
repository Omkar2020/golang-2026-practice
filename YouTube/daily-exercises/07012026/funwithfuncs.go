package main

import "fmt"

var y int = 42

func oddOrEven(x int) {
	if x%2 == 0 {
		fmt.Println("The Number ", x, "is Even")
	} else {
		fmt.Println("The Number ", x, "is Odd")
	}
	var y int = 15
	if y%2 == 0 {
		fmt.Println("The Number ", y, "is Even")
	} else {
		fmt.Println("The Number ", y, "is Odd")
	}
}

func main() {
	var numberToCheck int = 9
	oddOrEven(numberToCheck)

}
