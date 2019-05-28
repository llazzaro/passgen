package main

import (
    "fmt"
    "github.com/llazzaro/passgen/pkg/generator"
)

func main () {
    fmt.Println(generator.ProcessRule("test", "c$2$0$1$9"))
    //for rule := range rules {
    //    parser.ParseRule(rule)
    //}
}
