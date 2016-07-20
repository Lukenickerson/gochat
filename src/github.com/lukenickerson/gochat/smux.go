// Simple Mux (HTTP request multiplexer)
package main

import (
	//"net"
	"net/http"
	"io/ioutil" // read/write files
	"log"		// write to the screen (?)
	//"fmt"
	"strings"
	"strconv"
	"github.com/cbroglie/mustache"
)

var count int = 0

// ~ Object prototype
type MuxHandler struct {
}

// ~ Add a method into the MuxHandler prototype
func (this *MuxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	count = count + 1
	var path = r.URL.Path[1:] // do a slice to remove leading "/"
	var ip = r.RemoteAddr
	//var addresses, err = net.InterfaceAddrs()

	// strconv.Itoa -- the most confusing name for "convert int to string"
	log.Println(strconv.Itoa(count) + ": " + ip + " " + r.URL.Path) 

	if (strings.HasPrefix(path, "api/")) {
		w.Header().Add("Content Type", "application/json")
		w.Write([]byte(`{"m": "API Response goes here"}`))
		addMessage("Ok")
	} else if (strings.HasPrefix(path, "public/")) {
		data, err := ioutil.ReadFile(string(path))
		if err == nil {
			var contentType string = getContentTypeByPath(path)
			w.Header().Add("Content Type", contentType)
			w.Write(data)
		} else {
			writeError(w, 404)
		}
	} else if path == "" {
		// Look for the index template and populate it
		type pageDataType struct {
			Messages []Message
			Name string
		}
		pageData := new(pageDataType)
		pageData.Messages = Messages
		pageData.Name = "Person"

		writeTemplate(w, "index", pageData)
	} else {
		writeError(w, 403)
	}

}



func writeTemplate(w http.ResponseWriter, templateName string, data interface{}) {
	pageData, err := ioutil.ReadFile(string("templates/" + templateName + ".html"))
	
	if err == nil {
		w.Header().Add("Content Type", "text/html")
		html := string(pageData[:]) // convert bytes to string
		
		//mustacheData, mustacheErr := mustache.Render(html, map[string]string{"c": "world", "messages": messages[0].content})
		mustacheData, mustacheErr := mustache.Render(html, data)

		if mustacheErr == nil {
			w.Write([]byte(mustacheData))
		} else {
			// Should we return the un-filled template or an error?
			//w.Write(pageData)
			writeError(w, 500)
		}
	} else {
		// Template not found
		writeError(w, 404)
	}
}


func writeError(w http.ResponseWriter, errorCode int) {
	w.WriteHeader(errorCode)
	w.Write([]byte(strconv.Itoa(errorCode) + " - " + http.StatusText(errorCode)))
}

func getContentTypeByPath(path string) string{
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
	return contentType
}
