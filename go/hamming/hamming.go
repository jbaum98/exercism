package hamming

import "fmt"

const testVersion = 5

// chunkSize is the maximum size of a slice to be handled serially. Strings
// larger than this will be broken up into chunks of chunkSize or smaller and
// processed in parallel.
const chunkSize = 5000

// Distance computes the Hamming difference between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, fmt.Errorf("DNA strands were not of equal length")
	}

	if len(a) < 2*chunkSize {
		return distance(a, b), nil
	}

	return parDist(a, b, chunkSize), nil
}

// parDist omputes the Hamming difference in parallel by splitting the strings
// in. chunks of size chunkSize or smaller.
func parDist(a, b string, chunkSize int) int {
	var dif int

	l := len(a)

	nChunks := l / chunkSize
	if l%chunkSize != 0 {
		nChunks++
	}

	ch := make(chan int)

	var chunkStart, chunkEnd int
	for chunkStart, chunkEnd = 0, chunkSize; chunkEnd < l; chunkStart, chunkEnd = chunkEnd, chunkEnd+chunkSize {
		aChunk, bChunk := a[chunkStart:chunkEnd], b[chunkStart:chunkEnd]
		go sendDist(aChunk, bChunk, ch)
	}
	go sendDist(a[chunkStart:], b[chunkStart:], ch)

	for waiting := nChunks; waiting > 0; waiting-- {
		dif += <-ch
	}
	return dif
}

func sendDist(aChunk, bChunk string, ch chan int) {
	ch <- distance(aChunk, bChunk)
}

// distance computes the Hamming difference serially.
func distance(a, b string) int {
	var dif int
	for i := range a {
		if a[i] != b[i] {
			dif++
		}
	}
	return dif
}
