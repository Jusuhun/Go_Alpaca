package fizzbuzz

import "fmt"

func Execute() {
	for count := 1; count <= 100; count++ {
		Fizz := (count%3 == 0)
		Buzz := (count%5 == 0)

		if Fizz && Buzz {
			fmt.Println(count, "FizzBuzz")
		} else if Fizz {
			fmt.Println(count, "Fizz")
		} else if Buzz {
			fmt.Println(count, "Buzz")
		}
	}
}
