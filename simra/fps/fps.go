package fps

var (
	fpsTimerContainer = make(map[int]*fps)
	mapChan           = make(chan op)
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

func init() {
	go fpsTimerContainerDaemon()
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
	mapChan <- op{"add", fps}
	return fps.c
}

func (f *fps) progress() (int, bool) {
	f.elapsed++
	if f.elapsed >= f.timeToFire {
		f.c <- struct{}{}
		return f.id, true
	}
	return f.id, false
}

// Progress progresses elapsed frames for all timers
func Progress() {
	for _, v := range fpsTimerContainer {
		if id, fired := v.progress(); fired {
			mapChan <- op{"delete", id}
		}
	}
}

func fpsTimerContainerDaemon() {
	for op := range mapChan {
		switch op.op {
		case "add":
			fps := op.value.(*fps)
			fpsTimerContainer[fps.id] = fps
		case "delete":
			id := op.value.(int)
			delete(fpsTimerContainer, id)
		}
	}
}
