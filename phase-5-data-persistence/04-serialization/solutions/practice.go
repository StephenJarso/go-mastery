package solutions

import (
	"bytes"
	"encoding/gob"
)


type CacheItem struct {
	Key   string
	Value []byte
}

func GobSerialize(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func GobDeserialize(b []byte, target interface{}) error {
	buf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(buf)
	return dec.Decode(target)
}
