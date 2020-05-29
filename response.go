package ape

import "net/http"

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

//NewOk builds a Response for Ok http response.
func NewOk(payload interface{}) Response {
	return Response{Payload: payload, StatusCode: http.StatusOK}
}

//NewCreated builds a Response for Created http response
func NewCreated(payload interface{}) Response {
	return Response{Payload: payload, StatusCode: http.StatusCreated}
}

//NewAccepted builds a Response for Accepted http response
func NewAccepted(payload interface{}) Response {
	return Response{Payload: payload, StatusCode: http.StatusAccepted}
}

//NewInternalServerError builds a Response for InternalServerError http response.
func NewInternalServerError(message string) Response {
	return Response{Payload: ErrorPayload{Message: message},
		StatusCode: http.StatusInternalServerError}
}

//NewNotFound builds a Response for NotFound http response.
func NewNotFound(message string) Response {
	return Response{Payload: ErrorPayload{Message: message}, StatusCode: http.StatusNotFound}
}

//NewBadRequest builds a Response for a BadRequest http response
func NewBadRequest(message string) Response {
	return Response{Payload: ErrorPayload{Message: message}, StatusCode: http.StatusBadRequest}
}

//NewConflict builds a Response for a Conflict http response
func NewConflict(message string) Response {
	return Response{Payload: ErrorPayload{Message: message}, StatusCode: http.StatusConflict}
}

//NewForbidden builds a Response for a Forbidden http response
func NewForbidden(message string) Response {
	return Response{Payload: ErrorPayload{Message: message}, StatusCode: http.StatusForbidden}
}
