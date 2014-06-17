package configloader

import (
    "io/ioutil"
    "fmt"
    
    . "github.com/Aerathis/secret-archer/configparser"
)

func GetConfig(configFile string) (config HostConfiguration) {
    configContents, err := ioutil.ReadFile(configFile)
    if err != nil {
        panic(err)
    }
    
    configString := string(configContents[:])
    
    fmt.Println(configString)
    
    configs, configErr := ParseConfigString(configString)
    if configErr != nil {
        panic(configErr)
    }
    
    fmt.Println(configs)
    
    commandList := []CommandConfiguration{CommandConfiguration{"", ""}}
    config = HostConfiguration{"", "0", commandList}
    return
}