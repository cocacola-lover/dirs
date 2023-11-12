package economy

import (
	"os"
	"strings"
	"testing"
)

func TestNewSimpleEconomy(t *testing.T) {
	os.Setenv("friends", "[\"http://localhost:3000\", \"http://localhost:3002\"]")

	ecWithCorrectEnv := NewSimpleEconomy()

	if len(ecWithCorrectEnv.friends) != 2 {
		t.Errorf("ERROR: TestNewSimpleEconomy. Expected [\"http://localhost:3000\", \"http://localhost:3002\"] but got %s", turnMapToString(ecWithCorrectEnv.friends))
	}

	os.Setenv("friends", "")

	ecEmptyEnv := NewSimpleEconomy()

	if len(ecEmptyEnv.friends) != 0 {
		t.Errorf("ERROR: TestNewSimpleEconomy. Expected [] but got %s", turnMapToString(ecWithCorrectEnv.friends))
	}
}

func turnMapToString(set map[string]bool) string {
	var strArray []string

	for key := range set {
		strArray = append(strArray, key)
	}

	return strings.Join([]string{"[", strings.Join(strArray, ","), "]"}, "")
}
