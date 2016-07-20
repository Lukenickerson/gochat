package main

import (
	"net/http"
	"io/ioutil" // read/write files
	"log"		// write to the screen (?)
	"fmt"
	"strings"
	"strconv"
)

var count int = 0

// ~ Object prototype
type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count = count + 1
	path := r.URL.Path[1:] // do a slice
	// strconv.Itoa -- the most confusing name for "convert int to string"
	log.Println(path + " " + strconv.Itoa(count)) 

	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else if strings.HasSuffix(path, ".svg") {
			contentType = "image/svg+xml"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content Type", contentType)
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}
}

// Can't assign in global scope using infered type
// x := 5  // will initialize and infer the type
// var x int = 5 // explicitly set the type and the value

func main() {
	fmt.Printf("Running ServeMux\n")
	//myMux := http.NewServeMux()
	//myMux.HandleFunc("/", router)
	//http.ListenAndServe(":8080", myMux)

	http.Handle("/", new(MyHandler))
	http.ListenAndServe(":8080", nil)
}

func router(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("<html><body><h1>Hello Self</h1></body></html>"))
}
