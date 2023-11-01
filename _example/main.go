package main

import (
	"fmt"
	go_broadcaster "github.com/pefish/go-broadcaster"
	"time"
)

func main() {
	broadcaster := go_broadcaster.NewBroadcaster[int]()

	go func() {
		key := "1"
		broadcaster.Sub(key)
		for {
			select {
			case result := <-broadcaster.C(key):
				fmt.Printf("<%s> %d\n", key, result)
			}
		}
	}()

	go func() {
		key := "2"
		broadcaster.Sub(key)
		for {
			select {
			case result := <-broadcaster.C(key):
				fmt.Printf("<%s> %d\n", key, result)
			}
		}
	}()

	go func() {
		key := "3"
		broadcaster.Sub(key)
		for {
			select {
			case result := <-broadcaster.C(key):
				fmt.Printf("<%s> %d\n", key, result)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	broadcaster.Pub(10)
	broadcaster.Pub(11)

	time.Sleep(5 * time.Second)
}
