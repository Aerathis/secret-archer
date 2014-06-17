package configparser

import "errors"

type CommandConfiguration struct {
    CommandName string
    CommandUri string
}

type HostConfiguration struct {
    HostName string
    Port string
    Commands []CommandConfiguration
}

func lines(input string) (lines []string) {
    lines = make([]string, 0, 0)
    lineStart := 0
    for i := 0; i < len(input); i++ {
        if input[i] == '\n' {
            lines = append(lines, input[lineStart:i])
            lineStart = i + 1;
        }
    }
    return
}

func parseLine(input string) (lineLabel string, lineValue string, err error) {
    processed := false
    lineLabel = ""
    lineValue = ""    
    err = nil
    for i := 0; i < len(input); i++ {
        if !processed {
            if input[i] == '|' {
                lineLabel = input[0:i]
                lineValue = input[i+1:]
                processed = true
            }
        }        
    }
    if lineLabel == "" || lineValue == "" {
        lineLabel = ""
        lineValue = ""
        err = errors.New("Malformed Configuration Line")        
    }
    return
}

func ParseConfigString(input string) (resultConfig *HostConfiguration, err error) {
    hostName := ""
    port := "-1"
    
    commandList := make([]CommandConfiguration, 0, 0)
    
    configLines := lines(input)
    
    for i := 0; i < len(configLines); i++ {
        label, value, lineErr := parseLine(configLines[i])
        if lineErr != nil {
            err = lineErr
            return
        }
                
        if label == "HostName" {
            hostName = value
        } else if label == "Port" {
            port = value
        } else {
            commandList = append(commandList, CommandConfiguration{label, value})
        }        
    }
    
    if hostName != "" && port != "-1" {
        resultConfig = &HostConfiguration {hostName, port, make([]CommandConfiguration, 0, 0)}
        err = nil
    } else {
        resultConfig = nil
        err = errors.New("Not Implemented")        
    }    
    return
}
