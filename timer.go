package main

// Need to simplify this shit big time. Strip out the time package, don't need it. Time can be expressed as a simple int, 3600 seconds is one hour
// Every second just decrement one from the time int. Easy

type Timer struct {
	Id         int  `json:"Id"`      // Timer unique ID, 1-20
	Elapsed    int  `json:"Elapsed"` // Time elapsed, in seconds (eg 3600 for 1 hour)
	Running    bool `json:"Running"`
	OutOfOrder bool `json:"OutOfOrder"`
	Member     bool `json:"Member"`
	Reserved   bool `json:"Reserved"` // Reserved for member
	Expired    bool `json:"Expired"`  // True if time left is 0
}

func (t *Timer) Init(index int) {
	t.Elapsed = 3600
	t.Running = false
	t.Id = index
	t.Reserved = false
	t.OutOfOrder = false
	t.Member = false
	t.Expired = false
}

func (t *Timer) StartTimer() {
	t.Running = true
}

func (t *Timer) PauseTimer() {
	t.Running = false
}

func (t *Timer) ClearTimer() {
	t.Elapsed = 3600
	t.Running = false
	t.OutOfOrder = false
	t.Member = false
	t.Reserved = false
	t.Expired = false
}

func (t *Timer) GetElapsed() int {
	return t.Elapsed
}

func (t *Timer) SetOutOfOrder() {
	if !t.OutOfOrder {
		t.OutOfOrder = true
	} else {
		t.OutOfOrder = false
	}
}

func (t *Timer) SetMember() {
	if !t.Member {
		t.Member = true
	} else {
		t.Member = false
	}
}

func (t *Timer) SetReserved() {
	if !t.Reserved {
		t.Reserved = true
	} else {
		t.Reserved = false
	}
}
