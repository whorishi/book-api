package main

import (
	"bytes"
	"net/http"
	"testing"
	"time"
	"gofr.dev/pkg/gofr/request"
)

func TestIntegration(t *testing.T) {
	go main()
	time.Sleep(3* time.Second)

	tests := []struct {
		desc string
		method     string
		endpoint   string
		statusCode int
		body       []byte
	}{
		{"get books",http.MethodGet,"books",http.StatusOK,nil},
		{"post books",http.MethodPost,"books",http.StatusCreated,[]byte(`{
				"title": "hiiiiii",
				"author": "hellllo",
				"publisher": "holla",
				"price": 123,
				"category": "sayonara"
		}`),
	}}

	for i,tc := range tests {
		req,_ := request.NewMock(tc.method,"http://localhost:8000/"+tc.endpoint,bytes.NewBuffer(tc.body))
		c:=http.Client{}

		resp, err := c.Do(req)
		if err!= nil {
			t.Errorf("TEST[%v] Failed.\tHTTP request encountered Err: %v\n%s", i, err, tc.desc)
			continue
		}
		
		if resp.StatusCode!=tc.statusCode {
			t.Errorf("TEST[%v] Failed.\tExpected %v\tGot %v\n%s", i, tc.statusCode, resp.StatusCode, tc.desc)
		}

		_ = resp.Body.Close()
	}

}


