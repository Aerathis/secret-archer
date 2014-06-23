//Package secret-archer provides a load tester for a ZooTycoon server

package main

import (       
    "os" 
    "fmt"
    "github.com/Aerathis/secret-archer/configloader"
    "github.com/Aerathis/secret-archer/sendtest"
)

func main() {
    argsToProg := os.Args[1:]       
    
    if len(argsToProg) < 1 {
        // Print usage notes
        fmt.Println("Not enough args")
        return
    }
    configFile := argsToProg[0]
    
    //concurrencyLevel := argsToProg[1]
    
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        fmt.Println("Config file not present -", configFile)
        return
    }
    
    testConfiguration := configloader.GetConfig(configFile)      
    fmt.Println(testConfiguration)
    
    testConfiguration.SendTest()
    
    testResp := sendtest.TestNet()
    fmt.Println(testResp.Body)
}