package matchmaker

import (
	"dirs/pkg/logger"
	"os"
	"testing"
)

func TestNewMatchmaker(t *testing.T) {
	os.Setenv("knownInfo", "{\"GoodBook\" : \"this is a story about a pet\", \"BadBook\" : \"this is a story about a man\"}")
	_, _, Error := logger.NullLogger()

	matchmaker := NewMatchmaker(Error).(actualMatchmaker)

	val, ok := matchmaker.store["GoodBook"]
	if val != "this is a story about a pet" || !ok {
		t.Fatalf("NewMatchmaker factory fails")
	}
}
