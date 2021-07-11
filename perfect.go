package main

import (
    "fmt"
)

func find_sum_of_divisors(number int) int {
    sum := 0
    for i := 1; i <= number/2; i++ { // TODO: Find a better way to sum factors
        if number % i == 0 {
            sum += i
        }
    }
    return sum
}

func main() {
    var deficient, abundant, perfect = 0,0,0
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
    fmt.Printf("Deficient: %d\n", deficient)
    fmt.Printf("Perfect:   %d\n", perfect)
    fmt.Printf("Abundant:  %d\n", abundant)
}