package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: fig NUM")
		os.Exit(1)
	}

	nums, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("invalid parameter, %s\n", os.Args[1])
		os.Exit(1)
	}

	if nums <= 1 {
		fmt.Println(os.Args[1])
		return
	}

	a := big.NewInt(0)
	b := big.NewInt(1)
	res := big.NewInt(0)
	for ; nums > 1; nums-- {
		tmp := res
		res.Add(a, b)

		a = b
		b = tmp
	}

	fmt.Println(res.String())
	return
}
