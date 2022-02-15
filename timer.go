package main

// Need to simplify this shit big time. Strip out the time package, don't need it. Time can be expressed as a simple int, 3600 seconds is one hour
// Every second just decrement one from the time int. Easy

type Timer struct {
	Id         int  `json:"Id"`      // Timer unique ID, 1-20
	Elapsed    int  `json:"Elapsed"` // Time elapsed, in seconds (eg 3600 for 1 hour)
	Running    bool `json:"Running"`
	Reserved   bool `json:"Reserved"` // Reserved for member
	OutOfOrder bool `json:"OutOfOrder"`
}

func (t *Timer) Init(index int) {
	t.Elapsed = 3600
	t.Running = false
	t.Id = index
}

func (t *Timer) StartTimer() {
	t.Running = true
}

func (t *Timer) PauseTimer() {
	t.Running = false
}

func (t *Timer) ClearTimer() {
	t.Elapsed = 0
	t.Running = false
}

func (t *Timer) GetElapsed() int {
	return t.Elapsed
}
