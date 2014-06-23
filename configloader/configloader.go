package configloader

import (
    "io/ioutil"    
    
    . "github.com/Aerathis/secret-archer/sendtest"
    . "github.com/Aerathis/secret-archer/configparser"
)

func GetConfig(configFile string) (config *HostConfiguration) {
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