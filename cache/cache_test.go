package cache

import "testing"

func TestDummy(t *testing.T) {
}

func TestHash(t *testing.T) {
	hash := Hash("foo")
	if hash == "" {
		t.Error("Hash is null")
	}
}
