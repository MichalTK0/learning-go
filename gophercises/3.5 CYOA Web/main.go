package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
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

func parseJSON(filename string) (story map[string]*StorySection) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &story)

	return
}

func renderSection(w http.ResponseWriter, section *StorySection) {
	tmpl := `
		<h1>{{.Title}}</h1>
		{{range .Story}}
			<p>{{.}}</p>
		{{end}}
		<ul>
			{{range $index, $option := .Options}}
				<li><a href="/{{ $option.Arc }}">{{ $option.Text }}</a></li>
			{{end}}
		</ul>
	`

	if len(section.Options) == 0 {
		end_option := StoryOption{Text: "Back to start."}
		section.Options = append(section.Options, end_option )
	}

	t := template.Must(template.New("section").Parse(tmpl))
	err := t.Execute(w, section)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Handler function for each story section
func sectionHandler(story map[string]*StorySection) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arc := r.URL.Path[1:] // Extract the arc from the URL
		if section, ok := story[arc]; ok {
			renderSection(w, section)
		} else {
			http.NotFound(w, r)
		}
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intro", http.StatusSeeOther)
}


func main() {
	story := parseJSON("gopher.json")

	// Register view handler
	http.HandleFunc("/", viewHandler)

	// Register handler function for each story section
	for arc := range story {
		http.HandleFunc("/"+arc, sectionHandler(story))
	}

	log.Fatal(http.ListenAndServe(":8000", nil))
}

