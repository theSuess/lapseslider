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

var running bool

func main() {
	// Render engine
	r := render.New(render.Options{
		Layout: "layout",
	})

	// Handlers
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "home", nil)
	})
	router.HandleFunc("/about", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "about", nil)
	})
	router.HandleFunc("/start", func(w http.ResponseWriter, req *http.Request) {
		time := req.FormValue("time")
		stepsPerMinute, _ := strconv.Atoi(req.FormValue("steps"))
		var hours, minutes, seconds int
		fmt.Sscanf(time, "%d:%d:%d", &hours, &minutes, &seconds)
		fmt.Println(hours, minutes, seconds)
		go startLapse((hours*3600)+(minutes*60)+seconds, stepsPerMinute)
		r.HTML(w, http.StatusOK, "home", nil)
	}).Methods("POST")

	// HTTP Server
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":3000")
}

func startLapse(seconds int, stepsPerMinute int) {
	go spin(stepsPerMinute)
	time.Sleep(time.Second * time.Duration(seconds))
	running = false
}

func spin(stepsPerMinute int) {
	var pauseTime float64
	pauseTime = float64(60) / float64(stepsPerMinute)
	pfd := piface.NewPiFaceDigital(spi.DEFAULT_HARDWARE_ADDR, spi.DEFAULT_BUS, spi.DEFAULT_CHIP)

	// initializes pifacedigital board
	err := pfd.InitBoard()
	if err != nil {
		fmt.Printf("Error on init board: %s", err)
		return
	}

	pfd.OutputPins[2].SetValue(0)
	running = true
	for running {
		pfd.OutputPins[2].Toggle()
		time.Sleep(time.Second / 10)
		pfd.OutputPins[2].Toggle()
		time.Sleep(time.Duration(float64(time.Second) * pauseTime))
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
