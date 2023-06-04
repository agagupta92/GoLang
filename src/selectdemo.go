package main

import "fmt"

func main(){

    myChannel := make(chan String)

    anotherChannel := make(chan String)

    go func(){
        myChannel <- "data"
    }()

    go func(){
            anotherChannel <- "data2"
        }()

    /* msg := <-myChannel

    fmt.Println(msg) */

    select{
    case msgFromMyChannel := <- myChannel
        fmt.Println(msgFromMyChannel)
    case msgFromAnotherChannel := <- anotherChannel
            fmt.Println(msgFromAnotherChannel)
    }
}