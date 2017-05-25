package utils

import (
	"fmt"
	"os"
	"math/rand"
)

func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error ", err.Error())
		panic(err)
	}
}

func RandInt(min,max int) int {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	return rand.Intn(max-min)+min
}