package main

import (
	"fmt"
	"net/http"

	"strings"

	"github.com/satori/go.uuid"
)

func init() {
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	u := uuid.NewV4()
	w.Header().Set("Content-type", "text/html")	
	fmt.Fprintf(w, "<div>this is demo code for session 4 to generate GUID. <br> newly generated UUID is  %v</div>", u.String())
	
}
