package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"reflect"
)

// StructTags demonstrates the use of struct tags for metadata.
// Tags provide information about how struct fields should be processed
// in various contexts (JSON marshaling, validation, ORM, etc.)

// Person demonstrates common struct tags
type Person struct {
	// json tag controls JSON marshaling/unmarshaling
	ID int `json:"id"`

	// Multiple tags can be used together
	FirstName string `json:"firstName" xml:"FirstName"`
	LastName  string `json:"lastName" xml:"LastName"`

	// Tag options: omitempty means exclude if zero value
	Email string `json:"email,omitempty"`

	// Tag option: unexported field won't be marshaled
	Password string `json:"-"` // "-" means ignore this field completely
}

// JSONTagOptions shows different JSON tag options
type Config struct {
	// Rename field in JSON
	Hostname string `json:"host"`

	// Include even if empty (default behavior)
	Port int `json:"port"`

	// Omit if empty/zero value
	Timeout int `json:"timeout,omitempty"`

	// Skip this field in JSON
	DebugKey string `json:"-"`

	// Multiple options separated by comma
	SSLCert string `json:"ssl_cert,omitempty"`
}

// BasicJSONTags demonstrates JSON marshaling and unmarshaling with tags
func BasicJSONTags() {
	fmt.Println("=== Basic JSON Tags ===")

	// Create a person
	person := Person{
		ID:        1,
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alice@example.com",
		Password:  "secret123", // Will be ignored due to json:"-"
	}

	// Marshal to JSON - converts struct to JSON bytes
	jsonData, err := json.MarshalIndent(person, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling: %v\n", err)
		return
	}

	fmt.Printf("Marshaled JSON:\n%s\n", string(jsonData))
	fmt.Println("Note: Password field is not in JSON because of json:\"-\"")
}

// JSONUnmarshaling demonstrates unmarshaling JSON back to struct
func JSONUnmarshaling() {
	fmt.Println("\n=== JSON Unmarshaling ===")

	// JSON string
	jsonStr := `{
		"id": 2,
		"firstName": "Bob",
		"lastName": "Smith",
		"email": "bob@example.com"
	}`

	// Unmarshal JSON to struct
	var person Person
	err := json.Unmarshal([]byte(jsonStr), &person)
	if err != nil {
		fmt.Printf("Error unmarshaling: %v\n", err)
		return
	}

	fmt.Printf("Unmarshaled struct: %+v\n", person)
	fmt.Printf("First Name: %s\n", person.FirstName)
}

// OmitEmptyTag demonstrates the omitempty option
func OmitEmptyTag() {
	fmt.Println("\n=== omitempty Option ===")

	// Config with timeout unset (zero value)
	config1 := Config{
		Hostname: "localhost",
		Port:     8080,
		// Timeout is 0 (not set)
		DebugKey: "debug_secret", // Will be ignored
		SSLCert:  "",             // Empty, will be omitted
	}

	jsonData1, _ := json.MarshalIndent(config1, "", "  ")
	fmt.Printf("Config 1 (with omitempty):\n%s\n", string(jsonData1))
	fmt.Println("Note: 'timeout' and 'ssl_cert' are omitted because they're zero values")

	// Config with timeout set
	config2 := Config{
		Hostname: "example.com",
		Port:     443,
		Timeout:  30,
		SSLCert:  "/path/to/cert.pem",
	}

	jsonData2, _ := json.MarshalIndent(config2, "", "  ")
	fmt.Printf("\nConfig 2 (all fields set):\n%s\n", string(jsonData2))
}

// IgnoreFieldTag demonstrates the "-" tag
func IgnoreFieldTag() {
	fmt.Println("\n=== Ignore Field Tag (\"-\") ===")

	person := Person{
		ID:        3,
		FirstName: "Charlie",
		LastName:  "Brown",
		Email:     "charlie@example.com",
		Password:  "very_secret_password",
	}

	jsonData, _ := json.MarshalIndent(person, "", "  ")
	fmt.Printf("Person JSON:\n%s\n", string(jsonData))
	fmt.Println("\nPassword field is completely ignored (not even present in JSON)")
	fmt.Println("This is useful for sensitive data that shouldn't be serialized")
}

// XMLTags demonstrates XML struct tags
type Article struct {
	ID    int    `xml:"id,attr"`    // attr makes it an XML attribute
	Title string `xml:"title"`      // element
	Body  string `xml:"body"`       // element
	Tags  []string `xml:"tag"`      // repeated element
}

func XMLTags() {
	fmt.Println("\n=== XML Tags ===")

	article := Article{
		ID:    1,
		Title: "Introduction to Go",
		Body:  "Go is a modern programming language...",
		Tags:  []string{"golang", "programming", "tutorial"},
	}

	xmlData, _ := xml.MarshalIndent(article, "", "  ")
	fmt.Printf("Article XML:\n%s\n", string(xmlData))
	fmt.Println("\nNote: 'id' is an XML attribute (id=\"1\")")
	fmt.Println("Tags are repeated elements in XML")
}

// CustomTagOptions shows how to parse custom tags
type Product struct {
	Name  string `validate:"required" json:"name"`
	Price float64 `validate:"min=0" json:"price"`
	Stock int    `validate:"min=0,max=1000" json:"stock"`
}

func ParseCustomTags() {
	fmt.Println("\n=== Custom Tags ===")

	product := Product{Name: "Laptop", Price: 999.99, Stock: 50}

	// Reflect to read custom tags
	t := reflect.TypeOf(product)

	fmt.Println("Field tags:")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		validateTag := field.Tag.Get("validate")
		jsonTag := field.Tag.Get("json")
		fmt.Printf("%s: validate=%q, json=%q\n", field.Name, validateTag, jsonTag)
	}
}

// FieldName shows that tags use backticks
type Example struct {
	// Tags must use backticks
	Field1 string `json:"field1"`
	
	// You can have multiple tags
	Field2 string `json:"field2" xml:"Field2" validate:"required"`
	
	// Spaces don't matter
	Field3 string `json:"field3" xml:"Field3"`
}

func TagSyntax() {
	fmt.Println("\n=== Tag Syntax ===")

	fmt.Println(`
Struct tags use backticks:
  type MyStruct struct {
    Field string `json:"field"`
  }

Multiple tags can be combined:
  Field string `json:"field" xml:"Field" validate:"required"`

Common tag formats:
  json:"fieldName"              - JSON marshaling
  xml:"fieldName"               - XML marshaling
  json:"fieldName,omitempty"    - Omit if empty
  json:"-"                       - Ignore field
  json:"fieldName,string"       - Convert to string
  
Tag options are comma-separated:
  json:"name,omitempty,string"
`)
}

// RealWorldExample combines multiple concepts
type User struct {
	ID        int    `json:"id" xml:"id,attr"`
	Username  string `json:"username" validate:"required,min=3,max=20"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"-" validate:"required,min=8"` // Never serialize password
	CreatedAt string `json:"created_at" xml:"CreatedAt"`
	UpdatedAt string `json:"updated_at,omitempty" xml:"UpdatedAt,omitempty"`
	IsAdmin   bool   `json:"is_admin" validate:"boolean"`
	Status    string `json:"status,omitempty" validate:"oneof=active inactive banned"`
}

func RealWorldExample() {
	fmt.Println("\n=== Real World Example ===")

	user := User{
		ID:        123,
		Username:  "john_doe",
		Email:     "john@example.com",
		Password:  "hashed_password_here",
		CreatedAt: "2024-01-01T10:00:00Z",
		IsAdmin:   false,
		Status:    "active",
	}

	jsonData, _ := json.MarshalIndent(user, "", "  ")
	fmt.Printf("User JSON:\n%s\n", string(jsonData))

	fmt.Println("\nKey observations:")
	fmt.Println("1. Password is not in JSON (json:\"-\")")
	fmt.Println("2. UpdatedAt is omitted because it's empty (json:\"...,omitempty\")")
	fmt.Println("3. Tags also include validation rules for external validators")
}

func main() {
	BasicJSONTags()
	JSONUnmarshaling()
	OmitEmptyTag()
	IgnoreFieldTag()
	XMLTags()
	ParseCustomTags()
	TagSyntax()
	RealWorldExample()

	fmt.Println("\n=== All Tag Examples Completed ===")
}
