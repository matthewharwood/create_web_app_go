package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content Type", "text/html")
		templates := template.New("template")
		templates.New("test").Parse(doc)
		templates.New("header").Parse(header)
		templates.New("footer").Parse(footer)
		context := Context{
			[3]string{"lemon", "apple", "banana"},
			"the title",

		}
		templates.Lookup("test").Execute(w, context)

	})
	http.ListenAndServe(":8000", nil)
}

const doc = `
{{template "header" .Message}}
  <body>
    <h1>List of Fruit</h1>
    <ul>
      {{range .Fruit}}
      	<li>{{.}}</li>
      {{end}}
    </ul>
  </body>
{{template "footer"}}
`
const  header =  `
    <body>
		<h1>heyhey</h1>
		<h1>{{.}}</h1>
	</body>
`

const  footer = `
   <p> i'm a footer</p>
    </html>
`
type Context struct {
	Fruit [3]string
	Message string
}

