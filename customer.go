package main

import "fmt"

func init() {
    fmt.Println("Customer init")
}

type CustomerContext struct {
    Id string
}
