package peerManager

import (
	"context"
	"testing"
)

func TestPeerManager(t *testing.T) {
	ctx := context.Background()
	pm, err := NewManager(ctx, 5, 6, nil)
	if err != nil {
		t.Fatalf("error creating peer manager: %s", err)
	}
	pm.SetSelfID("self")
}
