package main

import (
	"fmt"
	"time"
	b1 "github.com/hink/go-blink1"
)

const (
    USBVendorID  = 10168
    USBProductID = 493
)

func main() {
	blink, err := b1.OpenNextDevice()
	if err != nil {
		fmt.Println("Couldn't open Blink")
		return
	}
	blink.SetState(b1.State{Red: 255, Duration: time.Second * 30})
	fmt.Println("Called setState, exiting")
}

