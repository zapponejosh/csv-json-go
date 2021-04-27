package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

var (
	delimiter string
	output    *os.File
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getFile(args []string, delimiter string) [][]string {
	filename := args[0]

	raw, err := os.Open(filename)
	check(err)
	reader := csv.NewReader(raw)

	switch delimiter { // defaults to comma, optionally use tabs or pipes
	case "tab":
		reader.Comma = '\t'
	case "pipe", "|":
		reader.Comma = '|'
	default:
	}

	data, e := reader.ReadAll()
	check(e)

	return data
}

func processCSV(data [][]string) []map[string]string {
	processedData := make([]map[string]string, len(data)-1)
	for i, row := range data {
		if i > 0 {
			newMap := make(map[string]string)
			for k, val := range row {
				newMap[data[0][k]] = val
			}
			processedData[i-1] = newMap
		}
	}
	return processedData
}

func main() {
	fmt.Println("running app...")

	//  command line  args
	flag.StringVar(&delimiter, "d", ",", "file delimiter: tabs, pipes, or commas")
	flag.Parse()
	args := flag.Args()

	if len(args) == 2 {
		out, err := os.Create(args[1])
		output = out
		check(err)
	} else {
		output = os.Stdout
	}

	// read, transform, and write
	csvData := getFile(args, delimiter)
	processedData := processCSV(csvData)
	err := json.NewEncoder(output).Encode(processedData)
	check(err)
	fmt.Println("complete")

}
