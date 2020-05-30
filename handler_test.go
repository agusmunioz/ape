package ape

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

//Test Handler response write when no extra header is specified in Response.
func TestHandlerNoExtraHeader(t *testing.T) {
	expected := `{
				   "id": "1",
				   "name": "fake"
				 }`

	response := Response{Payload: FakePayload{ID: "1", Name: "fake"}}
	testHandler(t, response, expected, []HTTPHeader{})
}

//Test Handler response write when an extra header is specified in Response.
func TestHandlerWithExtraHeader(t *testing.T) {

	expected := `{
				   "id": "1",
				   "name": "fake"
				}`

	response := Response{Payload: FakePayload{ID: "1", Name: "fake"},
		Headers: []HTTPHeader{HTTPHeader{Name: "Extra", Value: "Some"}}}
	testHandler(t, response, expected, []HTTPHeader{HTTPHeader{Name: "Extra", Value: "Some"}})
}

//Test Handler returns a not nil wrapper function.
func TestNotNil(t *testing.T) {

	//Handler func wrapped by json.Handler.
	mockHandler := func(r *http.Request) Response {
		return Response{}
	}

	wrapperFunc := Handler(mockHandler)

	assert.NotNil(t, wrapperFunc)
}

//TestMarshalError test when marshalling the response returns an error.
func TestMarshalError(t *testing.T) {
	expected := `{"message": "Unexpected Internal Error"}`
	response := Response{Payload: map[string]interface{}{
		"foo": make(chan int),
	}}

	testHandler(t, response, expected, []HTTPHeader{})
}

//Test the Handler using the specified scenario.
func testHandler(t *testing.T, response Response, expectedJSON string, expectedHeaders []HTTPHeader) {

	//Handler func wrapped by Handler.
	mockHandler := func(r *http.Request) Response {
		return response
	}

	req := httptest.NewRequest("GET", "http://rfk.com/foo", nil)
	w := httptest.NewRecorder()

	Handler(mockHandler).ServeHTTP(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	var expectedIndentedBody bytes.Buffer
	err := json.Indent(&expectedIndentedBody, []byte(expectedJSON), "=", "\t")

	if err != nil {
		t.Error(err.Error())
	}

	var indentedBody bytes.Buffer
	err = json.Indent(&indentedBody, body, "=", "\t")

	if err != nil {
		t.Error(err.Error())
	}

	assertion := assert.New(t)
	assertion.Equal(expectedIndentedBody.String(), indentedBody.String(), "Unexpected json body")
	assertion.NotNil(w.Header(), "No headers found in response")
	assertion.Equal(1+len(expectedHeaders), len(w.Header()), "Unexpected amount of headers")
	assertion.Equal("application/json", w.Header().Get("Content-Type"), "Unexpected Content-Type")

	for _, header := range expectedHeaders {
		assertion.Equal(header.Value, w.Header().Get(header.Name), "Unexpected header "+header.Name+" value")
	}
}

type FakePayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
