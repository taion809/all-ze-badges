package main

import "fmt"

const quantityOfBadges = "all of them"

func howManyBadges() string {
	return quantityOfBadges
}

func main() {
	fmt.Println("How many badges do we have: ", howManyBadges())
}
