package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkInvalidIdV2(num int) bool {
	numAsText := fmt.Sprintf("%d", num)

	var sequence string

	for _, c := range numAsText {

		if len(sequence) != 0 {
			repetitions := strings.Count(numAsText, sequence)

			if len(numAsText) == repetitions*len(sequence) {
				return true
			}
		}

		sequence += string(c)
	}

	return false
}

func countInvalidIds(first int, last int, idsAddition *int, invalidIdsCountV2 *int, idsAdditionV2 *int) int {

	var invalidIds int

	for num := first; num <= last; num++ {
		if len(fmt.Sprintf("%d", num))%2 == 0 {
			numAsText := fmt.Sprintf("%d", num)
			firstHalf := numAsText[0:(len(numAsText) / 2)]
			secondHalf := numAsText[(len(numAsText) / 2):]

			if strings.Compare(firstHalf, secondHalf) == 0 {
				*idsAddition += num
				invalidIds++
			}
		}

		if checkInvalidIdV2(num) {
			*invalidIdsCountV2++
			*idsAdditionV2 += num
		}
	}

	return invalidIds
}

func main() {
	data, err := os.ReadFile("inputs/2/input.txt")

	if err != nil {
		log.Println("Error: ", err)
	}

	invalidIdsCount := 0
	invalidIdsCountV2 := 0

	idsAddition := 0
	idsAdditionV2 := 0

	rows := strings.SplitSeq(strings.TrimSpace(string(data)), ",")

	for v := range rows {
		var first, last int
		first, err = strconv.Atoi(strings.Split(v, "-")[0])

		if err != nil {
			log.Println("Invalid first element: ", err)
		}

		last, err = strconv.Atoi(strings.Split(v, "-")[1])

		if err != nil {
			log.Println("Invalid last element: ", err)
		}

		invalidIdsCount += countInvalidIds(first, last, &idsAddition, &invalidIdsCountV2, &idsAdditionV2)
	}

	log.Println("Invalid Ids:", invalidIdsCount)
	log.Println("Addition of invalid Ids: ", idsAddition)

	log.Println("Invalid Ids V2: ", invalidIdsCountV2)
	log.Println("Addition of invalid Ids V2: ", idsAdditionV2)
}
