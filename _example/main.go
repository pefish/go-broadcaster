package main

import (
	"fmt"
	go_broadcaster "github.com/pefish/go-broadcaster"
	"time"
)

func main() {
	testBroadcaster := go_broadcaster.NewBroadcaster[int]()

	go func() {
		name := "1"
		testBroadcaster.Sub(name)
		for {
			select {
			case result := <-testBroadcaster.C(name):
				fmt.Printf("<%s> %d\n", name, result)
			}
		}
	}()

	go func() {
		name := "2"
		testBroadcaster.Sub(name)
		for {
			select {
			case result := <-testBroadcaster.C(name):
				fmt.Printf("<%s> %d\n", name, result)
			}
		}
	}()

	go func() {
		name := "3"
		testBroadcaster.Sub(name)
		for {
			select {
			case result := <-testBroadcaster.C(name):
				fmt.Printf("<%s> %d\n", name, result)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	testBroadcaster.Pub(10)
	testBroadcaster.Pub(11)

	time.Sleep(5 * time.Second)
}
