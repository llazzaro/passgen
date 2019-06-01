package test

import (
    "testing"
    "gitlab.com/llazzaro/passgen/pkg/generator"
)

func checkRule(t *testing.T, word string, rule string, expected string) {
    input := make(chan string)
    output := generator.Generate(input, rule)
    input <- word

    found := <-output
    if found != expected {
        t.Errorf("Expected %s, found %s ", expected, found)
    }

}


func TestSimpleGenerationFromRule(t *testing.T) {
    rule := "c $2$0$1$9"
    expected := "Test2019"
    word := "test"
    checkRule(t, word, rule, expected)
}


func TestDoNothing(t *testing.T) {
    rule := ":"
    expected := "test"
    word := "test"
    checkRule(t, word, rule, expected)
}


func TestAppendOne(t *testing.T) {
    rule := "$1"
    expected := "test1"
    word := "test"
    checkRule(t, word, rule, expected)
}


func TestPrependOne(t *testing.T) {
    rule := "^1"
    expected := "1test"
    word := "test"
    checkRule(t, word, rule, expected)
}
