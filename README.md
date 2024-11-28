# Key-Value Store Project

This repository contains a simple key-value store implementation written in Go. It includes a generic, thread-safe storage system with basic CRUD operations and exposes HTTP endpoints for interacting with the store.

## Project Structure

- **Main Package**:
  - Contains the `main` entry point, initializing the server and starting the HTTP service.
- **Storer Interface**:
  - Defines the contract for the key-value store with methods for `Put`, `Get`, `Update`, and `Delete`.
- **KVStore**:
  - A thread-safe, generic implementation of the `Storer` interface using Go's generics and synchronization primitives.
- **Server**:
  - Implements HTTP handlers using the Echo framework for `Put` and `Get` operations, backed by the `KVStore`.

---

## Features

1. **Generic Key-Value Store**:

   - `KVStore` supports any comparable type for keys and any type for values using Go's generics.
   - Thread-safe operations using `sync.RWMutex` to ensure safe concurrent access.

2. **CRUD Operations**:

   - `Put`: Add or update a key-value pair.
   - `Get`: Retrieve a value by its key.
   - `Update`: Modify an existing key-value pair.
   - `Delete`: Remove a key-value pair from the store.

3. **HTTP API**:
   - `GET /put/:key/:value`: Adds a key-value pair to the store.
   - `GET /get/:key`: Retrieves the value for a given key.

---

## Project Architecture

### Design Patterns

1. **Generic Repository**:

   - The `Storer` interface abstracts the implementation of the key-value store, making it easy to swap or extend the storage backend.
   - `KVStore` implements this interface with a map for simplicity.

2. **Thread-Safe Singleton**:

   - The `KVStore` ensures thread safety with `sync.RWMutex`.
   - Read-heavy operations (`Get`) use a read lock, while write operations (`Put`, `Delete`) use a write lock.

3. **RESTful API Design**:
   - The `Server` exposes HTTP routes to interact with the store, adhering to REST principles for easy client integration.

---

## Installation and Setup

### Prerequisites

- Go 1.18 or later
- `github.com/labstack/echo/v4` for the HTTP server

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/your-repo/kv-store.git
   cd kv-store

   ```

2. Install dependencies:

   ```bash
   go mod tidy

   ```

3. Run the server:
   ```bash
   go run main.go
   ```
   The server will start on http://localhost:3000.

---

## Future Enhancements

1. Add more HTTP endpoints for Update and Delete operations.
2. Implement persistent storage to save the data beyond the application's lifecycle.
3. Introduce authentication and authorization for API endpoints.
4. Enhance error handling and logging mechanisms.
5. Extend support for additional data types for values (e.g., JSON objects).

---
