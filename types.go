package timemachine

import (
	"log"
	"time"
)

// GoDefaultFormat is the default formatting string for the Go time library.
const GoDefaultFormat = "Mon Jan 2 15:04:05 MST 2006"

// DefaultTime returns the default time used by the timemachine library.
func DefaultTime() time.Time {
	return must(time.Parse(GoDefaultFormat, GoDefaultFormat))
}

func must(t time.Time, err error) time.Time {
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return t
}

// TimeGetter lets you mock time.Now() and similar functions.
type TimeGetter func() time.Time

// OPCode is an operation. It is combined with data to make an instruction.
type OPCode = string

const (
	// OPAdd will increment the state by the attached amount. Returns the time.
	OPAdd OPCode = "inc"

	// OPRep resets the instruction location back to zero. Does not return the time.
	OPRep = "rep"

	// OPSet sets the current value to the attached time. Does not return the time.
	OPSet = "set"

	// OPGet gets the current time the machine is set to without changing its
	// state. Returns the time.
	OPGet = "get"
)
