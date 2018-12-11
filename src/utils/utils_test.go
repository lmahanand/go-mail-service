package utils

import (
	"testing"
)

func TestMessage(t *testing.T) {

	if Message(true, "Test Data")["status"] != true {
		t.Error("Expected status equal true")
	}

	if Message(true, "Test Data")["message"] != "Test Data" {
		t.Error("Expected message equal Test Data")
	}
}
