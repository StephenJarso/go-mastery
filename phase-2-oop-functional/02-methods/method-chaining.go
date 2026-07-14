package main

import (
	"fmt"
	"strings"
)

// MethodChainingExamples demonstrates the builder pattern and fluent interfaces
// Method chaining allows elegant, readable API design

// QueryBuilder is an example of a builder using method chaining
type QueryBuilder struct {
	table  string
	where  []string
	select1 []string
	limit  int
	offset int
}

// NewQueryBuilder creates a new query builder
func NewQueryBuilder() *QueryBuilder {
	return &QueryBuilder{}
}

// From sets the table name
// Returns *QueryBuilder to enable chaining
func (q *QueryBuilder) From(table string) *QueryBuilder {
	q.table = table
	return q // Return receiver for chaining
}

// Select specifies columns to select
// Returns *QueryBuilder to enable chaining
func (q *QueryBuilder) Select(columns ...string) *QueryBuilder {
	q.select1 = columns
	return q // Return receiver for chaining
}

// Where adds a WHERE condition
// Returns *QueryBuilder to enable chaining
func (q *QueryBuilder) Where(condition string) *QueryBuilder {
	q.where = append(q.where, condition)
	return q // Return receiver for chaining
}

// Limit sets the LIMIT clause
// Returns *QueryBuilder to enable chaining
func (q *QueryBuilder) Limit(limit int) *QueryBuilder {
	q.limit = limit
	return q // Return receiver for chaining
}

// Offset sets the OFFSET clause
// Returns *QueryBuilder to enable chaining
func (q *QueryBuilder) Offset(offset int) *QueryBuilder {
	q.offset = offset
	return q // Return receiver for chaining
}

// Build constructs the final query string
// This is the terminal operation (doesn't return builder)
func (q *QueryBuilder) Build() string {
	var query strings.Builder

	if len(q.select1) == 0 {
		query.WriteString("SELECT *")
	} else {
		query.WriteString("SELECT ")
		query.WriteString(strings.Join(q.select1, ", "))
	}

	if q.table != "" {
		query.WriteString(" FROM ")
		query.WriteString(q.table)
	}

	if len(q.where) > 0 {
		query.WriteString(" WHERE ")
		query.WriteString(strings.Join(q.where, " AND "))
	}

	if q.limit > 0 {
		query.WriteString(fmt.Sprintf(" LIMIT %d", q.limit))
	}

	if q.offset > 0 {
		query.WriteString(fmt.Sprintf(" OFFSET %d", q.offset))
	}

	return query.String()
}

// BasicMethodChaining demonstrates simple method chaining
func BasicMethodChaining() {
	fmt.Println("=== Basic Method Chaining ===")

	// Without method chaining (verbose)
	q1 := NewQueryBuilder()
	q1.From("users")
	q1.Select("id", "name", "email")
	q1.Where("age > 18")
	q1.Limit(10)
	fmt.Printf("Without chaining: %s\n", q1.Build())

	// With method chaining (elegant)
	q2 := NewQueryBuilder().
		From("users").
		Select("id", "name", "email").
		Where("age > 18").
		Limit(10)
	fmt.Printf("With chaining: %s\n", q2.Build())
}

// ComplexChaining shows more complex chaining scenarios
func ComplexChaining() {
	fmt.Println("\n=== Complex Method Chaining ===")

	query := NewQueryBuilder().
		From("orders").
		Select("id", "customer_id", "total", "created_at").
		Where("status = 'completed'").
		Where("total > 100").
		Where("created_at > '2024-01-01'").
		Limit(20).
		Offset(40)

	fmt.Printf("Complex query: %s\n", query.Build())
}

// HTTPRequestBuilder shows chaining for HTTP request building
type HTTPRequest struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

// NewHTTPRequest creates a new request
func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{
		method:  "GET",
		headers: make(map[string]string),
	}
}

// Method sets the HTTP method
func (r *HTTPRequest) Method(method string) *HTTPRequest {
	r.method = method
	return r
}

// URL sets the request URL
func (r *HTTPRequest) URL(url string) *HTTPRequest {
	r.url = url
	return r
}

// Header adds a header
func (r *HTTPRequest) Header(key, value string) *HTTPRequest {
	r.headers[key] = value
	return r
}

// Body sets the request body
func (r *HTTPRequest) Body(body string) *HTTPRequest {
	r.body = body
	return r
}

// Send is the terminal operation
func (r *HTTPRequest) Send() string {
	var result strings.Builder
	result.WriteString(fmt.Sprintf("%s %s\n", r.method, r.url))
	for k, v := range r.headers {
		result.WriteString(fmt.Sprintf("%s: %s\n", k, v))
	}
	if r.body != "" {
		result.WriteString("\n" + r.body)
	}
	return result.String()
}

// HTTPRequestExample shows HTTP request chaining
func HTTPRequestExample() {
	fmt.Println("\n=== HTTP Request Builder ===")

	request := NewHTTPRequest().
		Method("POST").
		URL("https://api.example.com/users").
		Header("Content-Type", "application/json").
		Header("Authorization", "Bearer token123").
		Body(`{"name": "Alice", "email": "alice@example.com"}`)

	fmt.Println("Built request:")
	fmt.Println(request.Send())
}

// ConfigBuilder shows configuration object builder
type Config struct {
	host    string
	port    int
	timeout int
	retries int
	debug   bool
}

// NewConfig creates a new config with defaults
func NewConfig() *Config {
	return &Config{
		host:    "localhost",
		port:    8080,
		timeout: 30,
		retries: 3,
		debug:   false,
	}
}

// WithHost sets the host
func (c *Config) WithHost(host string) *Config {
	c.host = host
	return c
}

// WithPort sets the port
func (c *Config) WithPort(port int) *Config {
	c.port = port
	return c
}

// WithTimeout sets the timeout
func (c *Config) WithTimeout(timeout int) *Config {
	c.timeout = timeout
	return c
}

// WithRetries sets the number of retries
func (c *Config) WithRetries(retries int) *Config {
	c.retries = retries
	return c
}

// WithDebug enables debug mode
func (c *Config) WithDebug(debug bool) *Config {
	c.debug = debug
	return c
}

// Build returns the final config
func (c *Config) BuildConfig() *Config {
	return c
}

// ConfigExample shows configuration builder pattern
func ConfigExample() {
	fmt.Println("\n=== Configuration Builder ===")

	config := NewConfig().
		WithHost("api.example.com").
		WithPort(443).
		WithTimeout(60).
		WithRetries(5).
		WithDebug(true).
		BuildConfig()

	fmt.Printf("Config: host=%s, port=%d, timeout=%d, retries=%d, debug=%v\n",
		config.host, config.port, config.timeout, config.retries, config.debug)
}

// MethodChainingBenefits explains why method chaining is useful
func MethodChainingBenefits() {
	fmt.Println("\n=== Method Chaining Benefits ===")
	fmt.Println(`
Fluid Interface (Method Chaining) Benefits:

1. Readability
   - Reads like natural language
   - Clear sequence of operations
   - Easy to understand intent

2. Maintainability
   - Easy to add/remove operations
   - Clear structure
   - Self-documenting

3. Flexibility
   - Can build objects in any order
   - Optional steps
   - Defaults for unset values

4. Discoverability
   - IDE autocomplete shows available methods
   - Easy to explore options
   - Clear API

Examples in Go:
- database/sql query builders
- HTTP client builders
- Configuration builders
- ORM query builders
`)
}

// ImplementingMethodChaining shows implementation tips
func ImplementingMethodChaining() {
	fmt.Println("\n=== Implementing Method Chaining ===")
	fmt.Println(`
Rules for Method Chaining:

1. Return *T (pointer to receiver)
   func (b *Builder) Method() *Builder {
       // do something
       return b  // return pointer to enable chaining
   }

2. Methods must be on pointer receiver
   - Not on value receiver
   - Need *T to return *T

3. Terminal operations
   - Last method often doesn't return builder
   - Returns final result instead
   - Example: Build(), Send(), Execute()

4. Order flexibility
   - Design so operations can be in any order
   - Or enforce order with zero-checks
   - Document expected order if it matters

5. Defaults
   - Provide sensible defaults
   - Allow override via methods
   - Make common cases easy
`)
}

func main() {
	BasicMethodChaining()
	ComplexChaining()
	HTTPRequestExample()
	ConfigExample()
	MethodChainingBenefits()
	ImplementingMethodChaining()

	fmt.Println("\n=== Method Chaining Examples Complete ===")
}
