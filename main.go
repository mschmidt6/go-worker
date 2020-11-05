package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

type imageRequest struct {
	Image string `json:"image"`
}

func tagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "test %s", r.URL.Path[1:])
		log.Printf("%s Request %s", r.Method, r.URL.Path[1:])
	case "POST":
		log.Printf("%s Request %s", r.Method, r.URL.Path[1:])
		var imgReq imageRequest
		err := decodeJSONBody(w, r, &imgReq)
		if err != nil {
			var mr *malformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.msg, mr.status)
			} else {
				log.Println(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		fmt.Fprintf(w, "Image %+v", imgReq)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Printf("%s Request %s", r.Method, r.URL.Path[1:])
	}
}

func main() {
	http.HandleFunc("/", tagHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
