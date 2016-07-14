package handlers

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := "The Dash Button service is running."
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDashHandler(t *testing.T) {
	var tests = []struct {
		Mac  string
		Mode string
		Cmd  string
	}{
		{"74:75:48:C3:B1:D0", "add", "/usr/local/bin/mosquitto_pub -h mosquitto -t home/dash_btn/0 -m ON"},
		{"74:75:48:C3:B1:D1", "del", "/usr/local/bin/mosquitto_pub -h mosquitto -t home/dash_btn/1 -m ON"},
		{"74:75:48:C3:B1:D2", "old", ""},
	}

	testConfig := []byte(`
'74:75:48:C3:B1:D0':
  'add+old': [ '/usr/local/bin/mosquitto_pub', '-h', 'mosquitto', '-t', 'home/dash_btn/0', '-m', 'ON' ]
'74:75:48:C3:B1:D1':
  'del': [ '/usr/local/bin/mosquitto_pub', '-h', 'mosquitto', '-t', 'home/dash_btn/1', '-m', 'ON' ]
'74:75:48:C3:B1:D2':
  'add+del': [ '/usr/local/bin/mosquitto_pub', '-h', 'mosquitto', '-t', 'home/dash_btn/2', '-m', 'ON' ]
`)
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewBuffer(testConfig)); err != nil {
		t.Error(err)
	}

	for _, v := range tests {
		req, err := http.NewRequest("GET", fmt.Sprintf("/dash?mac=%s&mode=%s", v.Mac, v.Mode), nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(DashHandler)

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body is what we expect.
		expected := v.Cmd

		if actual := rr.Body.String(); actual != expected {
			t.Errorf("Unexpected Body - Actual: %v; Expected: %v", actual, expected)
		}
	}
}
