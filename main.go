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

// GetState sends the state of every timer in []t as JSON.
func GetState(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") // set header to JSON
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var response bytes.Buffer

	state, err := json.Marshal(t)
	if err != nil {
		log.Fatal("JSON marshalling error: ", err)
	}
	response.Write(state)
	w.Write(response.Bytes())
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

	// get value to set timer to
	valueStr := strings.Split(ti, "/")[3]
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
