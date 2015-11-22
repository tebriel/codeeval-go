package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"
import "strconv"
import "math"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        var i float64
        //'scanner.Text()' represents the test case, do something with it
        bounds := strings.Split(scanner.Text(), " ")
        fizz, _ := strconv.ParseFloat(bounds[0], 64)
        buzz, _ := strconv.ParseFloat(bounds[1], 64)
        stop, _ := strconv.ParseFloat(bounds[2], 64)

        for i = 1; i <= stop; i++ {
            var output string
            if math.Mod(i, fizz) == 0 {
                output = "F"
            }

            if math.Mod(i, buzz) == 0 {
                output = output + "B"
            }

            if len(output) == 0 {
                output = strconv.Itoa(int(i))
            }

            fmt.Print(output + " ")
        }
        fmt.Print("\n")
    }
}
