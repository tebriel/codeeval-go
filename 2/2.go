package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strconv"
import "sort"

type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

func main() {
    var output_count int = -1
    var output_count64 int64
    var i int
    var lines []string
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        if output_count == -1 {
            output_count64, _ = strconv.ParseInt(string(scanner.Text()), 10, 0)
            output_count = int(output_count64)
        } else {
            lines = append(lines, scanner.Text())
        }
        // fmt.Println(string(bytes.ToLower(byteArray)))
    }

    sort.Sort(ByLength(lines))
    for i = len(lines) - 1; i >= len(lines) - output_count; i-- {
        fmt.Println(lines[i])
    }
}
