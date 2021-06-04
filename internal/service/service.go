package service

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/sheghun/containerized-webapp/internal/memory"
)

func FindHighestPrime(num int) int {
	originalNum := num
	var wg sync.WaitGroup
	resNum, err := memory.Get(context.Background(), fmt.Sprintf("%d", originalNum))
	if err == nil {
		resNumInt, err := strconv.ParseInt(resNum, 10, 64)
		if err == nil {
			return int(resNumInt)
		}
		log.Printf("error occured parsing redis returned int: %v\n", err)
	}

	if num < 10 {
		if num == 2 || num == 3 || num == 5 || num == 7 {
			return num
		}
		num--
		if num == 2 || num == 3 || num == 5 || num == 7 {
			return num
		}
	}

	primeCh := make(chan int)
	highestPrimeNum := 0

	// Search for a Prime number in first 10 numbers
PrimeNumberSearch:
	// Break the search into batches of 10
	for i := num; i >= num-10; i-- {
		// Break it into go routines
		wg.Add(1)
		go func(j int, primeCh chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			for x := 2; x <= j; x++ {
				if j%x == 0 && j != x {
					return
				}
				if j%x == 0 && j == x {
					primeCh <- x
					return
				}
			}
		}(i, primeCh, &wg)
	}

	go func() {
		for primeNums := range primeCh {
			if primeNums > highestPrimeNum {
				highestPrimeNum = primeNums
			}
		}
	}()

	wg.Wait()
	// Wait for highestPrimeNumber to get the last int from the channel
	time.Sleep(100 * time.Millisecond)

	if highestPrimeNum == 0 {
		num = num - 10
		// If a prime number was not gotten from the previous 10 batch run another 10 batch
		goto PrimeNumberSearch
	}

	if err := memory.Set(context.Background(), fmt.Sprintf("%d", originalNum), fmt.Sprintf("%d", highestPrimeNum)); err != nil {
		log.Printf("error occurred trying to save prime number to memory: %v\n", err)
	}

	return highestPrimeNum
}
