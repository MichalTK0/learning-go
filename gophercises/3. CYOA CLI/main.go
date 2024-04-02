package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StorySection struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

func parse_json(filename string) (story map[string]*StorySection) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &story)

	return
}

func process_section(section StorySection) (choices map[int]string){
	choices = make(map[int]string)
	
	for _, s := range(section.Story) {
		fmt.Printf("%s\n\n", s)
	}

	for i, s := range(section.Options) {
		fmt.Printf("%v. %s\n\n", i, s.Text)
		choices[i] = s.Arc
	}

	return
}

func main() {

	story := parse_json("gopher.json")

	starting_section := story["intro"]
	choices := process_section(*starting_section)

	for len(choices) != 0 {
		var user_input int


		for {
			fmt.Scanln(&user_input)

			if _, ok := choices[user_input]; ok {
				break
			} else {
				fmt.Println("Invalid input.")
			}

		}
		
		fmt.Println()
		choices = process_section(*story[choices[user_input]])
	}


}
