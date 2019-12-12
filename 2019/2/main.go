package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	p1 := CalculateP1(getArray("i1.txt"), 12, 2)
	fmt.Printf("The results for Part 1 are: %v \n", p1)
	p2 := CalculateP2("i2.txt")
	fmt.Printf("The results for Part 2 are: %v \n", p2)
}

// CalculateP1 Calculate the results of Part 1
func CalculateP1(arr []int64, noun int, verb int) int64 {
	arr[1] = int64(noun)
	arr[2] = int64(verb)
	for i := 0; i < len(arr); i+=4 {
		switch arr[i] {
		case 1:
			//fmt.Println("ADDING BIT FOUND")
			arr[arr[i+3]] = arr[arr[i+1]] + arr[arr[i+2]]
		case 2:
			//fmt.Println("MULT BIT FOUND")
			arr[arr[i+3]] = arr[arr[i+1]] * arr[arr[i+2]]
		case 99:
			break
		default:
			break
		}
	}
	
	return arr[0]
}

// CalculateP2 calculates the result of part 2
func CalculateP2(s string) int64 {

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			arr := getArray(s)
			if CalculateP1(arr, i, j) == 19690720 {
				return int64(100 * i + j)
			}
		}
	}
	return 0
}

func getArray(s string) []int64 {
	file, _ := os.Open(s)
	scanner := bufio.NewScanner(file)
	var raw string
	for scanner.Scan() {
		raw = scanner.Text()
	}
	
	strArr := strings.Split(raw, ",")
	intArr := make([]int64, len(strArr))
	for index, value := range strArr {
		k, _ := strconv.ParseInt(value, 10, 64)
		intArr[index] = k
	}
	//fmt.Printf("%+v", intArr)
	return intArr
}