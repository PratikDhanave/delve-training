package lib

import "testing"

func TestAdd(t *testing.T) {
	if Add(1, 2) != 3 {
		t.Error("Expected 3, got ", Add(1, 2))
	}
}
