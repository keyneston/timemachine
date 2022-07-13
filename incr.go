package timemachine

import "time"

type Incr struct {
	delta time.Duration
	time  time.Time
}

func (i *Incr) Time() time.Time {
	return i.time
}
