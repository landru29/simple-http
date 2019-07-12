package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type response struct {
	Message string `json:"message,omitempty"`
}

func main() {
	port := 3000
	host := "0.0.0.0"

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")

		bytes, errBody := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		if errBody != nil {
			http.Error(writer, errBody.Error(), http.StatusInternalServerError)
			return
		}

		resp := response{
			Message: string(bytes),
		}

		json.NewEncoder(writer).Encode(resp)
	})

	http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		log.Print("I'm alive")
		writer.Write([]byte("ok"))
	})

	log.Printf("Listening on %s:%d\n", host, port)

	http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
}
