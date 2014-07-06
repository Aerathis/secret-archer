package receive

import (
    "errors"
    "strings"    
    "net/http"
    "io/ioutil"
    "encoding/json"
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

func extractSession(jsonString string) (session string, err error) {
    session = ""
    var f interface{}
    err = json.Unmarshal([]byte(jsonString), &f)
    if err != nil {        
        return
    }
    
    response := f.(map[string]interface{})   
    
    if _, ok := response["Code"]; ok {
        for k,v := range response {
            if k == "Message" {                
                err = errors.New(v.(string))
            }
        }
    } else {
        if list, ok := response["DataList"]; ok {
            li := list.(map[string]interface{})
            for k, v := range li {
                if k == "SSID" {                    
                    session = v.(string)
                }
            }
        }
    }
    return
}

func extractSessionOld(json string) (session string, err error) {
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