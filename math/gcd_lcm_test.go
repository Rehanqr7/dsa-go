package math

import "testing"

func TestGCDLCM(t *testing.T) {
	if got := GCD(54, 24); got != 6 { t.Fatalf("GCD expected 6 got %d", got) }
	if got := LCM(4, 6); got != 12 { t.Fatalf("LCM expected 12 got %d", got) }
}
