package httptest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

// TestData are common fields for http request and expected values
type TestData struct {
	Request            interface{}
	Route              string
	ExpectedStatusCode int
	ExpectedBody       string
	ExpectedCookie     bool
	Description        string
	Headers            map[string]string
	ErrNil             bool
	Method             string
}

// SetupRequest creates *http.Request with provided fields in TestData
func (test *TestData) SetupRequest() (*http.Request, error) {
	if test.Method == "" {
		test.Method = http.MethodPost
	}
	req, err := http.NewRequest(test.Method, test.Route, parseToJSON(test.Request))
	for k, v := range test.Headers {
		req.Header.Add(k, v)
	}
	return req, err
}

// Run takes *http.Response and possible error and asserts error, cookie, status and body with TestData expected fields
func (test *TestData) Run(t *testing.T, res *http.Response, err error) {
	t.Helper()

	t.Run(test.Description, func(t *testing.T) {
		t.Log(parseToJSON(test.Request).String())
		AssertErrNil(t, err, test.ErrNil)
		AssertStatus(t, res, test.ExpectedStatusCode)
		AssertCookieExists(t, res, "jwt", test.ExpectedCookie)
		AssertBody(t, res, test.ExpectedBody)
	})
}
func parseToJSON(o interface{}) *bytes.Buffer {
	buffer := bytes.NewBuffer([]byte{})
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(o)
	if err != nil {
		panic(err)
	}
	return buffer
}
