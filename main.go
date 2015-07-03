package main

import (
	"net/http"
	"text/template"
	"os"

)

func main() {
	templates := populateTemplate()
	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request){
			requestedFile := req.URL.Path[1:]
			template := templates.Lookup(requestedFile + ".html")
			if template != nil {
				template.Execute(w, nil)
			} else {
				w.WriteHeader(404)
			}

		})
	http.ListenAndServe(":8000", nil)
}

func populateTemplate() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	//opens the folder using the os
	templateFolder, _ := os.Open(basePath)
	// will close the opening of the templatefolder when populateTemplate is finished.
	defer templateFolder.Close()

	// reads all the file names in a folder
	// -1 means until none left
	templatePathRaw, _ := templateFolder.Readdir(-1)

	//loops through the raw path strings and sets to a slice
	//while adding the base path behind.
	templatePaths := new([]string)
	for _, pathInfo := range templatePathRaw {
		*templatePaths = append(*templatePaths,
			basePath + "/" + pathInfo.Name())
	}

	//helper in go to parse a set of files.
	//uses ... to go through 1 by 1
	result.ParseFiles(*templatePaths...)
	return result
}