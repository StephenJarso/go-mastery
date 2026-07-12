package serialization

import (
	"bytes"
	"fmt"
)

// Protocol Buffers (protobuf) are Google's binary serialization format.
// To use them, you follow these steps:
// 1. Define a schema in a ".proto" file:
//
//    syntax = "proto3";
//    package serialization;
//    option go_package = "./serialization";
//
//    message UserInfo {
//      int32 id = 1;
//      string name = 2;
//      string email = 3;
//    }
//
// 2. Install the protobuf compiler (protoc) and the Go plugin (protoc-gen-go).
// 3. Compile the schema:
//    $ protoc --go_out=. --go_opt=paths=source_relative user.proto
// 4. Use the generated Go structs with the official "google.golang.org/protobuf/proto" library.

// Under the hood, Protobuf achieves its high speed and small payload size by:
// 1. Omitting field names on the wire: It uses field numbers (e.g., 1, 2) and wire types.
// 2. Varints: Variable-length integers, meaning small integers (like 1 or 5) only use 1 byte on the wire
//    instead of 4 or 8 bytes.

// MockUserInfo simulates a compiled Protobuf struct to demonstrate struct tags
// and how Go code interacts with Protobuf properties.
type MockUserInfo struct {
	// Standard generated tags specify: wire type, field tag number, option parameters, and schema name.
	ID    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

// EncodeVarint demonstrates how Varint encoding is done under the hood in Protobuf.
// A Varint uses the most significant bit (MSB) of each byte as a signpost indicating
// if there are more bytes to follow.
func EncodeVarint(value uint64) []byte {
	var buf []byte
	for value >= 0x80 {
		buf = append(buf, byte(value|0x80))
		value >>= 7
	}
	buf = append(buf, byte(value))
	return buf
}

// DecodeVarint decodes a Varint byte stream back to a uint64.
func DecodeVarint(r *bytes.Reader) (uint64, error) {
	var x uint64
	var s uint
	for i := 0; ; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return 0, err
		}
		if b < 0x80 {
			if i > 9 || i == 9 && b > 1 {
				return 0, fmt.Errorf("varint overflow")
			}
			return x | uint64(b)<<s, nil
		}
		x |= uint64(b&0x7f) << s
		s += 7
	}
}

// SimpleWireMarshal mock-marshals our UserInfo using a simplified wire format:
// [Field Number << 3 | Wire Type] [Length of Data (for strings)] [Value Bytes]
// - Wire Type 0: Varint (for int32)
// - Wire Type 2: Length-delimited (for strings)
func SimpleWireMarshal(u MockUserInfo) []byte {
	var buf bytes.Buffer

	// 1. Write ID (Field 1, Wire Type 0)
	// Key: (field_number << 3) | wire_type = (1 << 3) | 0 = 8
	buf.WriteByte(8)
	buf.Write(EncodeVarint(uint64(u.ID)))

	// 2. Write Name (Field 2, Wire Type 2)
	// Key: (2 << 3) | 2 = 18
	buf.WriteByte(18)
	nameBytes := []byte(u.Name)
	buf.Write(EncodeVarint(uint64(len(nameBytes))))
	buf.Write(nameBytes)

	// 3. Write Email (Field 3, Wire Type 2)
	// Key: (3 << 3) | 2 = 26
	buf.WriteByte(26)
	emailBytes := []byte(u.Email)
	buf.Write(EncodeVarint(uint64(len(emailBytes))))
	buf.Write(emailBytes)

	return buf.Bytes()
}

// SimpleWireUnmarshal mock-deserializes raw bytes back into MockUserInfo.
func SimpleWireUnmarshal(data []byte) (MockUserInfo, error) {
	var u MockUserInfo
	reader := bytes.NewReader(data)

	for {
		keyByte, err := reader.ReadByte()
		if err != nil {
			break // EOF
		}

		fieldNumber := keyByte >> 3
		wireType := keyByte & 0x07

		switch fieldNumber {
		case 1: // ID (Varint)
			if wireType != 0 {
				return u, fmt.Errorf("unexpected wire type %d for field 1", wireType)
			}
			val, err := DecodeVarint(reader)
			if err != nil {
				return u, err
			}
			u.ID = int32(val)
		case 2: // Name (Length-delimited)
			if wireType != 2 {
				return u, fmt.Errorf("unexpected wire type %d for field 2", wireType)
			}
			length, err := DecodeVarint(reader)
			if err != nil {
				return u, err
			}
			nameBuf := make([]byte, length)
			_, err = reader.Read(nameBuf)
			if err != nil {
				return u, err
			}
			u.Name = string(nameBuf)
		case 3: // Email (Length-delimited)
			if wireType != 2 {
				return u, fmt.Errorf("unexpected wire type %d for field 3", wireType)
			}
			length, err := DecodeVarint(reader)
			if err != nil {
				return u, err
			}
			emailBuf := make([]byte, length)
			_, err = reader.Read(emailBuf)
			if err != nil {
				return u, err
			}
			u.Email = string(emailBuf)
		default:
			// Unknown field, skip it based on wire type (standard protobuf behavior)
			if wireType == 0 {
				_, _ = DecodeVarint(reader)
			} else if wireType == 2 {
				length, _ := DecodeVarint(reader)
				_, _ = reader.Read(make([]byte, length))
			} else {
				return u, fmt.Errorf("unsupported wire type %d for unknown field %d", wireType, fieldNumber)
			}
		}
	}

	return u, nil
}
