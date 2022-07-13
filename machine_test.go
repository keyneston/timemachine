package timemachine

import (
	"testing"
	"time"
)

func TestTimeMachineRepeating(t *testing.T) {
	tm := New().Add(time.Second).Repeat()

	values := []time.Time{}

	for i := 0; i < 10; i++ {
		values = append(values, tm.Time())
	}

	expected := time.Second * 1
	for i := 1; i < 10; i++ {
		diff := values[i].Sub(values[i-1])

		if diff != expected {
			t.Errorf("%q - %q = %s want %s", values[i], values[i-1], diff, time.Second)
		}
	}
}

func TestTimeMachineSet(t *testing.T) {
	start := must(time.Parse(GoDefaultFormat, "Wed Jan 1 13:01:01 MST 2020"))

	tm := New().Set(start).Get()

	value := tm.Time()
	if value != start {
		t.Errorf("tm.Time() = %q; want %q", value, start)
	}
}
