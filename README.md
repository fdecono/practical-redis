# Practical Redis Examples

This project demonstrates various Redis operations using Go, including caching, key-value storage, and rate limiting.

## Prerequisites

- Go 1.19 or later
- Docker (for running Redis)

## Setup

1. **Start Redis using Docker:**
   ```bash
   docker run -d --name redis -p 6379:6379 redis
   ```

2. **Install Go dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the examples:**
   ```bash
   go run main.go [example]
   ```

## Available Examples

### 1. Basic Cache Example (`1-cache`)
Demonstrates basic caching with expiration:
```bash
go run main.go 1-cache
```
- Sets a weather cache with 5-minute expiration
- Retrieves and displays the cached value

### 2. Write Cache (`2-write-cache`)
Allows you to write custom key-value pairs to Redis:
```bash
go run main.go 2-write-cache <key> <value>
```
Example:
```bash
go run main.go 2-write-cache "user:123" "John Doe"
```

### 3. Read Cache (`3-read-cache`)
Retrieves values from Redis by key:
```bash
go run main.go 3-read-cache <key>
```
Example:
```bash
go run main.go 3-read-cache "user:123"
```

### 4. Rate Limiting (`4-rate-limit`)
Demonstrates rate limiting using Redis counters:
```bash
go run main.go 4-rate-limit
```
- Implements a simple rate limiter (5 requests per minute)
- Uses Redis INCR and EXPIRE commands
- Tracks requests by IP address

## Project Structure

- `main.go` - Main application with all Redis examples
- `go.mod` - Go module dependencies
- `go.sum` - Dependency checksums

## Redis Operations Used

- `SET` - Store key-value pairs with expiration
- `GET` - Retrieve values by key
- `INCR` - Increment counters
- `EXPIRE` - Set key expiration

## Error Handling

The examples include basic error handling for Redis operations, particularly for cache misses and connection issues.

## Cleanup

To stop and remove the Redis container:
```bash
docker stop redis
docker rm redis
```
