package main

import "fmt"

func fibonacci(n int) []int {
    if n < 1 {
        return []int{}
    }
    if n == 1 {
        return []int{1}
    }
    fib := make([]int, n)
    fib[0], fib[1] = 1, 1
    for i := 2; i < n; i++ {
        fib[i] = fib[i-1] + fib[i-2]
    }
    fmt.Printf("Length: %d, Capacity: %d\n", len(fib), cap(fib))
    return fib
}

func main() {
    var n int
    fmt.Print("gib eine Zahl ein: ")
    fmt.Scan(&n)
    fmt.Println(fibonacci(n))

    // Liste mit Monatsnamen
    months := []string{"Januar", "Februar", "März", "April", "Mai", "Juni", "Juli", "August", "September", "Oktober", "November", "Dezember"}

    // Monat hinzufügen
    var newMonth string
    fmt.Print("Gib einen neuen Monat an: ")
    fmt.Scan(&newMonth)
    months = append(months, newMonth)

    fmt.Println("Der Monat des Jahres ist:")
    for i, month := range months {
        fmt.Printf("%d. %s\n", i+1, month)
    }
}


