package service

import (
	"testing"

	m "../model"
)

func TestVerifyTime(t *testing.T) {
	email := m.Email{From: "abc@gmail.com", ScheduledTime: "09 Dec 18 4:36 UTC"}

	if email.ScheduledTime != "09 Dec 18 4:36 UTC" {
		t.Error("Expected ScheduledTime equal 09 Dec 18 4:36 UTC")
	}

	if email.From != "abc@gmail.com" {
		t.Error("Expected From equal abc@gmail.com")
	}
}
