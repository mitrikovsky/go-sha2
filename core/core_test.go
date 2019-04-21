package core

import "testing"

func TestCalculateHash(t *testing.T) {
	payload := "payload"
	hash := calculateHash(payload, 3)

	want := "a9f838f8c8311f03525e7879b7aeba2adebe97f0633e36b44fda50d873e0e2b8"
	if hash != want {
		t.Errorf("Wrong hash for '%s': got %s, want %s", payload, hash, want)
	}
}
