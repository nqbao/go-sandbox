package externalmergesort

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"

	sys "github.com/nqbao/learn-go/sys"
)

func ExternalMergeSort(input *string, chunkSize float64) {
	f, err := os.OpenFile(*input, os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	sliceSize := int(chunkSize*1024*1024) / 4
	chunk := make([]int32, sliceSize)

	log.Printf("Slice size %v", sliceSize)

	scanner := bufio.NewScanner(f)

	chunkIndex := 0
	count := 0
	for scanner.Scan() {
		parsed, _ := strconv.Atoi(scanner.Text())
		chunk[count] = int32(parsed)
		count = count + 1

		// finish 1 chunk, write it down
		if count >= sliceSize {
			log.Printf("Sorting and writing chunk %v", chunkIndex)
			quicksort(chunk, count)
			runtime.GC()
			writeChunk(chunk, count, chunkIndex)
			runtime.GC()
			sys.PrintMemUsage()

			// reset counter and increase chunk index
			count = 0
			chunkIndex = chunkIndex + 1
		}
	}

	// don't forget the final chunk
	if count > 0 {
		log.Printf("Sorting and writing chunk %v", chunkIndex)
		quicksort(chunk, count)
		runtime.GC()
		writeChunk(chunk, count, chunkIndex)
		runtime.GC()
		sys.PrintMemUsage()
	}

	log.Printf("Merging chunks ...")
	mergeChunks(chunkIndex + 1)
}

func writeChunk(input []int32, count int, chunkIndex int) {
	os.MkdirAll("chunk", 0766)

	chunkName := fmt.Sprintf("chunk/%v", chunkIndex)
	f, err := os.OpenFile(chunkName, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	writer := bufio.NewWriter(f)

	for i := 0; i < count; i++ {
		writer.WriteString(fmt.Sprintf("%v\n", input[i]))
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func mergeChunks(chunks int) {
	f, err := os.OpenFile("chunk/final", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	writer := bufio.NewWriter(f)

	readers := make([]*bufio.Scanner, chunks)
	for i := 0; i < chunks; i++ {
		f2, err := os.Open(fmt.Sprintf("chunk/%v", i))

		if err != nil {
			log.Fatal(err)
		}

		defer f2.Close()
		readers[i] = bufio.NewScanner(f2)
	}

	top := make([]*int, chunks)
	for i := 0; i < chunks; i++ {
		if top[i] == nil && readers[i].Scan() {
			tmp, _ := strconv.Atoi(readers[i].Text())
			top[i] = &tmp
		}
	}

	for {
		// find current min
		cur := int(math.MaxInt32)
		j := -1
		for i := 0; i < chunks; i++ {
			if top[i] != nil && cur >= *top[i] {
				// fmt.Printf("Selecting %v: %v %v\n", i, cur, *top[i])
				cur = *top[i]
				j = i
			}
		}

		// done
		if j == -1 {
			break
		}

		// fmt.Printf("Reading from chunk %v: %v\n", j, cur)
		writer.WriteString(fmt.Sprintf("%v\n", cur))

		if readers[j].Scan() {
			tmp, _ := strconv.Atoi(readers[j].Text())
			top[j] = &tmp
		} else {
			top[j] = nil
		}
	}

	writer.Flush()
}
