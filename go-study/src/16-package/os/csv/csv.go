package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {
	res := readCsv("./src/conf/test.csv")
	printCsv(res)
	writeCsv("./src/conf/res.csv", res)
}

func readCsv(filename string) [][]string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("read file err: %s\n", err)
	}

	reader := csv.NewReader(file)
	reader.TrimLeadingSpace = true
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("recodes err: %s\n", err)
	}

	return records
}

func printCsv(records [][]string) {
	for _, record := range records {
		fmt.Printf("%v\n", record)
	}
}

func writeCsv(filename string, data [][]string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create file\n", err)
	}
	w := csv.NewWriter(file)
	err = w.WriteAll(data)
	if err != nil {
		log.Fatalf("Failed to write csv file: \n", err)
	}
	w.Flush()
}
