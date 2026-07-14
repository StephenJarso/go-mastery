# Phase 4: I/O, Networking & Web Development

Welcome to **Phase 4** of your Go journey! In this phase, we move from the core language features into the practical tools required to build networked systems, handle file operations, and build high-performance web APIs.

Go was designed at Google to build large-scale, high-concurrency network services. As a result, its standard library for I/O, networking, and HTTP is one of the most powerful and widely used in the entire software ecosystem. 

In this guide, we will explore:
1. **File I/O**: How Go uses simple interfaces (`io.Reader` and `io.Writer`) to handle streams.
2. **Networking (TCP/UDP)**: Writing low-level socket clients and servers.
3. **HTTP & Web Servers**: Building production-grade HTTP servers using standard packages.
4. **Web Routers & Frameworks**: Deciding when to use the standard library vs. routers like Chi, Gin, or Echo.
5. **REST API Design**: Building structured, standard-compliant, type-safe API services.

---

## 🎯 Learning Objectives

By the end of this phase, you will understand:
* **The I/O abstractions**: Why `io.Reader` and `io.Writer` are the most important interfaces in Go.
* **Concurrent Networking**: How to write TCP/UDP servers that handle millions of connections using simple goroutines.
* **Production HTTP**: Writing custom routers, middleware, and HTTP client requests.
* **REST APIs**: Binding JSON bodies to structures, validating inputs, and standardizing error responses.

---

## 📚 Detailed Topic Explanations & Language Comparisons

### 1. File I/O Operations: Stream Abstractions

In many languages, file I/O operations are tightly bound to file classes (e.g., Python's `open()` file object, Java's `File` or `FileInputStream`). 

Go decouples I/O from the concrete medium (disk file, socket, HTTP response) using two simple interfaces:
1. **`io.Reader`**: Anything that can read bytes into a buffer.
2. **`io.Writer`**: Anything that can write bytes from a buffer.

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

#### ⚖️ Why this is a game-changer
Because files, network connections, HTTP request bodies, and in-memory byte buffers all implement `io.Reader` and `io.Writer`, they are completely interchangeable.
* Want to upload a file to an HTTP endpoint? You can stream it directly using `io.Copy(httpResponseWriter, fileReader)` without loading the whole file into RAM.
* Want to test a parser function that reads from a file? Pass it an in-memory `strings.Reader` or `bytes.Buffer` instead, avoiding slow disk operations in tests.

#### 🚰 Streaming Large Files
For a beginner, the natural instinct is to read a whole file into memory (e.g., `os.ReadFile("large.txt")`). However, if the file is 10GB, the application will crash due to out-of-memory errors. 

Go's standard way to stream large files is using a chunk buffer:
```go
buf := make([]byte, 4096) // 4KB buffer
for {
    n, err := file.Read(buf)
    if err == io.EOF {
        break // Reached end of file
    }
    // Process buf[:n]
}
```
Or simply use `io.Copy(writer, reader)` which manages this buffer-streaming loop automatically.

---

### 2. Networking: Low-level Sockets (TCP/UDP)

In traditional languages, concurrent TCP servers require complex, non-blocking I/O multiplexers (like Java's NIO/Netty, Node.js event-loop, or Python's `asyncio`) or spawning heavyweight OS threads.

In Go, because goroutines are so lightweight, you can write high-performance concurrent socket servers using a simple **blocking execution model**:

```go
listener, _ := net.Listen("tcp", ":8080")
for {
    conn, _ := listener.Accept() // Blocks waiting for connection
    go handleConnection(conn)    // Spawns a lightweight goroutine for each client
}
```

#### ⚖️ Networking Comparison

| Feature | Java NIO / Node.js | Go (`net` package) |
| :--- | :--- | :--- |
| **Model** | Async Event Loop / Reactor Pattern. | Blocking calls inside lightweight goroutines. |
| **Code Style** | Callback-heavy or Promise/Future-based. | Clean, sequential, linear code. |
| **Resource Overhead** | High complexity to manage threads / event loops. | Negligible overhead (2KB per goroutine connection). |

---

### 3. HTTP & Web Development: The Power of Standard Library

In Node.js, developers almost always import Express. In Python, they use Flask/Django. In Java, they use Spring Boot.

In Go, the standard library `net/http` is production-ready, highly optimized, and supports HTTP/1.1, HTTP/2, and TLS out of the box. Many massive production apps at companies like Netflix, Cloudflare, and Uber use **only** the standard library.

#### 🔀 HTTP Request Handling
Go uses a Multiplexer (`http.ServeMux`) to route incoming requests to handlers:

```go
mux := http.NewServeMux()
mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
})
http.ListenAndServe(":8080", mux)
```

#### 🛡️ The Middleware Pattern
Middleware is a function that wraps an `http.Handler` to run code before or after the request (e.g., logging, auth checks, recovery).

```go
func LoggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Println("Received request:", r.URL.Path)
        next.ServeHTTP(w, r) // Pass control to the next handler
    })
}
```
*Compare to*: Express middleware `(req, res, next) => { ... next() }` or Spring filters.

---

### 4. Web Routers vs. Frameworks

When your application routes become complex (e.g., nested paths like `/users/:id/posts/:post_id`), the standard `http.ServeMux` (prior to Go 1.22) can feel limited. This is when developers choose libraries or frameworks.

#### 1. **Chi Router** (Lightweight/Idiomatic)
* **What it is**: A fast, lightweight router built on top of `net/http`.
* **Why use it**: 100% compatible with the standard library handlers (`http.HandlerFunc`). It adds URL routing parameters, grouping, and sub-routers without changing Go's standard paradigms.

#### 2. **Gin & Echo** (High-Performance Frameworks)
* **What they are**: Full-fledged HTTP web frameworks.
* **Why use it**:
  * Faster routing (using Radix tree structures).
  * Richer context objects (`gin.Context`, `echo.Context`) which group request, response, and helper methods.
  * Built-in JSON data binding (`context.BindJSON(&myStruct)`).
  * Integrated validation, logger, and recovery middlewares.

---

### 5. REST API Development

Building a clean REST API in Go requires:
1. **Structuring Route Handlers**: Mapping `GET`, `POST`, `PUT`, `DELETE` methods.
2. **JSON Data Binding**: Extracting the JSON body into a Go struct.
3. **Data Validation**: Using library tags (e.g., `validate:"required,email"`) to ensure data is clean.
4. **Structured JSON Responses**: Returning uniform payload wraps for success and error conditions.

```go
type CreateUserRequest struct {
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required,email"`
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    var req CreateUserRequest
    // Decode & Bind JSON
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(map[string]string{"error": "Invalid JSON"})
        return
    }
    // Process ...
}
```

---

## 🗂️ Phase 4 Code Directory Structure

* **[01-file-io/](file:///home/sjarso/go-mastery/phase-4-io-net-web/01-file-io)**: Accessing physical disks.
  * [basics.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/01-file-io/basics.go): Read/Write primitives and operations.
  * [buffered.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/01-file-io/buffered.go): Writing with `bufio.Writer` and read scanner loops.
  * [filepaths.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/01-file-io/filepaths.go): Traversal, path cleaning, directory scans.
  * [streaming.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/01-file-io/streaming.go): Processing large files with static buffer windows.
* **[02-networking/](file:///home/sjarso/go-mastery/phase-4-io-net-web/02-networking)**: Sockets and concurrency.
  * [tcp_server.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/02-networking/tcp_server.go): Accepting concurrent connection streams.
  * [tcp_client.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/02-networking/tcp_client.go): Sending payloads to a TCP address.
  * [udp_server.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/02-networking/udp_server.go) & [udp_client.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/02-networking/udp_client.go): Connectionless UDP packet routing.
* **[03-http-web/](file:///home/sjarso/go-mastery/phase-4-io-net-web/03-http-web)**: Standard HTTP.
  * [server.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/03-http-web/server.go): Setting up custom routers and mapping request paths.
  * [client.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/03-http-web/client.go): Constructing external GET/POST requests.
  * [middleware.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/03-http-web/middleware.go): Creating custom handler interceptors.
* **[04-web-frameworks/](file:///home/sjarso/go-mastery/phase-4-io-net-web/04-web-frameworks)**: Modern routers.
  * [chi_router.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/04-web-frameworks/chi_router.go): Lightweight, standard-compliant routing.
  * [gin_framework.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/04-web-frameworks/gin_framework.go) & [echo_framework.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/04-web-frameworks/echo_framework.go): Rich routing, parameters binding.
* **[05-rest-api/](file:///home/sjarso/go-mastery/phase-4-io-net-web/05-rest-api)**: Building complete APIs.
  * [main.go](file:///home/sjarso/go-mastery/phase-4-io-net-web/05-rest-api/main.go): Designing endpoints with binding and standardized responses.

---

## 🚀 How to Run the Code

Start the examples by launching them in your terminal:

```bash
# Go to the HTTP web directory
cd phase-4-io-net-web/03-http-web

# Run the server example
go run server.go middleware.go
```

Verify implementation and exercises via testing:
```bash
go test -v ./...
```
