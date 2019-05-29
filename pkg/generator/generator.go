package generator

import (
    "fmt"
    "sync"
    "errors"
    "strings"
)

func ProcessRule(word string, rule string) (string, error) {
    var res string
    var prepend, appen , toggleN bool
    var appendStr, prependStr string
    if len(word) < 1 {
        return "", nil
    }

    for _, char := range rule {
        if char == ' ' {
            continue
        } else if char == ':' {
            continue
        } else if appen {
            appendStr += string(char)
            appen = false
        } else if prepend {
            prepend = false
            prependStr += string(char)
        }  else if toggleN {
            return "", errors.New("Not implemented")
        } else if char == 'c' {
            res += strings.Title(word)
        } else if char == 'd' {
            res += word
        } else if char == '^' {
            prepend = true
        } else if char == '$' {
            appen = true
        } else if char == 'l' {
            res += strings.ToLower(word)
        } else if char == 'u' {
            res += strings.ToUpper(word)
        } else if char == 'C' {
            res += strings.ToLower(string(word[0]))
            res += strings.ToUpper(word[1:])
        } else if char == 't' {
            return "", errors.New("Not implemented")
        } else if char == 'T' {
            return "", errors.New("Not implemented")
        }
    }
    return prependStr + res + appendStr, nil
}

func Generate(input chan string, rule string) chan string {
    var wg sync.WaitGroup
    var err error
    var generatedWord string
    wg.Add(1)

    output := make(chan string)

    go func() {
        for word := range input {
            generatedWord, err = ProcessRule(word, rule)
            if err != nil {
                fmt.Println(err)
            }
            output <- generatedWord
        }

        wg.Done()
    }()

    go func() {
        wg.Wait()
        close(output)
    }()
    return output
}
