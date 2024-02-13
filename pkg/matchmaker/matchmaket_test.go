package matchmaker

import (
	"os"
	"testing"
)

func TestNewMatchmaker(t *testing.T) {
	os.Setenv("knownInfo", "{\"GoodBook\" : \"this is a story about a pet\", \"BadBook\" : \"this is a story about a man\"}")

	matchmaker := NewMatchmaker()

	val, ok := matchmaker.store["GoodBook"]
	if val != "this is a story about a pet" || !ok {
		t.Fatalf("NewMatchmaker factory fails")
	}
}
