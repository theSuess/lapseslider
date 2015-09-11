package main

import (
	"fmt"
	"github.com/luismesas/goPi/piface"
	"github.com/luismesas/goPi/spi"
	"time"
)

func main() {

	// creates a new pifacedigital instance
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// initializes pifacedigital board
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}

	for {
		if pfd.InputPins[1].Value == 1 {
			pfd.Leds[7].Toggle()
		}
		time.Sleep(time.Second)
	}
}
