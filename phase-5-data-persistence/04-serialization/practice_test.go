package serialization

import (
	"testing"
)


func TestGobSerialization(t *testing.T) {
	item := CacheItem{Key: "user-10", Value: []byte("active-session")}
	
	b, err := GobSerialize(item)
	if err != nil {
		t.Fatalf("failed to serialize: %v", err)
	}

	var decoded CacheItem
	err = GobDeserialize(b, &decoded)
	if err != nil {
		t.Fatalf("failed to deserialize: %v", err)
	}

	if decoded.Key != item.Key || string(decoded.Value) != string(item.Value) {
		t.Errorf("decoded struct does not match: %+v", decoded)
	}
}
