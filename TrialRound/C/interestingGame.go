package main

import "fmt"

func main() {
	var pointsToWin, cardsNumber int

	fmt.Scan(&pointsToWin, &cardsNumber)

	cards := setCards(cardsNumber)

	var winner string
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
			winner = "Draw"
		} else if vasyaPoints == pointsToWin {
			winner = "Vasya"
		} else if petyaPoints == pointsToWin {
			winner = "Petya"
		}

		if winner != "" {
			break
		}
	}

	fmt.Println(winner)
}

func setCards(n int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	return a
}
