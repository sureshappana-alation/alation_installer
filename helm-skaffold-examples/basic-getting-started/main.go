package main

import (
	"fmt"
	"time"
)

func main() {
	for counter := 0; ; counter++ {
		fmt.Println("Hello Alation!", counter)

		time.Sleep(time.Second * 3)
	}
}
