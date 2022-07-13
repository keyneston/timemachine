package timemachine

import (
	"log"
	"sync"
	"time"
)

// Instruction represents a single instruction in the virtual machine.
type Instruction struct {
	OP       OPCode
	Duration time.Duration
	Time     time.Time
}

// TimeMachine is the encapsulation of the time virtual machine.
type TimeMachine struct {
	Instructions []Instruction
	Pos          int
	Current      time.Time

	lock *sync.Mutex
}

// New creates a new TimeMachine ready to be used. The default start time is <TODO here>
func New() *TimeMachine {
	return &TimeMachine{
		lock:    &sync.Mutex{},
		Pos:     0,
		Current: DefaultTime(),
	}
}

// Time processes the virtual machine until it reaches a value to return.
func (tm *TimeMachine) Time() time.Time {
	tm.lock.Lock()
	defer tm.lock.Unlock()

	count := 0
	for {
		count++
		if count >= 256 {
			log.Fatalf("TimeMachine appears to be stuck in an infinite loop. Instructions: %v", tm.Instructions)
		}
		if len(tm.Instructions) < tm.Pos {
			log.Fatalf("Invalid instruction %d", tm.Pos) // TODO, what to do here? Fatal sucks
		}

		cur := tm.Instructions[tm.Pos]
		tm.Pos++

		switch cur.OP {
		case OPAdd:
			tm.Current = tm.Current.Add(cur.Duration)
			return tm.Current

		case OPSet:
			tm.Current = cur.Time

		case OPRep:
			tm.Pos = 0

		case OPGet:
			return tm.Current

		default:
			log.Fatalf("Unknown OP: %v", cur.OP)
		}
	}
}

// Add uses the builder pattern to add an Add instruction to the end of the
// Instruction list. Use a negative value to subtract.
func (tm *TimeMachine) Add(d time.Duration) *TimeMachine {
	tm.Instructions = append(tm.Instructions, Instruction{OP: OPAdd, Duration: d})
	return tm
}

// Repeat uses the builder
func (tm *TimeMachine) Repeat() *TimeMachine {
	tm.Instructions = append(tm.Instructions, Instruction{OP: OPRep})
	return tm
}

// Set uses the builder pattern to add a Set instruction to the end of the
// Instruction list.
func (tm *TimeMachine) Set(value time.Time) *TimeMachine {
	tm.Instructions = append(tm.Instructions, Instruction{OP: OPSet, Time: value})

	return tm
}

// Get uses the builder pattern to add a Get instruction to the end of the
// Instruction list.
func (tm *TimeMachine) Get() *TimeMachine {
	tm.Instructions = append(tm.Instructions, Instruction{OP: OPGet})

	return tm
}
