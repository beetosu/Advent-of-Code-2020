package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	content, _ := ioutil.ReadFile("input")

	numList := strings.Split(string(content), "\n")

	possibilities := make(map[int]map[int]int)

	start := time.Now()

	target := 2020

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
