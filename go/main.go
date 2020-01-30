package main

import (
	"fmt"
	"os"
	"runtime"
	"sort"
	"sync"
	"time"
)

const corpusRoute = "./corpus"
const xRoute = "./corpus-x"

var numCPU = runtime.NumCPU()
var lock = sync.RWMutex{}

var xs []string
var corpus []Tuple
var err error

func indexOf(arr []string, element string) int {
	for index, value := range arr {
		if value == element {
			return index
		}
	}
	return -1
}

func uniqueY(tuples []Tuple) []uint64 {
	keys := make(map[uint64]bool)
	uniq := make([]uint64, 0)
	for _, tuple := range tuples {
		if _, isset := keys[tuple.Y]; !isset {
			uniq = append(uniq, tuple.Y)
			keys[tuple.Y] = true
		}
	}
	return uniq
}

func compare(a, b []string) bool {
	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func init() {
	fmt.Printf("=> Importing the datasets.\n")
	xs = LoadX(xRoute)
	corpus, err = LoadCorpus(corpusRoute)
	if err != nil {
		fmt.Printf("Error loading the corpus.\n")
		os.Exit(1)
	}
}

func main() {
	fmt.Printf("=> Placing the elements in the intermediate table.\n")
	startSort := time.Now()

	// Sorting the keys using the stdlib.
	sort.Strings(xs)

	// We calculate the amount of Y. Very much worth the computational cost.
	ys := uniqueY(corpus)

	// Size of the keys space is the same as the x's
	table := make(map[string]map[uint64]string, len(xs))
	for _, x := range xs {
		table[x] = make(map[uint64]string, len(ys))
	}

	for _, tuple := range corpus {
		table[tuple.X][tuple.Y] = tuple.Value
	}

	diffSort := time.Since(startSort)
	fmt.Printf("=> Sorting finished.\n")

	fmt.Printf("=> Starting transpose of the table to its final form.\n")
	startTrans := time.Now()

	output := make([][]string, len(ys)+1)
	for i := range output {
		output[i] = make([]string, len(xs))
	}

	copy(output[0], xs)

	cTrans := make(chan bool, numCPU)
	qxs := len(xs) / numCPU
	transpose := func(keys []string, out [][]string, channel chan bool) {
		for _, x := range keys {
			yValue := table[x]
			xIndex := indexOf(xs, x)
			if xIndex < 0 {
				panic(fmt.Errorf("Invalid value: %s", x))
			}

			for y, value := range yValue {
				out[y+1][xIndex] = value
			}
		}
		channel <- true
	}

	for i := 0; i < numCPU; i++ {
		keys := xs[i*qxs : (i+1)*qxs]
		go transpose(keys, output, cTrans)
	}

	for i := 0; i < numCPU; i++ {
		<-cTrans
	}

	diffTrans := time.Since(startTrans)
	fmt.Printf("=> Transpose finished!\n")

	fmt.Printf("=> Ensuring the validity of the table.\n")
	startCheck := time.Now()
	header, tester := output[0], output[1:]
	if !compare(header, xs) {
		panic(fmt.Errorf("=> Transpose error: header => %v xs => %v", header, xs))
	}

	cCheck := make(chan bool)
	qys := len(ys) / numCPU
	check := func(y int, end int, test [][]string, channel chan bool) {
		for ; y < end; y++ {
			xValues := test[y]
			for xIndex, value := range xValues {
				x := xs[xIndex]
				computed := fmt.Sprintf("x=%s,y=%d", x, y)
				if value != computed {
					panic(fmt.Errorf("=> Transpose error: value => %s computed => %s", value, computed))
				}
			}
		}
		channel <- true
	}

	for i := 0; i < numCPU; i++ {
		go check(i*qys, (i+1)*qys, tester, cCheck)
	}

	for i := 0; i < numCPU; i++ {
		<-cCheck
	}

	diffCheck := time.Since(startCheck)
	diffTotal := time.Since(startSort)
	fmt.Printf("\n")
	fmt.Printf("Total execution time: %f\n", diffTotal.Seconds())
	fmt.Printf("---\n")
	fmt.Printf("Sorting execution time: %f\n", diffSort.Seconds())
	fmt.Printf("Transpose execution time: %f\n", diffTrans.Seconds())
	fmt.Printf("Check execution time: %f\n", diffCheck.Seconds())
}
