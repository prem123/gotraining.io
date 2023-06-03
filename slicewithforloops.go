package main

import "fmt"

func main() {
	cards := []string{"red", "blue", "black"}
	for index, card := range cards {
		fmt.Println(index, card)
	}
}
