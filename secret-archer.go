//Package secret-archer provides a load tester for a ZooTycoon server

package main

import (       
    "os" 
    "fmt"
    "strconv"
    "github.com/Aerathis/secret-archer/config"    
)

func main() {
    argsToProg := os.Args[1:]       
    
    if len(argsToProg) != 2 {
        // Print usage notes
        fmt.Println("Not enough args")
        return
    }
    configFile := argsToProg[0]
    
    concurrencyString := argsToProg[1]
    concurrencyLevel, err := strconv.ParseInt(concurrencyString, 10, 64)
    if err != nil {
        panic(err)
    }
    
    if _, err := os.Stat(configFile); os.IsNotExist(err) {
        fmt.Println("Config file not present -", configFile)
        return
    }
    
    testConfiguration := config.GetConfig(configFile)
    fmt.Println(testConfiguration)
    
    c := make(chan string)
    
    for i := 0; i < int(concurrencyLevel); i++ {
        userString := "stresstestuser" + strconv.Itoa(i)
        go testConfiguration.SendTest(userString, c)
    }
    for {
        fmt.Println(<-c)
    }    
}