package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

// think I need a global array of timers declared here
var t Timer

func main() {

	t.Init()

	var c Config
	c.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/start", Start)
	http.HandleFunc("/stop", Stop)

	log.Println("Starting server...")
	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, webpage")
	fmt.Println("Hello, console")
}

func ClearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Start(w http.ResponseWriter, r *http.Request) {

	if t.Running {
		fmt.Fprintf(w, "The timer is running")
	} else {
		fmt.Fprintf(w, "Timer not running.")
	}

	t.StartTimer()

	for t.Running {
		ClearConsole()
		fmt.Println(t.Elapsed)
		time.Sleep(1 * time.Second)
		t.Elapsed--
	}

	if !t.Running {
		fmt.Println("Timer not running.")
	}

	// so something is fucked with this, if I call StartTimer() my console prints that the timer started. The problem is that the webpage reports it has NOT started.
	// Calling the StartTimer() method seems to do FUCKING NOTHING
	// However, it works completely fine when directly modifying the bool from this function. What the fuck?
	// OK I figured it out. My methods were written as func (t Timer) instead of func (t *Timer) so my methods were not actually mutating the struct they were pointed at
}

func Stop(w http.ResponseWriter, r *http.Request) {
	t.PauseTimer()
}
