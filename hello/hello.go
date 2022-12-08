package main

import (
    "fmt"
    "rsc.io/quote"
    "go_training/greetings"
)

func main() {
    fmt.Println(quote.Go())

    message := greetings.Hello("Amid")

    fmt.Println(message)
}
