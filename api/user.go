package main

import (
	"fmt"
	"testdocker/rpc/service"
	"time"
)

func main() {
	fmt.Println("I am api")
	a := service.Service.GetOwnName
	fmt.Println(a)
	for {
		fmt.Println("I am api")
		time.Sleep(1 * time.Second)
	}
}
