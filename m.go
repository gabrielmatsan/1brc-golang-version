package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Measurement struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int64
}

func main() {
	start := time.Now()
	// abrindo arquivo
	mms, err := os.Open("measurements.txt")

	// verifica erros
	if err != nil {
		panic(err)
	}
	defer mms.Close()

	data := make(map[string]Measurement)

	scanner := bufio.NewScanner(mms)

	for scanner.Scan() {
		rawData := scanner.Text()

		semicolon := strings.Index(rawData, ";")

		location := rawData[:semicolon]
		temp, err := strconv.ParseFloat(rawData[semicolon+1:], 64)

		if err != nil {
			panic(err)
		}

		mm, ok := data[location]

		if !ok {
			mm = Measurement{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {

			mm.Sum += temp
			mm.Count++
			mm.Min = min(mm.Min, temp)
			mm.Max = max(mm.Max, temp)

		}

		data[location] = mm

	}

	locations := make([]string, 0, len(data))

	for name := range data {
		locations = append(locations, name)
	}

	sort.Strings(locations)

	fmt.Println("{")
	for _, name := range locations {
		mm := data[name]
		fmt.Printf("%s=%.1f/%.1f/%.1f\n", name, mm.Min, mm.Max, mm.Sum/float64(mm.Count))
	}
	fmt.Println("}\n")
	fmt.Println(time.Since(start))
}
