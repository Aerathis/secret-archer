package sendtest

import (    
    "net/http"
)

type CommandConfiguration struct {
    CommandName string
    CommandUri string
    CommandData string
}

type HostConfiguration struct {
    HostName string
    Port string
    Commands []CommandConfiguration
}

func (config *HostConfiguration) SendTest() *http.Response {
    resp, err := http.Get(config.HostName)
    if err != nil {
        panic(err)
    }
    return resp
}

func TestNet() *http.Response {
    resp, err := http.Get("http://google.com")
    if err != nil {
        panic(err)
    }
    return resp
}