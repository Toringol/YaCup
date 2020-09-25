package main

import (
	"fmt"
	"strings"
)

func main() {
	var jewelry string
	var stones string

	fmt.Scan(&jewelry, &stones)

	conjuctionInc := 0

	for _, stone := range stones {
		if strings.Contains(jewelry, string(stone)) {
			conjuctionInc++
		}
	}

	fmt.Println(conjuctionInc)
}
