package externalmergesort

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func GenerateData(input *string, count int) {
	fmt.Printf("Generating %v random numbers\n", count)
	f, err := os.OpenFile(*input, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)
	for i := 0; i < count; i++ {
		writer.WriteString(fmt.Sprintf("%v\n", rand.Int31()))
	}

	writer.Flush()
}

func ValidateData(input *string) {
	f, err := os.OpenFile(*input, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	prev := math.MinInt32
	line := 1
	for scanner.Scan() {
		cur, _ := strconv.Atoi(scanner.Text())

		if cur < prev {
			fmt.Printf("Error at line %v: %v > %v\n", line, cur, prev)
			break
		}

		prev = cur
		line++
	}
}
