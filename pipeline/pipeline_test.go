package pipeline

import (
	"context"
	"testing"
)

func TestExecute(t *testing.T) {
	// just tests that Execute doesn't deadlock
	Execute(context.Background())
}
