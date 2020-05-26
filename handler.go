package ape

import (
	"encoding/json"
	"log"
	"net/http"
)

//Handler is a wrapper that write a model.Response as Json in the http.ResponseWriter
func Handler(handler func(r *http.Request) Response) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := handler(r)
		body, err := json.Marshal(response.Payload)
		if err != nil {
			log.Println("There was an error when marshalling json response: " + err.Error())
			//TODO: return an error response.
		}
		w.Header().Set("Content-Type", "application/json")
		for _, header := range response.Headers {
			w.Header().Set(header.Name, header.Value)
		}
		w.WriteHeader(response.StatusCode)
		_, err = w.Write(body)
		if err != nil {
			log.Println("There was an error when writing json response: " + err.Error())
			//TODO: return an error response.
		}
	})
}