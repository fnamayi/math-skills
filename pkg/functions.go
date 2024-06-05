package pkg

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// this function reads the file parsed
func ReadDataFromFile(filepath string) []float64 {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileinfo, err := file.Stat()
	if fileinfo.Size() == 0 && err == nil {
		fmt.Printf("%s is empty.\n", fileinfo.Name())
		os.Exit(1)
	}

	scanner := bufio.NewScanner(file)

	var data []float64
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		line = strings.ReplaceAll(line, ",", "")
		if line == "" {
			continue
		}

		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			log.Fatalf("\nError parsing number:\n %v", err)
		}

		//checks overflow numbers
		if num > (math.Pow(2, 53)-1) || num < -(math.Pow(2, 53)-1) {
			log.Fatalf("\n %f is an overflow", num)
		}
		/* checks overflow using bitwise operator left shift
		if num > (1<<53)-1 || num < -(1-53){
			log.Fatalf("\n %f  is an overflow ",num)
		}*/

		data = append(data, num)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

// this func calculates mean/average
func CalculateAverage(data []float64) float64 {
	sum := 0.0
	for _, n := range data {
		sum += n
	}
	return sum / float64(len(data))
}

// this func calculates mediaan
func CalculateMedian(data []float64) float64 {
	sort.Float64s(data)
	dataLen := len(data)
	mid := dataLen / 2

	if dataLen%2 == 0 {
		return (data[mid-1] + data[mid]) / 2
	}
	return data[mid]
}

// func calculates variance
func CalculateVariance(data []float64, avg float64) float64 {
	sum := 0.0
	for _, dataPoint := range data {
		sum += math.Pow(dataPoint-avg, 2)
	}

	return sum / float64(len(data))
}

func CalculateStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
