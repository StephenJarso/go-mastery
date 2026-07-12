package serialization

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestXMLExample(t *testing.T) {
	p := Person{
		ID:    101,
		Name:  "StephenJarso",
		Email: "stephen@example.com",
		City:  "Nairobi",
	}

	// Marshal XML
	xmlBytes, err := MarshalPersonXML(p)
	if err != nil {
		t.Fatalf("MarshalPersonXML failed: %v", err)
	}

	xmlStr := string(xmlBytes)
	if !strings.HasPrefix(xmlStr, `<?xml version="1.0" encoding="UTF-8"?>`) {
		t.Error("expected XML output to contain the standard XML header")
	}

	if !strings.Contains(xmlStr, `id="101"`) || !strings.Contains(xmlStr, `<city>Nairobi</city>`) {
		t.Errorf("unexpected XML content:\n%s", xmlStr)
	}

	// Unmarshal XML
	p2, err := UnmarshalPersonXML(xmlBytes)
	if err != nil {
		t.Fatalf("UnmarshalPersonXML failed: %v", err)
	}

	if p2.ID != p.ID || p2.Name != p.Name || p2.City != p.City {
		t.Errorf("unmarshaled values mismatch: got %+v, expected %+v", p2, p)
	}
}

func TestMsgpackExample(t *testing.T) {
	entry := LogEntry{
		Level:     "error",
		Message:   "database connection closed",
		Timestamp: time.Date(2026, 7, 13, 1, 0, 0, 0, time.UTC),
		Host:      "server-01",
	}

	// Marshal MessagePack
	msgpackBytes, err := MarshalLogEntry(entry)
	if err != nil {
		t.Fatalf("MarshalLogEntry failed: %v", err)
	}

	// Unmarshal MessagePack
	entry2, err := UnmarshalLogEntry(msgpackBytes)
	if err != nil {
		t.Fatalf("UnmarshalLogEntry failed: %v", err)
	}

	if entry2.Level != entry.Level || entry2.Message != entry.Message || entry2.Host != entry.Host {
		t.Errorf("unmarshaled log entry mismatch: got %+v, expected %+v", entry2, entry)
	}
	if !entry2.Timestamp.Equal(entry.Timestamp) {
		t.Errorf("unmarshaled timestamp mismatch: got %v, expected %v", entry2.Timestamp, entry.Timestamp)
	}
}

func TestProtobufVarintEncoding(t *testing.T) {
	// Test a range of integers to verify MSB varint encoding behavior
	testCases := []uint64{0, 1, 127, 128, 300, 16384, 9007199254740991}

	for _, tc := range testCases {
		encoded := EncodeVarint(tc)
		reader := bytes.NewReader(encoded)
		decoded, err := DecodeVarint(reader)
		if err != nil {
			t.Fatalf("failed to decode varint for %d: %v", tc, err)
		}
		if decoded != tc {
			t.Errorf("expected varint %d, got %d", tc, decoded)
		}
	}
}

func TestMockProtobufWireFormat(t *testing.T) {
	u := MockUserInfo{
		ID:    42,
		Name:  "StephenJarso",
		Email: "stephen@example.com",
	}

	// Marshal via simple wire format
	rawBytes := SimpleWireMarshal(u)

	// Unmarshal
	u2, err := SimpleWireUnmarshal(rawBytes)
	if err != nil {
		t.Fatalf("SimpleWireUnmarshal failed: %v", err)
	}

	if u2.ID != u.ID || u2.Name != u.Name || u2.Email != u.Email {
		t.Errorf("unmarshaled mock protobuf mismatch: got %+v, expected %+v", u2, u)
	}
}
