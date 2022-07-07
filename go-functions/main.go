package main

import (
	"errors"
	"fmt"
	"log"
)

func Sum(x, y int) (int, error) {
	sum := x + y

	if sum >= 100 {
		return sum, errors.New("Sum is more than 100")
	}

	return sum, nil
}

func main() {
	sum, err := Sum(80, 20)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum = ", sum)
}
