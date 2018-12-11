package utils

import (
	"testing"
)

func TestMessage(t *testing.T) {
	res := Message(true, "Test Data")

	if res["status"] != true && res["message"] == "Test Data" {
		t.Error("Expected status equal true")
	}
}
