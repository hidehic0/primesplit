package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: primesplit <number>")
		return
	}

	var n int64
	n, err := strconv.ParseInt(os.Args[1], 10, 64)

	if err != nil {
		fmt.Println("Usage: primesplit <number>")
		return
	}

	if n == 1 {
		fmt.Println("1: 1")
		return
	}

	res := [][]int64{}
	var i int64 = 2

	for i*i <= n {
		if n%i == 0 {
			res = append(res, []int64{i, 1})
			n /= i
		}

		for n%i == 0 {
			res[len(res)-1][1]++
			n /= i
		}

		i++
	}

	if n > 1 {
		res = append(res, []int64{n, 1})
	}

	var ind = 0

	fmt.Print(n, ": ")

	for ind < len(res) {
		if res[ind][1] == 1 {
			fmt.Print(res[ind][0])
		} else {
			fmt.Print(res[ind][0], "^", res[ind][1])
		}

		if ind+1 != len(res) {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
		ind++
	}
}
