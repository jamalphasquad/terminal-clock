package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

var (
	//⚪ ⚫ 🟥 🔴 ⬜ ⬛
	visible   = "⬜"
	invisible = "⬛"
)

func getTimeArray() [][][]int8 {
	timeArray := [][][]int8{
		{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 0, 0, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}},
		{{0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}},
		{{1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 1, 1, 1}},
		{{1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}},
		{{1, 0, 0, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}},
		{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}},
		{{1, 1, 1, 1}, {1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}},
		{{1, 1, 1, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}, {0, 0, 0, 1}},
		{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}},
		{{1, 1, 1, 1}, {1, 0, 0, 1}, {1, 1, 1, 1}, {0, 0, 0, 1}, {1, 1, 1, 1}},
	}

	return timeArray
}

func getDoubleDigits(num int) [5][9]int8 {

	if num > 60 {
		log.Fatalf(" num is greater than 99 %d", num)
	}

	digits := getTimeArray()
	var digitsArr [5][9]int8
	digStr := strconv.Itoa(num)

	if num < 10 {
		digStr = "0" + digStr
	}

	first, _ := strconv.Atoi(string(digStr[0]))
	second, _ := strconv.Atoi(string(digStr[1]))

	for ind, digit := range digits[first] {
		digitsArr[ind] = [9]int8(append(append(digit[:], 0), digits[second][ind]...))
	}

	return digitsArr
}

func main() {

	ticker := time.Tick(time.Second)
	log.SetFlags(0)

	for {
		<-ticker

		currentTime := time.Now()
		hours := currentTime.Hour()
		minutes := currentTime.Minute()
		seconds := currentTime.Second()
		amPm := currentTime.Format("PM")

		_ = amPm

		secondsArray := getDoubleDigits(seconds)
		minutesArray := getDoubleDigits(minutes)
		hoursArray := getDoubleDigits(hours)

		fmt.Print("\033[u\033[K")

		ch1 := make(chan bool)
		ch2 := make(chan bool)
		ch3 := make(chan bool)

		go func() {
			for ind, rows := range hoursArray {
				<-ch1
				for _, elem := range rows {
					if elem == 1 {
						fmt.Print("\033[31m", visible, "\033[0m")
					} else {
						fmt.Print(invisible)
					}
				}
				if ind == 1 || ind == 3 {
					fmt.Print(invisible + visible + invisible)
				} else {
					fmt.Print(invisible + invisible + invisible)
				}
				ch2 <- true
			}
		}()

		go func() {
			for ind, rows := range minutesArray {
				<-ch2
				for _, elem := range rows {
					if elem == 1 {
						fmt.Print("\033[31m", visible, "\033[0m")
					} else {
						fmt.Print(invisible)
					}
				}
				if ind == 1 || ind == 3 {
					fmt.Print(invisible + visible + invisible)
				} else {
					fmt.Print(invisible + invisible + invisible)
				}
				ch3 <- true
			}
		}()

		go func() {
			for _, rows := range secondsArray {
				<-ch3
				for _, elem := range rows {
					if elem == 1 {
						fmt.Print("\033[31m", visible, "\033[0m")
					} else {
						fmt.Print(invisible)
					}
				}
				fmt.Println("")
				ch1 <- true
			}
		}()
		ch1 <- true
	}

}
