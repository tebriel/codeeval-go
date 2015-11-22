// https://www.codeeval.com/open_challenges/3/
package main

import "fmt"
// import "log"
// import "bufio"
// import "os"
import "strconv"
import "math/big"

func main() {
    found_it := false
    NumGenerator:
        for num := 1000; num >= 0; num-- {
            str_num := []byte(strconv.Itoa(num))
            for i := 0; i < len(str_num) / 2; i++ {
                if str_num[i] != str_num[len(str_num) - 1 - i] {
                    found_it = false
                    break
                } else {
                    found_it = true
                }
            }
            big_int := big.NewInt(int64(num))
            if found_it && big_int.ProbablyPrime(20) {
                fmt.Println(num)
                break NumGenerator
            }
        }
}
