# Phase 4: I/O, Networking & Web Development

**Estimated Duration:** 4-5 weeks  
**Prerequisites:** Phase 3 (Advanced Language Features)  
**Focus:** File Systems, Network Protocols, HTTP Servers/Clients, Web Frameworks, and REST API Development.

---

## 🎯 Learning Objectives

After completing this phase, you'll be able to:

- [ ] Read, write, stream, and manipulate files and directories efficiently using the standard library.
- [ ] Understand and implement low-level TCP and UDP network servers and clients.
- [ ] Build robust, production-grade HTTP servers using standard `net/http` and custom routers.
- [ ] Construct and send HTTP requests using `http.Client` with custom configurations (timeouts, headers).
- [ ] Utilize popular Web Frameworks (Gin, Echo, Chi) and compare their trade-offs.
- [ ] Design and implement a standard RESTful API featuring request validation, JSON processing, error wrapping, and structured response routing.

---

## 📚 Topics Covered

### 1. File I/O Operations (`01-file-io`)
Go's file I/O operations are built on top of `io.Reader` and `io.Writer` interfaces, facilitating interoperability between different data sources.
- **File Access**: Reading and writing files using `os` package primitives (`os.Open`, `os.Create`, `os.OpenFile`).
- **Buffered I/O**: Improving performance for small, frequent reads/writes with the `bufio` package.
- **FilePath & Directory**: Navigating directory trees, creating/deleting directories, and path normalization with `path/filepath`.
- **Large File Streaming**: Processing very large files without exceeding memory limits using buffer-based streaming and `io.Copy`.

### 2. Networking Basics (`02-networking`)
Go offers first-class socket programming utilities through the `net` package.
- **TCP Programming**: Setting up TCP listeners (`net.Listen`) and dialing connections (`net.Dial`).
- **UDP Programming**: Connectionless datagram delivery using `net.ListenUDP` and `net.DialUDP`.
- **Client-Server Architecture**: Handling concurrent client connections using goroutines.

### 3. HTTP & Web Development (`03-http-web`)
The `net/http` package provides everything needed to build modern web services.
- **HTTP Server**: Setting up servers, multiplexers (`http.ServeMux`), and writing handlers.
- **HTTP Client**: Building custom HTTP requests, setting headers, cookies, and managing request lifecycles.
- **Middleware Pattern**: Wrapping handler functions to intercept requests for logging, authentication, recovery, etc.

### 4. Web Frameworks (`04-web-frameworks`)
While the standard library is powerful, third-party frameworks simplify routing and standardizing request pipelines.
- **Chi Router**: Lightweight, idiomatic router fully compatible with `net/http`.
- **Gin & Echo**: High-performance HTTP frameworks offering context-rich middleware, parameter binding, and fast routing.

### 5. REST API Development (`05-rest-api`)
A synthesis of routing, middleware, JSON mapping, and clean architecture.
- **API Design**: Mapping REST routes (`GET`, `POST`, `PUT`, `DELETE`).
- **Data Binding & Validation**: Unmarshaling JSON request payloads and validating input fields.
- **Standardized Responses**: Returning uniform JSON structures for success and error conditions.

---

## 🚀 How to Run the Code

Navigate to any subfolder to run the example files or run the tests:

```bash
# Go to the File I/O folder
cd phase-4-io-net-web/01-file-io

# Run the tests to verify practice exercises
go test -v ./...
```
