package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Metric struct {
	Count int64
	Min   float64
	Max   float64
	Sum   float64
}

func printResults(metrics map[string]Metric) {
	// sort alphabetically
	keys := make([]string, 0, len(metrics))
	for k := range metrics {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	fmt.Print("{")
	i := 0
	metric := metrics[keys[i]]
	mean := metric.Sum / float64(metric.Count)
	fmt.Printf("%s=%.1f/%.1f/%.1f", keys[i], metric.Min, mean, metric.Max)
	for i < len(keys)-1 {
		i++
		fmt.Printf(",%s=%.1f/%.1f/%.1f", keys[i], metric.Min, mean, metric.Max)
	}
	fmt.Print("}")
}

func main() {
	file, err := os.Open("measurements.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	metrics := make(map[string]Metric)

	for scanner.Scan() {
		line := scanner.Text()
		splitted := strings.Split(line, ";")
		weatherstation := splitted[0]
		temperature, err := strconv.ParseFloat(splitted[1], 64)
		if err != nil {
			log.Fatal(err)
		}

		if metric, exists := metrics[weatherstation]; exists {
			var min float64
			if temperature < metric.Min {
				min = temperature
			} else {
				min = metric.Min
			}
			var max float64
			if temperature > metric.Max {
				max = temperature
			} else {
				max = metric.Min
			}
			metrics[weatherstation] = Metric{
				Count: metric.Count + 1,
				Min:   min,
				Max:   max,
				Sum:   temperature + metric.Sum,
			}
		} else {
			metrics[weatherstation] = Metric{
				Count: 1,
				Min:   temperature,
				Max:   temperature,
				Sum:   temperature,
			}
		}
	}
}
