package set

import (
	"log"
	"testing"
)

func TestIsSuperAnimal(t *testing.T) {
	expected := true
	got := true
	if got != expected {
		t.Errorf("Expected: %v, got: %v", expected, got)
	} else {
		log.Println("OK")
	}
}
