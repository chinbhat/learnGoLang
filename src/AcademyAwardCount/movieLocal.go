package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

func main() {

	const (
		fileName = "/tmp/academy_awards.csv"
	)

	type movieNomin struct {
		name  string
		count int
	}
	nomineeMap, winnerMap := convertNomineeToMap(fileName)

	if nomineeMap == nil {
		panic("Error while Parsing CSV File.")
	}

	var nomineeSlice []movieNomin
	var winnerSlice []movieNomin

	for k, v := range *nomineeMap {
		nomineeSlice = append(nomineeSlice, movieNomin{k, v})
	}

	for k, v := range *winnerMap {
		winnerSlice = append(winnerSlice, movieNomin{k, v})
	}

	sort.Slice(nomineeSlice, func(i, j int) bool {
		return nomineeSlice[i].count > nomineeSlice[j].count
	})

	sort.Slice(winnerSlice, func(i, j int) bool {
		return winnerSlice[i].count > winnerSlice[j].count
	})

	for _, v := range nomineeSlice {
		if v.count != -1 {
			fmt.Printf("%v has been nominated %v times, without winning Academy Awards once\n", v.name, v.count)
		}
	}

	for _, v := range winnerSlice {
		fmt.Printf("%v has won the academy awards %v times \n", v.name, v.count)
	}

}

func convertNomineeToMap(csvFileName string) (*map[string]int, *map[string]int) {

	fmt.Println("Inside the function" + csvFileName)

	fileHandler, err := os.Open(csvFileName)
	checkError(err)

	defer fileHandler.Close()

	csvHandler := csv.NewReader(fileHandler)

	nomineeMap := make(map[string]int)
	winnerMap := make(map[string]int)

	for {
		csvLine, err := csvHandler.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		} else {
			entryCriteria := strings.Contains(csvLine[1], "Actor") || strings.Contains(csvLine[1], "Actress")

			if entryCriteria == true {

				name := csvLine[2]

				if csvLine[4] == "NO" {
					if nomineeMap[name] != -1 {
						nomineeMap[name] = nomineeMap[name] + 1
					}
				} else {
					winnerMap[name] = winnerMap[name] + 1
					nomineeMap[name] = -1
				}

			}

		}
	}
	return &nomineeMap, &winnerMap

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
