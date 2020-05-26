package ape

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Test json.Handler response write when no extra header is specified in Response.
func TestHandlerNoExtraHeader(t *testing.T) {
	expected := `{
				   "id": "1",
				   "name": "fake"
				 }`

	response := Response{Payload: FakePayload{ID: "1", Name: "fake"}}
	testHandler(t, response, expected, []HTTPHeader{})
}

//Test json.Handler response write when an extra header is specified in Response.
func TestHandlerWithExtraHeader(t *testing.T) {

	expected := `{
				   "id": "1",
				   "name": "fake"
				}`

	response := Response{Payload: FakePayload{ID: "1", Name: "fake"},
		Headers: []HTTPHeader{HTTPHeader{Name: "Extra", Value: "Some"}}}
	testHandler(t, response, expected, []HTTPHeader{HTTPHeader{Name: "Extra", Value: "Some"}})
}

//Test json.Handler returns a not nil wrapper function.
func TestNotNil(t *testing.T) {

	//Handler func wrapped by json.Handler.
	mockHandler := func(r *http.Request) Response {
		return Response{}
	}

	wrapperFunc := Handler(mockHandler)

	assert.NotNil(t, wrapperFunc)
}

//Test the json.Handler using the specified scenario.
func testHandler(t *testing.T, response Response, expectedJSON string, expectedHeaders []HTTPHeader) {

	//Handler func wrapped by json.Handler.
	mockHandler := func(r *http.Request) Response {
		return response
	}

	req := httptest.NewRequest("GET", "http://rfk.com/foo", nil)
	w := httptest.NewRecorder()

	Handler(mockHandler).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var expectedIndentedBody bytes.Buffer
	json.Indent(&expectedIndentedBody, []byte(expectedJSON), "=", "\t")

	var indentedBody bytes.Buffer
	json.Indent(&indentedBody, body, "=", "\t")

	assert := assert.New(t)
	assert.Equal(expectedIndentedBody.String(), indentedBody.String(), "Unexpected json body")
	assert.NotNil(w.Header(), "No headers found in response")
	assert.Equal(1+len(expectedHeaders), len(w.Header()), "Unexpected amount of headers")
	assert.Equal("application/json", w.Header().Get("Content-Type"), "Unexpected Content-Type")

	for _, header := range expectedHeaders {
		assert.Equal(header.Value, w.Header().Get(header.Name), "Unexpected header "+header.Name+" value")
	}
}

type FakePayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}