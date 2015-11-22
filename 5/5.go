/* https://www.codeeval.com/open_challenges/5/
expected output:
6 3 1
49
1 2 3
*/

package main

import "fmt"
import "log"
import "bufio"
import "os"
import "strings"

type pointer struct {
    index int
    value string
    items *[]string
}

func next(item pointer) pointer {
    result := pointer{items: item.items}
    result.index = item.index + 1
    // fmt.Println(result.index, len(*item.items))
    if result.index >= len(*item.items) {
        result.index = len(*item.items) - 1
    }
    result.value = (*item.items)[result.index]

    return result
}

func find_cycle(items []string) []string {
    var result []string
    base := pointer{index: 0, value: items[0], items: &items}
    tortoise := next(base)
    hare := next(next(base))
    for tortoise.value != hare.value {
        // fmt.Println(tortoise, hare)
        tortoise = next(tortoise)
        hare = next(next(hare))
    }
    // fmt.Println(tortoise, hare)


    mu := 0
    tortoise = base
    _ = "breakpoint"
    for tortoise.value != hare.value {
        tortoise = next(tortoise)
        hare = next(hare)
        mu += 1
    }

    // fmt.Println(mu, tortoise, hare)

    lam := 1
    hare = next(tortoise)
    for tortoise.value != hare.value {
        hare = next(hare)
        lam += 1
    }

    // fmt.Println(mu, lam)

    for i := mu; i < mu + lam; i++ {
        result = append(result, items[i])
    }

    return result
}

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        stringArray := strings.Split(scanner.Text(), " ")
        fmt.Println(strings.Join(find_cycle(stringArray), " "))
    }
}
