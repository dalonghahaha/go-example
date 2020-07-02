package main

import (
	"log"
	"os"
	"text/template"
)

const letter = `
Dear {{.Name}},

{{if .Attended}}
	It was a pleasure to see you at the wedding.{{else}}
	It is a shame you couldn't make it to the wedding.{{end}}
	{{with .Gift}}Thank you for the lovely {{.}}.
{{end}}

											Best wishes,
											Josie
`

type Recipient struct {
	Name     string
	Gift     string
	Attended bool
}

func main() {
	var recipients = []Recipient{
		Recipient{
			Name:     "Aunt Mildred",
			Gift:     "bone china tea set",
			Attended: true,
		},
		Recipient{
			Name:     "Uncle John",
			Gift:     "moleskin pants",
			Attended: false,
		},
		Recipient{
			Name:     "Cousin Rodney",
			Gift:     "",
			Attended: false,
		},
	}
	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))
	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}
}
