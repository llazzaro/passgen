package test

import (
    "testing"
    "github.com/llazzaro/passgen/pkg/generator"
)

func TestSimpleGenerationFromRule(t *testing.T) {
    input := make(chan string)
    ruleCase := "c $2$0$1$9"
    output := generator.Generate(input, ruleCase)
    input <- "test"

    expected := "Test2019"
    found := <-output
    if found != expected {
        t.Errorf("Expected %s, found %s ", expected, found)
    }
}
