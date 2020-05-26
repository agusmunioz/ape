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

type Response struct {
	Payload    interface{}
	StatusCode int
	Headers    []HTTPHeader
}

//WithHeader adds a header in the Response.
func (response *Response) WithHeader(name string, value string) *Response {

	if response.Headers == nil {
		response.Headers = []HTTPHeader{}
	}

	response.Headers = append(response.Headers, HTTPHeader{Name: name, Value: value})

	return response
}

//Contains the payload of an error
//swagger:model ErrorPayload
type ErrorPayload struct {
	Message string `json:"message"`
}

//HTTPHeader models an http header for being set in an http response.
type HTTPHeader struct {
	Name  string
	Value string
}

//NewInternalServerError builds a model.Response for InternalServerError http response.
func NewInternalServerError(message string) Response {
	return Response{Payload: ErrorPayload{Message: message},
		StatusCode: http.StatusInternalServerError}
}

//NewNotFound builds a model.Response for NotFound http response.
func NewNotFound(message string) Response {
	return Response{Payload: ErrorPayload{Message: message}, StatusCode: http.StatusNotFound}
}

//NewOk builds a model.Response for Ok http response.
func NewOk(payload interface{}) Response {
	return Response{Payload: payload, StatusCode: http.StatusOK}
}

//NewBadRequest builds a model.Response for a BadRequest http response
func NewBadRequest(message string) Response {
	return Response{Payload: ErrorPayload{Message: message}, StatusCode: http.StatusBadRequest}
}
