package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"sort"
)

type record struct {
	Key   [10]byte
	Value [90]byte
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 3 {
		log.Fatalf("Usage: %v inputfile outputfile\n", os.Args[0])
	}

	log.Printf("Sorting %s to %s\n", os.Args[1], os.Args[2])

	// Read input file name
	inputFileName := os.Args[1]

	// Open input file
	inputFile, err := os.Open(inputFileName)

	if err != nil {
		log.Fatalf("Error in opening input file - %v", err)
	}

	// Declare records array which will store individual record
	records := []record{}
	for {
		var key [10]byte
		var value [90]byte
		// Read first 10 bytes into key
		_, err := inputFile.Read(key[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		// Read the next 90 bytes into value
		_, err = inputFile.Read(value[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		// Create record object with key and value
		rec := record{key, value}
		// Append record to the records array
		records = append(records, rec)
	}

	// Close input file
	inputFile.Close()

	// Custom comparator for sorting records array by key
	sort.Slice(records, func(i, j int) bool {
		// Sort the two records by the key
		isLessThan := bytes.Compare(records[i].Key[:], records[j].Key[:])
		if isLessThan <= 0 {
			return true
		}
		return false
	})

	// Read write file name
	writeFileName := os.Args[2]
	// Create output file
	outputFile, err := os.Create(writeFileName)
	if err != nil {
		log.Fatalf("Error creating output file - %v", err)
	}
	// Writing records to the output file
	for _, rec := range records {
		// Write key into the file
		_, err = outputFile.Write(rec.Key[:])
		if err != nil {
			log.Println(err)
		}
		// Write value into the file
		_, err = outputFile.Write(rec.Value[:])
		if err != nil {
			log.Println(err)
		}
	}
	// Closing the output file
	outputFile.Close()
}
