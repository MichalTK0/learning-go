package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"time"
	"strings"
	"math/rand"
 )


 func parse_csv(file_name string) (data [][]string) {

	// Open the CSV file
	file, err := os.Open(file_name)
	if err != nil {
		panic(err) // https://gobyexample.com/panic
	}
	defer file.Close() // https://gobyexample.com/defer

	// Read the CSV data
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 2 // Allow variable number of fields - will throw an error if CSV is in wrong format (wrt fields)!
	data, err = reader.ReadAll()
	if err != nil {
		panic(err)
	}

	return
}

func clean_string(str string) (ret string) {
	ret = strings.ToLower(strings.TrimSpace(str))
	return
}

func quiz_game(data [][]string, duration int) (correct int) {

    // Create a channel to receive user input
    userInput := make(chan string)
    defer close(userInput)

    // Create a channel to receive timeout signal
    timeout := make(chan bool)
    defer close(timeout)

    fmt.Printf("Press enter to start! Time limit=%ds\n", duration)
    fmt.Scanln()

    // Start the timer in its own go routine. Set timeout channel to true when finished.
    go func() {
        time.Sleep(time.Duration(duration) * time.Second)
        timeout <- true // < -- sets timeout channel to true
    }() // <-- these brackets mean the function instantly starts and does not need to be explicitly called.

    // Main game loop
    for _, row := range data {
        // Print question
        fmt.Printf("%s\n", row[0])

		ans := clean_string(row[1])

        // Prompt user for input
        go func() {
            var user_input string
            fmt.Scanln(&user_input)
            userInput <- clean_string(user_input)
        }()

        // Wait for user input or timeout
        select { // A select statement is like a switch for concurrency, and acts depending on channel input/output states.
        case <-timeout:
			// If the timeout channel has been set, exit.
            fmt.Println("\nTime's up!")
            return correct
        case user_input := <-userInput:
			// if input has been written from the userInput channel to the user input var, check the answer.
            // Check the answer
            if ans == user_input {
                fmt.Println("Correct!")
                correct++
            } else {
                fmt.Println("Incorrect!")
            }
            fmt.Println()
        }
    }

    return correct
}

func randomise_slice(slc [][]string) ([][]string){

	r := rand.New(rand.NewSource(time.Now().Unix()))
	ret := make([][]string, len(slc))
	perm := r.Perm(len(slc))

	for i, randIndex := range perm {
	  ret[i] = slc[randIndex]
	}

	return ret
}

func main() {
	// args
	// flag name, default value, help string
	file_name := flag.String("file", "problems.csv", "Problem file name. Default: problems.csv")
	time := flag.Int("time", 30, "Quiz time limit. Default: 30 (seconds)")
	random := flag.Bool("random", false, "Randomises order of questions. Default: false")

	/*
	Or:
	    var file_name string
    	flag.StringVar(&file_name, "file", "problems.csv", "a string var")
	*/

	flag.Parse() // Need to call this to get the actual flags in

	// the flag returns a pointer, so need to use * to get the actual value. Uni C flashbacks. Oh lord.
	data := parse_csv(*file_name)

	// Shuffle if set
	if *random {
		data = randomise_slice(data)
	}

	// Tracking for correct / incorrect values
	correct := quiz_game(data, *time)
	incorrect := len(data) - correct

	fmt.Printf("\nGame finished!\nStats:\n" +
	"Correct Answers: %v\n" +
	"Incorrect Answers: %v",
	correct, incorrect)

}