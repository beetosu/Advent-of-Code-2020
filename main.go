package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	content, _ := ioutil.ReadFile("day 1/input")

	numList := strings.Split(string(content), "\n")

	possibilities := make(map[int]map[int]int)

	start := time.Now()

	target := 2020

	//the solution to the first part was basically the inner for loop, so O(n)
	//but I modified it to find 3 numbers, which ended up making it O(n^2).
	//there's definately some redundancy that could be smoothed out, espically if i sorted the list beforehand
	//but atm it isn't *too* bad
	//(the runtime with the given output is currently 1 millisecond)
	for outerIdx, outerElem := range numList {
		firstKey, firstConvIssue := strconv.Atoi(outerElem)
		if firstConvIssue == nil {
			possibilities[firstKey] = make(map[int]int)
			for innerIdx, innerElem := range numList {
				secondKey, secondConvIssue := strconv.Atoi(innerElem)
				thirdKey, found := possibilities[firstKey][target-firstKey-secondKey]
				if found && innerIdx != outerIdx {
					fmt.Println(firstKey, secondKey, thirdKey, firstKey*secondKey*thirdKey)
					elapsed := time.Since(start)
					fmt.Println("finding this took", elapsed.Microseconds(), "microseconds.")
					return
				} else if secondConvIssue == nil {
					possibilities[firstKey][secondKey] = secondKey
				}
			}
		}
	}
}
