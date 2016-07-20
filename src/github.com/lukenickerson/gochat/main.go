/*
[x] get mustahce
[ ] Get IP address of client
[x] Message object: text, ip
[x] Array of message objs
[x] Populate template with data

[ ] api to add data
[ ] ui to hit api and add data

[ ] name/unique id login
[ ] Keep list of unique names
[ ] Read json data in external file of pokemon names
[ ] Use pokemon names + IP for unique identification
*/
package main

import (
	//"net"
	"net/http"
	//"io/ioutil" // read/write files
	//"log"		// write to the screen (?)
	"fmt"
	//"strings"
	//"strconv"
	//"github.com/cbroglie/mustache"
)


// Can't assign in global scope using infered type
// x := 5  // will initialize and infer the type
// var x int = 5 // explicitly set the type and the value

func main() {

	fmt.Printf("Running ServeMux. Waiting for requests on localhost...\n")
	//myMux := http.NewServeMux()
	//myMux.HandleFunc("/", router)
	//http.ListenAndServe(":8080", myMux)

	http.Handle("/", new(MuxHandler))
	http.ListenAndServe(":8080", nil)
}

