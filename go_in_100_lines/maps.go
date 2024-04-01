package main

import "fmt"

func main() {
	/* Define a map containing the release year of several languages */
	firstReleases := map[string]int{
		"C": 1972, "C++": 1985, "Java": 1996,
		"Python": 1991, "JavaScript": 1996, "Go": 2012,
	}

	/* Loop through each entry and output the name and release year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first released in %d\n", k, v)
		// Note the random order. 
	}

	// You can also use the make keyword to make a map and use it later
	/*
		patrons := make(map[int]string)

		patrons[0] = "Terrence"
		patrons[1] = "Evelyn"
	*/
}