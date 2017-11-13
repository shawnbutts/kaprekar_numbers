package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func checkKap(intChan chan int64, start time.Time) {

	for {
		select {
		case num := <-intChan:

			sq := math.Pow(float64(num), 2)

			r := strconv.FormatFloat(sq, 'f', 0, 64)

			sp := strings.Split(r, "")

			for i := 1; i < len(sp); i++ {
				fs := strings.Join(sp[0:i], "")
				ls := strings.Join(sp[i:], "")

				f, _ := strconv.ParseInt(fs, 0, 64)
				l, _ := strconv.ParseInt(ls, 0, 64)

				if (f+l) == num && f > 0 && l > 0 {
					t := time.Now()

					fmt.Printf("%v %v (%v, %v) - time: %v\n", num, int64(sq), f, l, t.Sub(start))
				}
			}
		}
	}
}

func main() {
	var i int64

	concurrency := 6
	intChan := make(chan int64, concurrency*2)
	start := time.Now()
	fmt.Printf("%T", start)

	for g := 0; g < concurrency; g++ {
		go checkKap(intChan, start)
	}

	for {
		i++
		intChan <- i
	}

}
