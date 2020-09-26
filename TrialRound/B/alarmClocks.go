package main

import (
	"fmt"
	"math"
)

func main() {
	var alarmClocksNumber, freqCalls, alarmClocksToWakeUp int

	fmt.Scan(&alarmClocksNumber, &freqCalls, &alarmClocksToWakeUp)

	alarmsTimers := setAlarmsTimers(alarmClocksNumber)

	alarmsTimers = deleteDupAlarms(alarmsTimers, freqCalls)

	// sort.Slice(alarmsTimers, func(i, j int) bool {
	// 	return alarmsTimers[i]%freqCalls < alarmsTimers[j]%freqCalls
	// })

	fmt.Println(countTimeWakeUp(alarmsTimers, freqCalls, alarmClocksToWakeUp))
}

func setAlarmsTimers(n int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	return a
}

func deleteDupAlarms(alarmsTimers []int, freqCalls int) []int {
	for i := 0; i < len(alarmsTimers)-1; i++ {
		for j := i + 1; j < len(alarmsTimers); j++ {
			if alarmsTimers[i]%freqCalls == alarmsTimers[j]%freqCalls {
				if alarmsTimers[i] < alarmsTimers[j] {
					alarmsTimers = remove(alarmsTimers, j)
				} else {
					alarmsTimers = remove(alarmsTimers, i)
				}
			}
		}
	}

	return alarmsTimers
}

func remove(slice []int, pos int) []int {
	return append(slice[:pos], slice[pos+1:]...)
}

func alarmsRangToExactTime(alarmTimers []int, freqCalls, time int) int {
	alarmsRangCounter := 0
	alarmsRang := 0

	for _, alarmTimer := range alarmTimers {
		alarmsRang = (time - alarmTimer) / freqCalls

		if alarmsRang > 0 {
			alarmsRangCounter += alarmsRang + 1
		}

		if alarmsRang == 0 {
			alarmsRangCounter++
		}
	}

	return alarmsRangCounter
}

func countTimeWakeUp(alarmsTimers []int, freqCalls, alarmClocksToWakeUp int) int {
	low := 0
	median := 0
	high := math.MaxInt32

	for low <= high {
		median = (low + high) / 2

		alarmsRang := alarmsRangToExactTime(alarmsTimers, freqCalls, median)

		if alarmsRang < alarmClocksToWakeUp {
			low = median + 1
		} else {
			high = median - 1
		}

		if alarmsRang == alarmClocksToWakeUp {
			break
		}

	}

	return median
}
