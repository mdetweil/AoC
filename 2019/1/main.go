package main


import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    fmt.Println("Total for P1 is: %1", CalculateP1("input.txt"))
    fmt.Println("Total for P2 is: %1", CalculateP2("input.txt"))
}

func CalculateP1(s string) int {
    file, _ := os.Open(s)
    scanner := bufio.NewScanner(file)
    total := 0
	for scanner.Scan() {
        line := scanner.Text()
        i, _ := strconv.Atoi(line)
        total += CalculateResults(i)
	}
    return total
}

func CalculateP2(s string) int {
    file, _ := os.Open(s)
    scanner := bufio.NewScanner(file)
    total := 0
	for scanner.Scan() {
        line := scanner.Text()
        i, _ := strconv.Atoi(line)
        sub := CalculateResults(i)
        sub1 := sub
        for (sub1 >= 0) {
            t1 := CalculateResults(sub1)
            if (t1 >= 0) {
                sub += t1
            }
            sub1 = t1
        }
        total += sub
	}
    return total
}

func CalculateResults(i int) int {
    d := i/3
    return (d - 2)
}

/*
for (t1 >= 0) {
        t2 := CalculateResults(t1)
        if (t2 >= 0) {
            fmt.Println("Adding t1 %i", t2)
            total += t2            
        }
        t1 = t2
    }

    fmt.Println(total)
*/