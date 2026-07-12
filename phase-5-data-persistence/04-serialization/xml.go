package serialization

import (
	"encoding/xml"
	"fmt"
)

// XML (Extensible Markup Language) is commonly used in legacy systems, enterprise integrations,
// and configurations. Go's standard library provides the "encoding/xml" package,
// which is syntactically very similar to the "encoding/json" package.

// Person represents a data structure mapped to XML.
// XML tags specify:
// - xml:"person": Maps the struct to the XML root element <person>.
// - xml:"id,attr": Maps the field to an XML attribute id="value".
// - xml:"name": Maps to a standard child element <name>value</name>.
// - xml:"address>city": Maps to a nested child element structure <address><city>value</city></address>.
type Person struct {
	XMLName xml.Name `xml:"person"`
	ID      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email,omitempty"`
	City    string   `xml:"address>city"` // Nested tag matching
}

// MarshalPersonXML serializes a Person struct to XML format.
func MarshalPersonXML(p Person) ([]byte, error) {
	// xml.MarshalIndent pretty-prints the output with indent prefixes.
	data, err := xml.MarshalIndent(p, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal XML: %w", err)
	}

	// It is best practice to prepend the standard XML header declaration.
	header := []byte(xml.Header)
	fullXML := append(header, data...)
	return fullXML, nil
}

// UnmarshalPersonXML deserializes XML bytes into a Person struct.
func UnmarshalPersonXML(data []byte) (Person, error) {
	var p Person
	err := xml.Unmarshal(data, &p)
	if err != nil {
		return Person{}, fmt.Errorf("failed to unmarshal XML: %w", err)
	}
	return p, nil
}
