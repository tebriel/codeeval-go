// Sample code to read in test cases:
package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        input := scanner.Text()
        var total int64 = 0
        for _, value := range input {
            // fmt.Println(index)
            // fmt.Printf("%#U", value)
            intval, _ := strconv.ParseInt(string(value), 10, 0)
            total += intval
        }
        fmt.Println(total)
    }
    //'scanner.Text()' represents the test case, do something with it
}
