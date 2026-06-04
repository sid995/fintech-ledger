package version

import "testing"

func TestBanner(t *testing.T) {
	const want = "fintech-ledger 0.0.0-phase0"
	if got := Banner(); got != want {
		t.Fatalf("Banner() = %q, want %q", got, want)
	}
}
