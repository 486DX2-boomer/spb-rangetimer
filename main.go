package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

// think I need a global array of timers declared here
var t [21]Timer

func main() {

	for i := range t {
		t[i].Init()
	}

	var c Config
	c.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/start/", Start)
	http.HandleFunc("/stop/", Stop)
	http.HandleFunc("/refresh/", Refresh)

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

	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/start/")[1]

	x, _ := strconv.Atoi(timerIndex)

	t[x].StartTimer()

	// so something is fucked with this, if I call StartTimer() my console prints that the timer started. The problem is that the webpage reports it has NOT started.
	// Calling the StartTimer() method seems to do FUCKING NOTHING
	// However, it works completely fine when directly modifying the bool from this function. What the fuck?
	// OK I figured it out. My methods were written as func (t Timer) instead of func (t *Timer) so my methods were not actually mutating the struct they were pointed at
}

func Stop(w http.ResponseWriter, r *http.Request) {
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/stop/")[1]

	x, _ := strconv.Atoi(timerIndex)
	t[x].PauseTimer()
}

// I tried to write a buffered version here but it doesn't help at all.
// I also do not understand the concept of double buffering so this is useless until I figure it out
// func Refresh(w http.ResponseWriter, r *http.Request) {
// 	var s string
// 	var buffer string
// 	for {
// 		fmt.Print(buffer)
// 		time.Sleep(1 * time.Second)
// 		ClearConsole()
// 		s = ""
// 		for i := 1; i < 21; i++ {
// 			if t[i].Running {
// 				s += "Timer "
// 				s += strconv.Itoa(i)
// 				s += ": "
// 				s += strconv.Itoa(t[i].Elapsed)
// 				s += "\n"
// 				t[i].Elapsed--
// 			} else if !t[i].Running {
// 				s += "Timer "
// 				s += strconv.Itoa(i)
// 				s += ": "
// 				s += "PAUSED"
// 				s += "\n"
// 			}
// 		}
// 		buffer = s
// 	}
// }

func Refresh(w http.ResponseWriter, r *http.Request) {
	for {
		time.Sleep(1 * time.Second)
		ClearConsole()
		for i := 1; i < 21; i++ {
			if t[i].Running {
				fmt.Println("Timer ", i, ":", t[i].Elapsed)
				t[i].Elapsed--
			} else if !t[i].Running {
				fmt.Println("Timer ", i, ":", "PAUSED")
			}
		}
	}
}
