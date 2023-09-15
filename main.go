package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func averageCalc(num []int) int {
	var sum float64
	for i := 0; i < len(num); i++ {
		sum += float64(num[i])
	}
	average := float64(sum) / float64((len(num)))

	return int(math.Round(average))
}

func medianCalc(num []int) int {
	bubbleSort(num)
	return num[len(num)/2]
}

func bubbleSort(list []int) {

	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list); j++ {
			if list[i] < list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func summ(num []int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum += num[i]
	}
	return sum
}

func variancer(num []int) int {
	mean := summ(num) / len(num)

	sqDifference := 0

	for i := 0; i < len(num); i++ {
		sqDifference += (num[i] - mean) * (num[i] - mean)
	}
	return sqDifference / len(num)
}

func standardDeviation(num []int) int {
	base := float64(variancer(num))
	return int(math.Round(math.Sqrt(base)))
}

func main() {
	if len(os.Args) == 2 {
		var intArray []int
		dataFileName := os.Args[1]
		dataRow, _ := os.ReadFile(dataFileName)
		numList := strings.Split(string(dataRow), "\n")
		for j := 0; j < len(numList)-1; j++ {
			number, _ := strconv.Atoi(numList[j])
			intArray = append(intArray, number)
		}
		// fmt.Println("----------------------------------")
		// fmt.Println(intArray)
		// fmt.Println("----------------------------------")
		// fmt.Println(intArray)
		fmt.Println("----------------------------------")
		fmt.Println("Average: ", averageCalc(intArray))
		fmt.Println("----------------------------------")
		fmt.Println("Median: ", medianCalc(intArray))
		fmt.Println("----------------------------------")
		fmt.Println("Variance: ", variancer(intArray))
		fmt.Println("----------------------------------")
		fmt.Println("Standard Deviation: ", standardDeviation(intArray))
		fmt.Println("----------------------------------")
	} else {
		fmt.Println("File not found or too many added")
		return
	}
}
