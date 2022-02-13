package main

import (
	"encoding/json"
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
var refreshInit bool = false

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
	http.HandleFunc("/getrunning/", GetRunning)

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

	if !t[x].Running {
		t[x].StartTimer()
	}
}

func Stop(w http.ResponseWriter, r *http.Request) {
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/stop/")[1]

	x, _ := strconv.Atoi(timerIndex)

	if t[x].Running {
		t[x].PauseTimer()
	}
}

func Refresh(w http.ResponseWriter, r *http.Request) {

	// NOTE: Need a method to make sure refresh only runs ONCE, multiple refreshes make the timer run at sonic the hedgehog speed

	if refreshInit {
		fmt.Println("Refresh already init")
		return
	}

	refreshInit = true

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

func GetRunning(w http.ResponseWriter, r *http.Request) {

	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/getrunning/")[1]
	index, _ := strconv.Atoi(timerIndex)

	w.Header().Set("Content-Type", "application/json") // set header to JSON
	w.Header().Set("Access-Control-Allow-Origin", "null")
	response := make(map[string]string)
	if t[index].Running {
		response["Running"] = "true"
	} else {
		response["Running"] = "false"
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal("JSON marshalling error: ", err)
	}
	// fmt.Println(string(jsonResponse))
	w.Write(jsonResponse)
}
