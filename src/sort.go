package main

import (
	"log"
	"os"
	"io"
	"encoding/binary"
	"sort"
	"unsafe"
)

type record struct {
	Key [10]byte 
	Value [10]byte
}

func checkByteOrder() binary.ByteOrder {
	var byteOrder binary.ByteOrder
	buffer := [2]byte{}
	*(*uint16)(unsafe.Pointer(&buffer[0])) = uint16(0xABCD)

	switch buffer {
	case [2]byte{0xCD, 0xAB}:
		byteOrder = binary.LittleEndian
	case [2]byte{0xAB, 0xCD}:
		byteOrder = binary.BigEndian
	default:
		log.Fatalln("Cannot determine the byte order")
	}
	return byteOrder
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

	records := []record{}
	byteOrder := checkByteOrder()
	for {
		rec := record{}
		err = binary.Read(inputFile, byteOrder, &rec)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Println(err)
		}
		// Append record to the records array
		records = append(records, rec)
	}
	// Close the input file
	inputFile.Close()

	// Custom comparator for sorting records array by key
	sort.Slice(records, func(i, j int) bool {
		// Sort the two recrods by the key
		return string(records[i].Key[:]) < string(records[j].Key[:])
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
		binary.Write(outputFile, byteOrder, rec)
	}
	// Closing the output file
	outputFile.Close()
}
