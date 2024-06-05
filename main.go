package main

import (
	"fmt"
	"log"
	"math"
	"os"

	stat "math-skills/pkg"
)

func main() {
	// check to allow only 2 argument
	if len(os.Args) != 2 {
		fmt.Println("usage <go run . data.txt >")
		return
	}

	// selects first arg to read from
	filepath := os.Args[1]
	data := stat.ReadDataFromFile(filepath)
	if len(data) == 0 {
		log.Fatalf("\n %s lines are empty", filepath)
	}
	// calculates statistics
	avg := stat.CalculateAverage(data)
	med := stat.CalculateMedian(data)
	variance := stat.CalculateVariance(data, avg)
	stdDev := stat.CalculateStandardDeviation(variance)

	fmt.Println("Average:", int(math.Round(avg)))
	fmt.Println("Median:", int(math.Round(med)))
	fmt.Println("Variance:", int(math.Round(variance)))
	fmt.Println("Standard Deviation:", int(math.Round(stdDev)))
}
