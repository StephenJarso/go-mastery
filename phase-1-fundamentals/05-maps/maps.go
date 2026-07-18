package mapsdemo

import (
	"fmt"
)

// DemonstrateMaps shows how maps are declared, initialized, and used in Go.
func DemonstrateMaps() {
	// 1. Declaration and Initialization using make
	// A map maps keys to values. Zero value of a map is nil.
	// A nil map has no keys, nor can keys be added.
	var nilMap map[string]int
	fmt.Printf("Nil map: %v, len: %d\n", nilMap == nil, len(nilMap))

	// To create a map, use make:
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 82
	fmt.Printf("Scores after make: %v, len: %d\n", scores, len(scores))

	// 2. Map Literal
	// You can also initialize a map with values using a literal:
	lookup := map[string]string{
		"US": "United States",
		"CA": "Canada",
		"FR": "France",
	}
	fmt.Printf("Lookup map: %v\n", lookup)

	// 3. Retrieving values and the comma-ok idiom
	// If a key exists, ok is true. If not, val is the zero value and ok is false.
	country, ok := lookup["CA"]
	fmt.Printf("Key 'CA': val = %q, ok = %t\n", country, ok)

	missing, ok := lookup["JP"]
	fmt.Printf("Key 'JP' (missing): val = %q, ok = %t\n", missing, ok)

	// 4. Deleting keys
	// The built-in delete function removes an element from a map.
	delete(lookup, "FR")
	fmt.Printf("After deleting 'FR': %v\n", lookup)

	// 5. Iterating over maps
	// Iteration order over maps is not specified and is not guaranteed to be the same from one iteration to the next.
	fmt.Println("Iterating over scores:")
	for name, score := range scores {
		fmt.Printf("- %s: %d\n", name, score)
	}
}
