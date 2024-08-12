package main

import (
	"fmt"
	"in-memory-cache/cache"
	"time"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42, time.Second*5)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
