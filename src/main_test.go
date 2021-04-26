package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSayHiHandler(t *testing.T){
	tt := []struct{
		method string
		endpoint string
		expected string
	}{
		{method: "GET", endpoint: "localhost:8080/", expected: "Hi, im glad you call\n" },
		{method: "POST", endpoint: "localhost:8080/", expected: "Hi, im glad you call\n" },
		{method: "PUT", endpoint: "localhost:8080/", expected: "Hi, im glad you call\n" },
		{method: "DELETE", endpoint: "localhost:8080/", expected: "Hi, im glad you call\n" },

	}
	for _ , tc:= range tt{

		req, err:= http.NewRequest(tc.method, tc.endpoint, nil)
		
		if err!= nil{
			t.Fatalf("Could not create any request: %v", err)
		}
		
		rec := httptest.NewRecorder()
		
		sayHiHandler(rec, req)
		
		res := rec.Result()
		
		defer res.Body.Close()
		
		b, err := ioutil.ReadAll(res.Body)
		
		if err!= nil{
			t.Fatalf("Could not read body from response: %v", err)
		}

		if res.StatusCode != http.StatusOK{
			t.Fatalf("Did not get status 200")
		}
		
		msg:= string(bytes.TrimSpace(b))

		expectedMsg :=  strings.TrimSpace(tc.expected)

		if msg != expectedMsg {
			t.Error("Recived and expected body do not match.\nRecived: ", msg, "\nExpected", expectedMsg )
		}
		
	}


}

func TestServerAndRouting(t *testing.T){
	server := httptest.NewServer(initHandlers())
	defer server.Close()

	res, err:= http.Get(fmt.Sprintf("%v/", server.URL))

	if err!= nil{
		t.Fatal("There was an error while getting a response from the server:", err)
	}

	if res.StatusCode != http.StatusOK{
		t.Error("Expected status 200 but got", res.StatusCode)
	}

	b, err:= ioutil.ReadAll(res.Body)
	
	defer res.Body.Close()

	if err!= nil{
		t.Error("There was an error trying to read response's body:", err)
	}

	if msg:= string(bytes.TrimSpace(b)); msg != strings.TrimSpace("Hi, im glad you call\n"){
		t.Error("Expected response does not match with recived one")
	}
		
	
}