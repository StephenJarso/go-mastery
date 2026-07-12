package serialization

import (
	"encoding/xml"
	"fmt"

	"github.com/vmihailenco/msgpack/v5"
)

// PRACTICE EXERCISE #1: XML Config Exporter
// Design a configuration struct that marshals to XML with the following structure:
// <config port="8080">
//   <host>localhost</host>
//   <features>
//     <feature name="metrics" active="true"></feature>
//     <feature name="auth" active="false"></feature>
//   </features>
// </config>

type XMLConfig struct {
	XMLName  xml.Name     `xml:"config"`
	Port     int          `xml:"port,attr"`
	Host     string       `xml:"host"`
	Features []XMLFeature `xml:"features>feature"` // Double nested path mapping
}

type XMLFeature struct {
	Name   string `xml:"name,attr"`
	Active bool   `xml:"active,attr"`
}

func ExportConfigToXML(cfg XMLConfig) ([]byte, error) {
	data, err := xml.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to export XML config: %w", err)
	}
	return append([]byte(xml.Header), data...), nil
}

// PRACTICE EXERCISE #2: MessagePack Cache Cacher
// MessagePack is excellent for cache payloads.
// Implement a CacheItem structure and functions to serialize/deserialize it using msgpack.
//
// Struct requirements:
//  - key (string)
//  - value (byte slice)
//  - ttl_seconds (int64)

type CacheItem struct {
	Key        string `msgpack:"key"`
	Value      []byte `msgpack:"value"`
	TTLSeconds int64  `msgpack:"ttl_seconds"`
}

func SerializeCacheItem(item CacheItem) ([]byte, error) {
	data, err := msgpack.Marshal(item)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize cache item: %w", err)
	}
	return data, nil
}

func DeserializeCacheItem(data []byte) (CacheItem, error) {
	var item CacheItem
	err := msgpack.Unmarshal(data, &item)
	if err != nil {
		return CacheItem{}, fmt.Errorf("failed to deserialize cache item: %w", err)
	}
	return item, nil
}
