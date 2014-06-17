//Package secret-archer provides a load tester for a ZooTycoon server

package main

import (       
    "os" 
    "fmt"
    "github.com/Aerathis/secret-archer/configloader"
)

func main() {
    argsToProg := os.Args[1:]       
    
    if len(argsToProg) < 1 {
        // Print usage notes
        fmt.Println("Not enough args")
        return
    }
    configFile := argsToProg[0]
    
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        fmt.Println("Config file not present -", configFile)
        return
    }
    
    testConfiguration := configloader.GetConfig(configFile)
    fmt.Println(testConfiguration)
}