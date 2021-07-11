package main

import (
    "fmt"
    "strconv"
)

func find_sum_of_divisors(number int) int {
    sum := 0
    for i := 1; i < number; i++ { // TODO: Find a better way to sum factors
        if number % i == 0 {
            sum += i
        }
    }
    return sum
}

func main() {
    deficient := 0
    perfect   := 0
    abundant  := 0
    for number := 1; number <= 20000; number++ {
        sum := find_sum_of_divisors(number)
        if sum == number {
            perfect++
        } else if sum < number {
            deficient++
        } else {
            abundant++
        }
    }
    fmt.Println("Deficient: " + strconv.Itoa(deficient))
    fmt.Println("Perfect:   " + strconv.Itoa(perfect))
    fmt.Println("Abundant:  " + strconv.Itoa(abundant))
}