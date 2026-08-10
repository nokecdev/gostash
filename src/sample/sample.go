package main

import (
	"cache"
	"fmt"
	"time"
)

func main() {
	defaultExpiration, _ := time.ParseDuration("0.5h")
	gcInterval, _ := time.ParseDuration("3s")
	c := cache.NewCache(defaultExpiration, gcInterval)

	k1 := "Initial cache data"
	expiration, _ := time.ParseDuration("5s")

	c.Set("k1", k1, expiration)
	s, _ := time.ParseDuration("10s")
	if v, found := c.Get("k1"); found {
		fmt.Println("Found k1: ", v)
	} else {
		fmt.Println("Not found k1")
	}
	// Pause for 10 seconds
	time.Sleep(s)
	// Now k1 should have been cleared
	if v, found := c.Get("k1"); found {
		fmt.Println("Found k1: ", v)
	} else {
		fmt.Println("Not found k1")
	}
}
