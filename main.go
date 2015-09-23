package main

import (
	"fmt"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/luismesas/goPi/piface"
	"github.com/luismesas/goPi/spi"
	"github.com/unrolled/render"
	"net/http"
	"strconv"
	"time"
)

var progress int
var running bool

func main() {
	// Render engine
	r := render.New(render.Options{
		Layout: "layout",
	})

	// Handlers
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "home", progress)
	})
	router.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "about", nil)
	})
	router.HandleFunc("/start", func(w http.ResponseWriter, req *http.Request) {
		time := req.FormValue("time")
		stepsPerMinute, _ := strconv.Atoi(req.FormValue("steps"))
		var hours, minutes, seconds int
		fmt.Sscanf(time, "%d:%d:%d", &hours, &minutes, &seconds)
		go startLapse((hours*3600)+(minutes*60)+seconds, stepsPerMinute)
		r.HTML(w, http.StatusOK, "home", progress)
	}).Methods("POST")

	// HTTP Server
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}

func startLapse(seconds int, stepsPerMinute int) {
	running = true
	seconds = seconds * 10
	var pauseTime float64
	pauseTime = float64(stepsPerMinute) / float64(60.0)
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// initializes pifacedigital board
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}

	for i := 0; i < seconds && running; i++ {
		pfd.OutputPins[0].Toggle()
		time.Sleep(time.Duration(float64(time.Second) * pauseTime))
		pfd.OutputPins[0].Toggle()
	}
}

func testMotor() {
	// creates a new pifacedigital instance
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// initializes pifacedigital board
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}
	pfd.OutputPins[0].Toggle()
	time.Sleep(time.Second / 10)
	pfd.OutputPins[0].Toggle()
}
