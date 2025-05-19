package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	n, _ = strconv.Atoi(os.Args[1])
	res := [][]int{}
	var i int = 2

	for n != 1 {
		if n%i == 0 {
			res = append(res, []int{i, 1})
			n /= i
		}

		for n%i == 0 {
			res[len(res)-1][1]++
			n /= i
		}

		i++
	}

	i = 0

	fmt.Print(n, ": ")

	for i < len(res) {
		if res[i][1] == 1 {
			fmt.Print(res[i][0])
		} else {
			fmt.Print(res[i][0], "^", res[i][1])
		}

		if i+1 != len(res) {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
		i++
	}
}
