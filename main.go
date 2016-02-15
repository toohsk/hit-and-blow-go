package main

import (
	"math/rand"
	"time"
)

func getThreeTargetNum() []int {

	targetNumAry := make([]int, 3)
	rand.Seed(time.Now().UnixNano())

	for i := range targetNumAry {
		val := rand.Intn(9)
		print(val)
		targetNumAry[i] = val
	}

	return targetNumAry

}

func main() {

	getThreeTargetNum()

}
