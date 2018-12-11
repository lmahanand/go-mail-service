package utils

import (
	"testing"
)

func TestMessage(t *testing.T) {

	if Message(true, "Test Data")["status"] != true && Message(true, "Test Data")["message"] == "Test Data" {
		t.Error("Expected status equal true")
	}
}
