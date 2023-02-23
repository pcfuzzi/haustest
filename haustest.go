package main

import (
"bufio"
"fmt"
"math/rand"
"os"
"time"
)

func main() {
const maxCount int = 5
var count int
rand.Seed(time.Now().UnixNano())

fmt.Print("Geben Sie die Anzahl der Zahlen ein (maximal 10): ")
reader := bufio.NewReader(os.Stdin)
input, _ := reader.ReadString('\n')
fmt.Sscanf(input, "%d", &count)

if count > maxCount {
	fmt.Printf("Es können nur maximal %d Zahlen eingegeben werden.\n", maxCount)
	return
}

numbers := make([]int, count)
for i := 0; i < count; i++ {
	numbers[i] = rand.Intn(100)
	fmt.Printf("Die zufällige %d. Zahl ist: %d\n", i+1, numbers[i])
}

sum := 0
for _, number := range numbers {
	sum += number
}
average := float64(sum) / float64(count)

min := numbers[0]
max := numbers[0]
for _, number := range numbers {
	if number < min {
		min = number
	}
	if number > max {
		max = number
	}
}

fmt.Printf("Summe: %d\n", sum)
fmt.Printf("Durchschnitt: %.2f\n", average)
fmt.Printf("Kleinste Zahl: %d\n", min)
fmt.Printf("Größte Zahl: %d\n", max)
}
