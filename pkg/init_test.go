package gkudu

import "testing"

func TestGetKnownResult(t *testing.T) {
	got := GetKnownResult()
	if got != "known" {
		t.Error("Known result not received, test failed.")
	}
}
