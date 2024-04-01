package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"flag"
 )

func main() {

	file_name := flag.String("file", "problems.csv", "Problem file name. Default: problems.csv")
	/*
	Or:
	    var file_name string
    	flag.StringVar(&file_name, "file", "problems.csv", "a string var")
	*/

	flag.Parse() // Need to call this to get the actual flags in

   // Open the CSV file
   file, err := os.Open(*file_name)
   if err != nil {
       panic(err) // https://gobyexample.com/panic
   }
   defer file.Close() // https://gobyexample.com/defer

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2 // Allow variable number of fields - will throw an error if CSV is in wrong format (wrt fields)!
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Tracking for correct / incorrect values
	correct, incorrect := 0, 0

	// Main game loop
	for _, row := range data {
		// Print question
		fmt.Printf("%s", row[0])
		fmt.Println()

		// Convert expected ans to an integer - although I guess we could do this all as string comparisons.
		ans_str := row[1]
		ans, err := strconv.Atoi(ans_str)
		if err != nil {
			panic(err)
		}

		var user_input int
		fmt.Scanln(&user_input)

		if ans == user_input {
			fmt.Println("Correct!")
			correct += 1
		} else {
			fmt.Println("Incorrect!")
			incorrect += 1
		}

		fmt.Println()

	}

	fmt.Printf("\nGame finished!\nStats:\n" +
	"Correct Answers: %v\n" +
	"Incorrect Answers: %v",
	correct, incorrect)

}