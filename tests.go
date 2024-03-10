package httptest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

	type TestData struct {
		Request            interface{}
		Route              string
		ExpectedStatusCode int
		ExpectedBody       string
		ExpectedCookie     bool
		Description        string
		Headers map[string]string
	}

	func (test *TestData) Run(t *testing.T, res *http.Response, err error){
		t.Helper()
		
		t.Run(test.Description, func(t *testing.T) {
            t.Log(parseToJSON(test.Request).String())
			req, _ := http.NewRequest(http.MethodPost, test.Route, parseToJSON(test.Request))
			for k, v := range test.Headers{
				req.Header.Add(k, v)
			}
            AssertErrNotNil(t, err)
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
