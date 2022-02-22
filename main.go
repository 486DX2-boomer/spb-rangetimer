package main

import (
	"bytes"
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
var updateInit bool = false

func main() {

	for i := range t {
		t[i].Init(i)
	}

	var c Config
	c.Load()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/start/", Start)
	http.HandleFunc("/stop/", Stop)
	http.HandleFunc("/clear/", Clear)
	http.HandleFunc("/update/", Update)
	http.HandleFunc("/getrunning/", GetRunning)
	http.HandleFunc("/getstate/", GetState)
	http.HandleFunc("/member/", SetMember)
	http.HandleFunc("/reserved/", SetReserved)
	http.HandleFunc("/outoforder/", SetOutOfOrder)
	http.HandleFunc("/settime/", setTime)

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

func Clear(w http.ResponseWriter, r *http.Request) {
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/clear/")[1]

	x, _ := strconv.Atoi(timerIndex)

	t[x].ClearTimer()
}

func Update(w http.ResponseWriter, r *http.Request) {

	// Ensure Update is run only once
	if updateInit {
		// Update already initialized
		return
	}

	updateInit = true
	// Update timer each second
	for {
		time.Sleep(1 * time.Second)
		// ClearConsole()
		for i := 1; i < 21; i++ {
			if t[i].Running {
				t[i].Elapsed--
				// Stop timer if out of time
				if t[i].Elapsed <= 0 {
					t[i].Running = false
					t[i].Expired = true
				}
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

func GetState(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // set header to JSON
	w.Header().Set("Access-Control-Allow-Origin", "null")

	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/getstate/")[1]
	index, _ := strconv.Atoi(timerIndex)

	var response bytes.Buffer

	state, err := json.Marshal(t[index])
	if err != nil {
		log.Fatal("JSON marshalling error: ", err)
	}
	response.Write(state)
	// fmt.Println(response.String())
	w.Write(response.Bytes())
	// Reset only needed because I was trying to send everything in one JSON with a loop
	// I want to refactor this so I send one request and get the state of all 20 timers with one request
	// Instead of sending 20 requests every second
	// But I don't know how
	response.Reset()
}

func SetMember(w http.ResponseWriter, r *http.Request) {
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/member/")[1]

	x, _ := strconv.Atoi(timerIndex)

	t[x].SetMember()
}

func SetReserved(w http.ResponseWriter, r *http.Request) {
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/reserved/")[1]

	x, _ := strconv.Atoi(timerIndex)

	t[x].SetReserved()
}

func SetOutOfOrder(w http.ResponseWriter, r *http.Request) {
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string
	timerIndex := strings.Split(ti, "/outoforder/")[1]

	x, _ := strconv.Atoi(timerIndex)

	t[x].SetOutOfOrder()
}

func setTime(w http.ResponseWriter, r *http.Request) {
	// get timer
	ti := fmt.Sprint(r.URL) // Write the r.URL to a string

	timerIndex := strings.Split(ti, "/")[2]
	x, _ := strconv.Atoi(timerIndex)

	fmt.Println(timerIndex)

	// get value to set timer to
	valueStr := strings.Split(ti, "/")[3]
	fmt.Println(valueStr)
	y, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Println("Error: Not an integer? ", err)
		// fallback
		y = 3600
	}
	// Validate the input
	if (y*60 > 3600) || (y < 0) {
		fmt.Println("Error: Time value must be between 0 and 3600 seconds")
		// fallback
		y = 3600
	}

	t[x].Elapsed = y * 60
}
