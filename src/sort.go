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

	// Read input file
	inputFile, err := os.Open(inputFileName)

	if err != nil {
		log.Fatalf("Error in opening input file - %v", err)
	}
	// data, err := ioutil.ReadFile(inputFileName)

	// if err != nil {
	// 	log.Fatalf("Error in reading input file - %v", err)
	// }

	// Declare records array which will store individual record
	records := []record{}
	for {
		// var key [10]byte
		// // read first 10 bytes and copy it into key
		// copy(key[:], data[:10])
		// // remove first 10 bytes from data
		// data = data[10:]
		// var value [90]byte
		// // read first 90 bytes and copy it into value
		// copy(value[:], data[:90])
		// // remove first 90 bytes from data
		// data = data[90:]
		// // Create record object with key and value
		// rec := record{key, value}
		// // Append record to the records array
		// records = append(records, rec)

		var key [10]byte
		var value [90]byte
		_, err := inputFile.Read(key[:])
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
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
		_, err = outputFile.Write(rec.Key[:])
		_, err = outputFile.Write(rec.Value[:])
		if err != nil {
			log.Println(err)
		}
	}
	// Closing the output file
	outputFile.Close()
}
