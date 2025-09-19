package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [example]")
		fmt.Println("Example:\n 1-cache\n 2-write-cache <key> <value>\n 3-read-cache <key>\n 4-rate-limit")
		return
	}

	switch os.Args[1] {
	case "1-cache":
		cacheExample()
	case "2-write-cache":
		writeCache()
	case "3-read-cache":
		readCache()
	case "4-rate-limit":
		rateLimit()
	default:
		fmt.Println("Invalid example number")
	}
}

func cacheExample() {
	rdb.Set(ctx, "weather:bsas", "Sunny 22 C", 5*time.Minute)
	value, _ := rdb.Get(ctx, "weather:bsas").Result()
	fmt.Println("Cached weather in Buenos Aires:", value)
}

func writeCache() {
	rdb.Set(ctx, os.Args[2], os.Args[3], 5*time.Minute)
	fmt.Println("Cached", os.Args[2], "with value:", os.Args[3])
}

func readCache() {
	value, err := rdb.Get(ctx, os.Args[2]).Result()
	if err != nil {
		fmt.Println("Error getting", os.Args[2], ">", err)
		return
	}
	fmt.Println("Got", os.Args[2], "with value:", value)
}

func rateLimit() {
	key := "rate:ip:127.0.0.1"
	count, _ := rdb.Incr(ctx, key).Result()
	if count == 1 {
		rdb.Expire(ctx, key, 1*time.Minute)
	}
	fmt.Printf("Request nr #%d\n", count)
	if count > 5 {
		fmt.Println("Too many requests")
	}
}
