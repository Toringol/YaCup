package main

import "fmt"

func main() {
	var pointsToWin, cardsNumber int

	fmt.Scan(&pointsToWin, &cardsNumber)

	cards := setCards(cardsNumber)

	fmt.Println(whoIsWinner(cards, pointsToWin))
}

func setCards(n int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	return a
}

func whoIsWinner(cards []int, pointsToWin int) string {
	petyaPoints, vasyaPoints := 0, 0

	for _, card := range cards {
		if card%15 == 0 {
			continue
		} else if card%5 == 0 {
			vasyaPoints++
		} else if card%3 == 0 {
			petyaPoints++
		}

		if vasyaPoints == petyaPoints && vasyaPoints == pointsToWin {
			return "Draw"
		} else if vasyaPoints == pointsToWin {
			return "Vasya"
		} else if petyaPoints == pointsToWin {
			return "Petya"
		}
	}

	if vasyaPoints > petyaPoints {
		return "Vasya"
	} else if petyaPoints > vasyaPoints {
		return "Petya"
	} else {
		return "Draw"
	}
}
