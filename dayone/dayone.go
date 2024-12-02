package dayone

import (
	"fmt"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/maxiim3/advent-of-code-2024/utils"
)

type InputType string

const (
	Sample InputType = "dayone/sample.txt"
	Input            = "dayone/input.txt"
)

func PartOne(inputType InputType) (int, error) {
	lines, err := utils.Parse(string(inputType))

	if err != nil {
		return 0, err
	}

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

	sort.Ints(left)
	sort.Ints(right)

	length := len(left)
	res := 0

	for i := 0; i < length; i++ {
		diff := float64(right[i] - left[i])
		res += int(math.Abs(diff))
	}

	return res, nil
}

func PartTwo(inputType InputType) {
	fmt.Println(inputType)
	// read file
	// extract lines

	lines, err := utils.Parse(string(inputType))

	if err != nil {
		log.Fatal(err)
	}

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

	// track the number already parsed to prevent from looping over right again
	res := make(map[string]int)
	total := 0
	for i := range left {
		val, ok := res[fmt.Sprint(left[i])]
		if ok {
			// if ok, get the occurence of left in right and multupliy by itself
			total += val * left[i]
			// else loop over right
		} else {
			count := 0
			for j := range right {
				if left[i] == right[j] {
					count++
				}
			}
			// if ok, get the occurence of left in right and multupliy by itself
			if count > 0 {
				res[fmt.Sprint(left[i])] = count
				total += count * left[i]
			}
		}
	}

	fmt.Println(total)
}
