package ape

import (
	"encoding/json"
	"log"
	"net/http"
)

//Handler is a wrapper that encapsulates JSON management for HTTP responses.
func Handler(handler func(r *http.Request) Response) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := handler(r)
		body, err := json.Marshal(response.Payload)
		statusCode := response.StatusCode

		if err != nil {
			log.Println("There was an error when marshalling JSON response: " + err.Error())
			body = []byte(`{"message": "Unexpected Internal Error"}`)
			statusCode = http.StatusInternalServerError
		}

		w.WriteHeader(statusCode)
		w.Header().Set("Content-Type", "application/json")
		for _, header := range response.Headers {
			w.Header().Set(header.Name, header.Value)
		}

		_, err = w.Write(body)
		if err != nil {
			log.Println("There was an error when writing JSON response: " + err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		}
	})
}
