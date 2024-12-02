package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rawFile, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(rawFile), "\n")
	lastIndex := len(lines) - 1
	lines = lines[:lastIndex]

	var left, right []int

	for _, line := range lines {
		f := strings.Fields(line)

		l, err := strconv.Atoi(f[0])

		if err != nil {
			log.Fatal(err)
		}

		r, err := strconv.Atoi(f[1])

		if err != nil {
			log.Fatal(err)
		}

		left = append(left, l)
		right = append(right, r)
	}

	length := len(left)
	sort.Ints(left)
	sort.Ints(right)

	diff := 0
	for i := 0; i < length; i++ {
		d := float64(right[i] - left[i])
		diff += int(math.Abs(d))
	}

	fmt.Println(diff)
}
