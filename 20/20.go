// https://www.codeeval.com/open_challenges/20/
package main

import "fmt"
import "log"
import "bufio"
import "os"
import "bytes"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        byteArray := []byte(scanner.Text())
        fmt.Println(string(bytes.ToLower(byteArray)))
    }
}
