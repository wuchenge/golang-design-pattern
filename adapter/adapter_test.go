package adapter

import "testing"

var expectedResponse = "Specific Request from Adapter"

func TestAdapter(t *testing.T) {
	adaptee := NewAdaptee()
	adapter := NewAdapter(adaptee)
	response := adapter.Request()

	if response != expectedResponse {
		t.Errorf("Expected response '%s', but got '%s'", expectedResponse, response)
	}
}
