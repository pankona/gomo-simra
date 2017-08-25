package fps

import "sync"

var (
	fpsTimerContainer sync.Map
	opQueue           = make(chan op)
	timerID           int
)

type fps struct {
	id         int
	timeToFire int64
	elapsed    int64
	c          chan struct{}
}

type op struct {
	op    string
	value interface{}
}

// After waits for the duration (fps based) to elapse
// and then sends the empty channel
func After(timeToFire int64) <-chan struct{} {
	fps := &fps{
		id:         timerID,
		timeToFire: timeToFire,
		c:          make(chan struct{}, 1),
	}
	timerID++
	if timerID > 65535 {
		timerID = 0
	}
	fpsTimerContainer.Store(fps.id, fps)
	return fps.c
}

func (f *fps) progress() (int, bool) {
	if f == nil {
		return -1, false
	}
	f.elapsed++
	if f.elapsed >= f.timeToFire {
		f.c <- struct{}{}
		return f.id, true
	}
	return f.id, false
}

// Progress progresses elapsed frames for all timers
func Progress() {
	fpsTimerContainer.Range(func(_, v interface{}) bool {
		fps := v.(*fps)
		if id, fired := fps.progress(); fired {
			fpsTimerContainer.Delete(id)
		}
		return true
	})
}
