package matchmaker

import (
	"dirs/pkg/logger"
	"os"
	"testing"
)

func TestNewMatchmaker(t *testing.T) {
	os.Setenv("knownInfo", "{\"GoodBook\" : \"this is a story about a pet\", \"BadBook\" : \"this is a story about a man\"}")
	i, _, e := logger.NullLogger()

	matchmaker := NewMatchmaker(i, e).(actualMatchmaker)

	val, ok := matchmaker.store["GoodBook"]
	if val != "this is a story about a pet" || !ok {
		t.Fatalf("NewMatchmaker factory fails")
	}
}
