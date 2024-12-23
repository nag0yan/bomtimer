package main

import (
	"fmt"
	"time"
)

func main() {
	var input int
	fmt.Println("Set bom timer")
	for {
		if _, err := fmt.Scan(&input); err != nil {
			fmt.Println("Time should be integer (seconds)")
			continue
		}
		if input == -1 {
			fmt.Println("Bye!")
			break
		} else if input < 0 {
			fmt.Println("Time should be more than zero")
			continue
		}
		c := make(chan string)
		go bom(c, time.Second*time.Duration(input))
		go func() { fmt.Println(<-c) }()
	}
}

func bom(c chan string, t time.Duration) {
	time.Sleep(t)
	c <- "BOM!"
}
