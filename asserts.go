package httptest

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

//Assert  calls t.Errorf with message and msgArgs as parameters if statement is false
func Assert(t testing.TB, statement bool, message string, msgArgs ...any) {
	if !statement {
		t.Errorf(message, msgArgs...)
	}
}
//AssertErrNotNil fails test with err.Error print if err is not nil
func AssertErrNil(t testing.TB, err error, exists bool) {
	t.Helper()
    if !exists{
	Assert(t, err == nil, "expected error to be nil, got:\n %v", err)
}else{

	Assert(t, err != nil, "expected error to exist, got nil")
}
}
//AssertBody fails test if body of resp  does not contain want
func AssertBody(t testing.TB, resp *http.Response, want string) {
	t.Helper()
	body, err := io.ReadAll(resp.Body)
	AssertErrNil(t, err, false)
	Assert(t, strings.Contains(string(body), want), "expected request body to contain %s, got: \n %s", want, body)
}

//AssertStatus fails test if status code of resp is not equal to want
func AssertStatus(t testing.TB, resp *http.Response, want int) {
	t.Helper()
	s := resp.StatusCode
	Assert(t, want == s, "expected status of %d, got %d", want, s)
}
//AssertCookieExists fails test if cookie with given name does/does not exist depending on exists parameter
func AssertCookieExists(t testing.TB, resp *http.Response, name string, exists bool) {
    t.Helper()
    _, e := GetCookieValue(resp, "jwt")
	if exists {
		Assert(t, e, "Expected %s name cookie to exist, but it doesn't", name)
	} else {
		Assert(t, !e, "Did not  expect %s name cookie to exist, but it does", name)
	}
}

//GetCookieValue returns value of cookie named cookieName and if it exists 
func GetCookieValue(resp *http.Response, cookieName string) (value string, exists bool) {
	cookies := resp.Cookies()
	for _, c := range cookies {
		if c.Name == cookieName {
			value = c.Value
			exists = true
			return
		}
	}
	return
}
