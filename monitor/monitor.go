package monitor

type Monitor interface {
    Start()
    Running() bool
}

type BaseMonitor struct {
    Channel chan string
}

// Base monitor implementation
func (m BaseMonitor) Start() {
    for {
        <-m.Channel
    }
}

func (m BaseMonitor) Running() bool {
    return true;
}