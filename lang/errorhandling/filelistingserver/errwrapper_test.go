package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func errPanic(http.ResponseWriter, *http.Request) error {
	panic(123)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(http.ResponseWriter, *http.Request) error {
	return testingUserError("user error")
}

func errNotFound(http.ResponseWriter, *http.Request) error {
	return os.ErrNotExist
}

func errNoPermission(http.ResponseWriter, *http.Request) error {
	return os.ErrPermission
}

func noErr(w http.ResponseWriter, r *http.Request) error {
	fmt.Fprintln(w, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNotFound, 404, "Not Found"},
	{errNoPermission, 403, "Forbidden"},
	{noErr, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http:localhost:8888", nil)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errorWrapper(tt.h)
		// 启动一个服务器
		server := httptest.NewServer(http.HandlerFunc(f))
		// 发送真正的请求
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(resp *http.Response, expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); got(%d, %s)", expectedCode, expectedMsg, resp.StatusCode, body)
	}
}
