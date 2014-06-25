package sendtest

import (    
    "net/http"    
    "strings"
    
    "github.com/Aerathis/secret-archer/receive"
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

func replaceUserToken(rawString, token string) (replacedString string) {
    resultBytes := make([]byte, 0)
    replaced := false
    for i := range rawString {
        if rawString[i] != '|' {
            resultBytes = append(resultBytes, rawString[i])
        } else {
            if !replaced {
                if rawString[i+1] == '|' {
                    for j := range token {
                        resultBytes = append(resultBytes, token[j])
                    }                    
                }
            }            
        }        
    }
    replacedString = string(resultBytes)
    return
}

func (config *HostConfiguration) SendTest(userToken string, channel chan string) () {
    receiver := receive.BaseReceiver{userToken}
    for i := range config.Commands {
        url := config.HostName + "/" + config.Commands[i].CommandUri
        rawData := config.Commands[i].CommandData
        commandString := replaceUserToken(rawData, userToken)        
        resp, err := http.Post(url, "application/json", strings.NewReader(commandString))
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()
        receiver.ReceiveResponse(resp, channel)
    }        
}