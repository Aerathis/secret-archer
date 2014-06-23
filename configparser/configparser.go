package configparser

import (
    "github.com/Aerathis/secret-archer/sendtest"
    "errors"
)

func lines(input string) (lines []string) {
    lines = make([]string, 0, 0)
    lineStart := 0
    for i := 0; i < len(input); i++ {
        if input[i] == '\n' {
            lines = append(lines, input[lineStart:i])
            lineStart = i + 1;
        }
    }
    lines = append(lines, input[lineStart:])
    return
}

func chop(input string, chopPoint byte) (front string, back string, err error) {
    processed := false
    front = ""
    back = ""
    err = nil
    for i := 0; i < len(input); i++ {
        if !processed {
            if input[i] == chopPoint {
                front = input[0:i]
                if i < len(input) {
                    back = input[i+1:]
                }                
                processed = true
            }
        }
    }
    if front == "" {
        err = errors.New("Cannot chop this string")
    }
    return
}

func ParseConfigString(input string) (resultConfig *sendtest.HostConfiguration, err error) {
    hostName := ""
    port := "-1"
    
    commandList := make([]sendtest.CommandConfiguration, 0, 0)
    
    configLines := lines(input)       
    
    for i := 0; i < len(configLines); i++ {        
        label, value, lineErr := chop(configLines[i], '|')
        if lineErr != nil {
            err = lineErr
            return
        }
                
        if label == "HostName" {
            hostName = value
        } else if label == "Port" {
            port = value
        } else {         
            commandValue, commandData, commandErr := chop(value, '~')
            if commandErr != nil {
                err = commandErr
                return
            }
            commandList = append(commandList, sendtest.CommandConfiguration{label, commandValue, commandData})
        }        
    }
    
    if hostName != "" && port != "-1" {
        resultConfig = &sendtest.HostConfiguration {hostName, port, commandList}
        err = nil
    } else {
        resultConfig = nil
        err = errors.New("Not Implemented")        
    }    
    return
}
