package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Content Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil {
			context := Context{"message wow"}
			tmpl.Execute(w, context)
		}
	})
	http.ListenAndServe(":8000", nil)
}

const doc  = `
<!DOCTYPE>
<html>
	<head>
		<title>example</title>
	</head>
	<body>
		<h1>heyhey</h1>
		<h1>{{.Message}}</h1>
	</body>
</html>

`

type Context struct {
	Message string
}

