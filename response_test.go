package ape

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

//Test the creation of an InternalServerError response.
func TestNewInternalServerError(t *testing.T) {

	expectedMessage := "A message"
	response := NewInternalServerError(expectedMessage)
	assertErrorResponse(t, response, http.StatusInternalServerError, expectedMessage)

}

//Test the creation of an NotFound response.
func TestNewNotFound(t *testing.T) {

	expectedMessage := "A message"
	response := NewNotFound(expectedMessage)
	assertErrorResponse(t, response, http.StatusNotFound, expectedMessage)

}

//Test the creation of a BadRequest response.
func TestNewBadRequest(t *testing.T) {

	expectedMessage := "A message"
	response := NewBadRequest(expectedMessage)
	assertErrorResponse(t, response, http.StatusBadRequest, expectedMessage)

}

//Test the creation of a BadRequest response.
func TestNewOk(t *testing.T) {

	var expectedPayload interface{} = 7
	response := NewOk(expectedPayload)
	assertResponse(t, response, http.StatusOK, expectedPayload)

}

//Test the adding of a header when response has no previous header which initializes the list of header.ñ
func TestWithHeaderWithNoHeaderListInResponse(t *testing.T) {

	response := Response{}
	response.WithHeader("A header", "A header value")

	assertion := assert.New(t)
	assertion.NotNil(response.Headers, "Unexpected nil list of headers")
	assertion.Equal(1, len(response.Headers), "Unexpected amount of headers")

	header := response.Headers[0]
	assertion.Equal("A header", header.Name, "Unexpected header name")
	assertion.Equal("A header value", header.Value, "Unexpected header value")
}

//Test the adding of a header when the response already has a header in its list.
func TestWithHeaderWithHeaderInResponse(t *testing.T) {

	response := Response{Headers: []HTTPHeader{HTTPHeader{Name: "First header", Value: "First value"}}}

	response.WithHeader("A header", "A header value")

	assertion := assert.New(t)
	assertion.NotNil(response.Headers, "Unexpected nil list of headers")
	assertion.Equal(2, len(response.Headers), "Unexpected amount of headers")

	header := response.Headers[0]
	assertion.Equal("First header", header.Name, "Unexpected header name for existing header")
	assertion.Equal("First value", header.Value, "Unexpected header value for existing header")

	header = response.Headers[1]
	assertion.Equal("A header", header.Name, "Unexpected header name")
	assertion.Equal("A header value", header.Value, "Unexpected header value")
}

//Asserts the error response
func assertErrorResponse(t *testing.T, response Response, expectedHTTPStatus int, expectedMessage string) {

	expectedPayload := ErrorPayload{Message: expectedMessage}
	assertResponse(t, response, expectedHTTPStatus, expectedPayload)
}

func assertResponse(t *testing.T, response Response, expectedHTTPStatus int, expectedPayload interface{}) {
	assertion := assert.New(t)
	assertion.NotNil(response, "Unexpected nil response")
	assertion.Equal(expectedHTTPStatus, response.StatusCode, "Unexpected HTTP status code")
	assertion.Equal(expectedPayload, response.Payload, "Unexpected response payload")
}
