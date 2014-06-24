package config

import (
    "github.com/Aerathis/secret-archer/sendtest"
    "errors"
    "io/ioutil"
)

func lines(input string) (lines []string) {
    lines = make([]string, 0, 0)
    lineStart := 0
    for i := range input {
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
    for i := range input {
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
    
    for i := range configLines {        
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

func GetConfig(configFile string) (config *sendtest.HostConfiguration) {
    configContents, err := ioutil.ReadFile(configFile)
    if err != nil {
        panic(err)
    }
    
    configString := string(configContents[:])    
    
    config, configErr := ParseConfigString(configString)
    if configErr != nil {
        panic(configErr)
    }                  
    return
}
