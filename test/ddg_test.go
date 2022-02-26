package goduckgo_test

import (
	"testing"

	"github.com/danielwetan/goddg"
)

func TestIsResultPresent(t *testing.T) {
	query, _ := goddg.Query("jakarta")

	if query == nil {
		t.Errorf("result not found")
	}
}
