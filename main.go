package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/luismesas/goPi/piface"
	"github.com/luismesas/goPi/spi"
	"github.com/unrolled/render"
	"net/http"
	"time"
)

func main() {

	// Render engine
	r := render.New(render.Options{
		Layout: "layout",
	})

	// Handlers
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "home", nil)
	})
	mux.HandleFunc("/blink", func(w http.ResponseWriter, req *http.Request) {
		testLed()
		r.Text(w, http.StatusOK, "OK")
	})

	// HTTP Server
	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":3000")
}

func testLed() {

	// creates a new pifacedigital instance
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// initializes pifacedigital board
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}
	pfd.Leds[0].Toggle()
	time.Sleep(time.Second)
	pfd.Leds[0].Toggle()
}
