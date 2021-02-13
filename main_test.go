package main

import (
	"sync"
	"testing"
	"time"
)

func TestAppend(t *testing.T) {
	input := []string{"63584162", "64441813", "64442336", "64448862", "64449735", "64449952", "64450126", "64430719", "64447862"}
	output := make([]string, 0, len(input))
	//output := []string{}
	//output := make([]string, len(input))

	wg := sync.WaitGroup{}
	wg.Add(len(input))
	for _, s := range input {
		go func() {
			defer wg.Done()
			output = append(output, s)
		}()
		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
	println("inputLen:", len(input), "   OutputLength:", len(output))

	for _, s := range output {
		print(s)
		print(",")
	}
}
