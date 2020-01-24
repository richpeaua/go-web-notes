package main

import (
	"net/http"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	
	// Instantiate test router
	r := newRouter()

	// Create mock server
	mockServer := httptest.NewServer(r)

	// Mock server runs server and exposes its location in ".URL" attribute
	// Concatenate that attribute with "/hello" route on router to create mock req endpoint
	resp, err := http.Get(mockServer.URL + "/hello")
	if err != nil {
		t.Fatal(err)
	}

	// Check if status is OK, throw error if not
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// Read response body
	defer resp.Body.Close()
	// Read body into a bunch of bites (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Convert bites into string
	respString := string(b)
	expected := "hello"

	// Test if "respString" returns "expected" value
	if respString != expected {
		t.Errorf("Response should be %s, but got %s", expected, respString)
	}

}

func TestRouterForNonExistentRoute(t *testing.T) {
	r := newRouter()
	mockServer := httptest.NewServer(r)

	resp, err := http.Post(mockServer.URL + "/hello", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("Status should be 405, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respString := string(b)
	expected := ""

	if respString != expected {
		t.Errorf("Response should be blank, got %s", respString)
	}
}

