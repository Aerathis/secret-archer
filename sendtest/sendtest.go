package sendtest

import (    
    "net/http"
    "io/ioutil"
    "strings"
    "fmt"
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

func (config *HostConfiguration) SendTest() () {
    // Iterate over the commands list
    for i := range config.Commands {
        url := config.HostName + "/" + config.Commands[i].CommandUri
        resp, err := http.Post(url, "application/json", strings.NewReader(config.Commands[i].CommandData))
        if err != nil {
            panic(err)
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            panic(err)
        }
        fmt.Println(string(body))
    }        
}