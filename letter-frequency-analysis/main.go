package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"unicode"
)

func main() {
	file, err := os.Open("shakespeare.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)

	// Create a map for the frequencys
	freqMap := make(map[string]int)

	// Count every instance of a letter but disregard case and anything that is not a letter
	for _, char := range b {
		//Convert Char to Uppercase
		upperCaseLetter := unicode.ToUpper(rune(char))

		if int(upperCaseLetter) >= 65 && int(upperCaseLetter) <= 90 {
			// Increment the count in the map
			freqMap[string(upperCaseLetter)] = freqMap[string(upperCaseLetter)] + 1
		}
	}

	// Dump out the frequencies in CSV format for simply plotting
	for char, count := range freqMap {
		fmt.Printf("%s,%d\n", char, count)
	}
}
