package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
)

func exit() {
	fmt.Print("Usage: ")
	c := color.New(color.FgRed, color.Bold)
	c.Println("primesplit <number of 0 more>")
	os.Exit(1)
}

func main() {
	c := color.New(color.Bold, color.FgBlue)
	if len(os.Args) != 2 {
		exit()
	}

	var n int64
	n, err := strconv.ParseInt(os.Args[1], 10, 64)

	if err != nil {
		exit()
	}

	if n <= 0 {
		exit()
	}

	if n == 1 {
		fmt.Print("1: ")
		c.Println("1")
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
			c.Print(res[ind][0])
		} else {
			c.Print(res[ind][0], "^", res[ind][1])
		}

		if ind+1 != len(res) {
			fmt.Print(" ")
		} else {
			fmt.Println()
		}
		ind++
	}
}
