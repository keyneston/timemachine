# TimeMachine

[![Go Reference](https://pkg.go.dev/badge/github.com/keyneston/timemachine.svg)](https://pkg.go.dev/github.com/keyneston/timemachine)

A library for mocking time in go. Using the builder pattern a tiny virtual
machine is programmed to take the desired steps for mocking.

## Adding to your structs

```golang
type MyStruct struct {
	time timemachine.TimeGetter
}

func (m MyStruct) Time() time.Time {
	return m.time()
}

func New() MyStruct {
	return MyStruct{
		time: time.Now, // Easily configure to use `time.Now` in production.
	}

}

func TestMyStruct(t *testing.T) {
	m := MyStruct{
		// And configure to use the mock during testing.
		time: timemachine.New().Set(myDefaultStartTime).Get().Repeat(),
	}

	if m.Time() != myDefaultStartTime {
		t.Error("Oh nooes!")
	}
}
```

## Examples

```golang
// Continually increase the clock by one second.
tm := timemachine.New().Add(time.Second).Repeat()

// Replicate issues of wonky NTP servers.
tm := timemachine.New().Add(time.Second).Add(time.Second * -1).Repeat()

// Set jumps in the time:
tm := timemachine.New().Set(time1).Get().Set(time2).Get().Repeat()
```
