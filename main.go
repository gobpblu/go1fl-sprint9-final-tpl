package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	var numbers = make([]int, size)

	for i := range size {
		numbers[i] = rand.Int() % 1000000
	}
	// fmt.Println(numbers)

	return numbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// fmt.Println(data)
	if len(data) <= 0 {
		return 0
	}

	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) < 8 {
		return maximum(data)
	}

	chunksMaxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup
	var mu sync.Mutex
	chunkSize := int(math.Floor(float64(len(data)) / CHUNKS))
	fmt.Println("CHUNK SIZE:", chunkSize, float64(SIZE)/8)

	for i := range CHUNKS {
		wg.Add(1)
		go func() {
			beginIndex := i * chunkSize
			endIndex := beginIndex + chunkSize

			var chunk []int
			if i == CHUNKS-1 {
				chunk = data[beginIndex:]
			} else {
				chunk = data[beginIndex:endIndex]
			}

			chunkMax := maximum(chunk)

			mu.Lock()
			chunksMaxValues[i] = chunkMax
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return maximum(chunksMaxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	numbers := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(now).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	now = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(now).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
