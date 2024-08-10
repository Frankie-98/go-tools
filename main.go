package main

import (
	"fmt"
	"math/rand"
	"time"
)

func estimatePi(numPoints int) float64 {
	insideCircle := 0
	for i := 0; i < numPoints; i++ {
		x := rand.Float64()
		y := rand.Float64()
		if x*x+y*y <= 1 {
			insideCircle++
		}
	}
	return 4 * float64(insideCircle) / float64(numPoints)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numPoints := 1000000000
	startTime := time.Now()

	approxPi := estimatePi(numPoints)
	fmt.Println("Approximate Pi:", approxPi)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Elapsed time: %.3f seconds\n", elapsedTime.Seconds())
}
