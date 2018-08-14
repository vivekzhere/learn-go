package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	now := time.Now()
	rand.Seed(now.UnixNano())
	fmt.Println(now)
	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())
	fmt.Println(now.UTC())
	fmt.Println(now.UTC().Unix())
	fmt.Println(now.UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(10))
}
