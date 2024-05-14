package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I am rpc")
		time.Sleep(1 * time.Second)
	}
}
