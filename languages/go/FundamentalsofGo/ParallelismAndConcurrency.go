package main

//This is a concurrent code in golang that does not support  parallelism
import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	runtime.GOMAXPROCS(1)
}

func main() {
	var numbers []int
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			numbers = append(numbers, i)
		}(i)
	}
	wg.Wait()
	fmt.Println(numbers)
}
