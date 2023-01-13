package main

import (
	"fmt"
	"time"
)

var hashList []string

func main() {
	for {
		h := createHash(createRandomString(5))

		if ok, v := hasThreeZeroInHash(h); ok {
			sortingAndAppendHashToSlice(v)
			fmt.Printf("%v", hashList)
		}

		<-time.Tick(1 * time.Millisecond)
	}

}

func createRandomString(strSize int) <-chan string {
	str := make(chan string)
	go func() {
		defer close(str)
		str <- randomString(strSize)
	}()

	return str
}

func createHash(str <-chan string) <-chan hash {
	h := make(chan hash)
	go func() {
		defer close(h)
		h <- newHash(<-str)
	}()

	return h
}

func hasThreeZeroInHash(hash <-chan hash) (bool, <-chan string) {
	validHash, ok := make(chan string), false
	go func() {
		defer close(validHash)
		h := <-hash

		if h[len(h)-1] == 0 && h[len(h)-2] == 0 && h[len(h)-3] == 0 {
			validHash <- h.String()
			ok = true
		}
	}()

	return ok, validHash
}

func sortingAndAppendHashToSlice(h <-chan string) {
	if len(hashList) == 0 {
		hashList = append(hashList, <-h)
		return
	}

	hashList = append(hashList, <-h)
	nums := make([]int, 0)

	for _, s := range hashList {
		nums = append(nums, <-calcSumHash(s))
	}

	go func(nums []int) {
		n := len(hashList)
		for i := 0; i < i-1; i++ {
			for j := 0; j < n-i-1; j++ {
				if nums[j] < nums[j+1] {
					nums[j], nums[j+1] = nums[j+1], nums[j]
					hashList[j], hashList[j+1] = hashList[j+1], hashList[j]
				}
			}
		}
	}(nums)
}

func calcSumHash(h string) <-chan int {
	sum := make(chan int)

	go func(currentHash string) {
		defer close(sum)

		var (
			num, s int
		)

		for _, char := range currentHash {
			if char >= '0' && char <= '9' {
				num = num*10 + int(char-'0')
			} else {
				s += num
				num = 0
			}
		}
		s += num

		sum <- s

	}(h)

	return sum
}
