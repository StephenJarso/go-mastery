package serialization

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
)

func TestExportConfigToXML(t *testing.T) {
	cfg := XMLConfig{
		Port: 9090,
		Host: "127.0.0.1",
		Features: []XMLFeature{
			{Name: "metrics", Active: true},
			{Name: "auth", Active: false},
		},
	}

	xmlData, err := ExportConfigToXML(cfg)
	if err != nil {
		t.Fatalf("ExportConfigToXML failed: %v", err)
	}

	// Unmarshal back to check equality
	var cfg2 XMLConfig
	err = xml.Unmarshal(xmlData, &cfg2)
	if err != nil {
		t.Fatalf("failed to unmarshal exported XML: %v", err)
	}

	if cfg2.Port != cfg.Port || cfg2.Host != cfg.Host {
		t.Errorf("unmarshaled config mismatch: got %+v, expected %+v", cfg2, cfg)
	}

	if len(cfg2.Features) != 2 || cfg2.Features[0].Name != "metrics" || !cfg2.Features[0].Active {
		t.Errorf("unmarshaled features mismatch: got %+v", cfg2.Features)
	}
}

func TestCacheItemSerialization(t *testing.T) {
	item := CacheItem{
		Key:        "user-session-123",
		Value:      []byte("session-secret-payload"),
		TTLSeconds: 3600,
	}

	msgpackBytes, err := SerializeCacheItem(item)
	if err != nil {
		t.Fatalf("SerializeCacheItem failed: %v", err)
	}

	item2, err := DeserializeCacheItem(msgpackBytes)
	if err != nil {
		t.Fatalf("DeserializeCacheItem failed: %v", err)
	}

	if item2.Key != item.Key || item2.TTLSeconds != item.TTLSeconds || !bytes.Equal(item2.Value, item.Value) {
		t.Errorf("deserialized cache item mismatch: got %+v, expected %+v", item2, item)
	}
}
