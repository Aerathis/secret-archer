package receive

import (
    "net/http"
    "io/ioutil"        
)

type TestReceiver interface {
    ReceiveRespose(*http.Response, chan string)
}

type BaseReceiver struct {
    UserName string
}

// Base implementation for a TestReceiver
func (r BaseReceiver) ReceiveResponse(resp *http.Response, channel chan string) {
    _, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    channel <- string(r.UserName)
}