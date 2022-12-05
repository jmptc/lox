package main

import (
	"fmt"

    "github.com/jmptc/lox/scanner"
)

func main() {
    fmt.Println("Hello lox!")
    /* 
    inputScanner := bufio.NewScanner(os.Stdin)
    for {
        inputScanner.Scan()
        line := inputScanner.Text()
        fmt.Println(line)

        s := scanner.NewScanner(line) 
        s.ScanTokens()
    }
    */

    source := "(){}"
    scanner := scanner.NewScanner(source)
    tokens := scanner.ScanTokens()
    fmt.Println(tokens)
}
