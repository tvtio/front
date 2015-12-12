package cache

import "testing"

func TestHash(t *testing.T) {
	hash := Hash("foo")
	if hash == "" {
		t.Error("Hash is null")
	}
}
