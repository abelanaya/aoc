package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const dialInitial = 50
const dialMin = 0
const dialMax = 100

func dialOperate(dial int, operation string, count int, dialNumberZeroes *int) int {
	switch operation {
	case "R":
		*dialNumberZeroes += (dial + count) / dialMax
		dial = (dial + count) % dialMax

	case "L":
		if dial == 0 {
			*dialNumberZeroes += count / dialMax
		} else if count >= dial {
			*dialNumberZeroes += 1 + (count-dial)/dialMax
		}

		dial = ((dial-count)%dialMax + dialMax) % dialMax

	default:
		fmt.Println("Invalid operation: . Operation count: . Current dial: ", operation, count, dial)
		os.Exit(1)
		return 0
	}

	return dial
}

func main() {
	fp, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	dialNumberZeroes := 0
	dialCurrent := dialInitial

	for scanner.Scan() {
		line := scanner.Text()
		operation := string(line[0])
		count, err := strconv.Atoi(line[1:])

		if err != nil {
			fmt.Println("Error: ", err)
		}

		dialCurrent = dialOperate(dialCurrent, operation, count, &dialNumberZeroes)
	}

	fmt.Println("Last dial is: ", dialCurrent)
	fmt.Println("Number of zeroes is: ", dialNumberZeroes)
}
