package receive

import (
    "errors"
    "strings"
    //"strconv"
    "net/http"
    "io/ioutil"    
)

type TestReceiver interface {
    ReceiveRespose(*http.Response, chan string)
}

type BaseReceiver struct {
    UserName string
}

type SessionReceiver struct {
    UserName string
    UserSession string    
}

func extractSession(json string) (session string, err error) {
    session = ""
    err = nil
    if !strings.Contains(json, "SSID") {
        err = errors.New("Result does not contain a session")
        return
    }    
    // TODO: Make less ugly
    end := json[strings.Index(json, "SSID") + len("\"SSID\""):]
    session = strings.Split(end, ",")[0]
    session = session[1:len(session) - 1]        
    return
}

// Base implementation for a TestReceiver
func (r BaseReceiver) ReceiveResponse(resp *http.Response, channel chan string) {
    _, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    channel <- string(r.UserName)
}

// Session enabled Receiver implementation
func (r SessionReceiver) ReceiveResponse(resp *http.Response, channel chan string) {
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }        
    
    session := ""       
    
    if r.UserSession == "" {
        session, err = extractSession(string(body))
        if err != nil {
            channel <- err.Error()
        }
        r.UserSession = session
    } else {
        session = r.UserSession
    }
    
    channel <- string(session)
}