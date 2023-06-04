package main

import "fmt"

func main(){

    myChannel := make(chan String)

    go func(){
        myChannel <- "data"
    }()

    msg := <-myChannel

    fmt.Println(msg)
}