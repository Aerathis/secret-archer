package monitor

import (
    "os"    
)

type Monitor interface {
    Start()
    Running() bool
}

// Base monitor to satisfy interface
type BaseMonitor struct {
    Channel chan string
}

// Output channel values to a file
type FileOutputMonitor struct {
    Channel chan string
    ReportFileName string
    IsRunning bool
    DebugChan chan string
}

func (m BaseMonitor) Start() {
    for {
        <-m.Channel
    }
}

func (m BaseMonitor) Running() bool {
    return true;
}

func (m FileOutputMonitor) Start() {    
    m.DebugChan <- m.ReportFileName
    f, err := os.OpenFile(m.ReportFileName, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        f, err = os.Create(m.ReportFileName)
        if err != nil {
            panic(err)
        }
    }
    defer f.Close()
    for m.IsRunning {
        recordString := <-m.Channel
        if _, err = f.WriteString(recordString + "\n"); err != nil {
            panic(err)
        }
        m.DebugChan <- recordString
    }
}

func (m FileOutputMonitor) Running() bool {
    return m.IsRunning
}