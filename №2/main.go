package main

import (
	"fmt"
	"sync"
)

func main() {
	var t int
	var err error
	_, err = fmt.Scan(&t)
	if err != nil {
		fmt.Print("ошибка")
		return
	}
	var n, p int
	for i := 1; i <= t; i++ {
		var er error
		_, er = fmt.Scan(&n, &p)
		if er != nil {
			fmt.Println("ошибка ввода")
			return
		}
		fmt.Printf("%.2f \n", roundSum(n, p))
	}
}

func roundSum(n, p int) float64 {
	var resultFloat, result float64
	var resultInt int
	mu := new(sync.Mutex)
	for i := 0; i < n; i++ {
		var number int
		var err error
		_, err = fmt.Scan(&number)
		if err != nil {
			panic("не целое число!")
		}
		go func(p *int) {
			mu.Lock()
			resultInt = int(float64(number) * float64(*p) / 100)
			resultFloat = float64(number) * float64(*p) / 100
			result += resultFloat - float64(resultInt)
			mu.Unlock()
		}(&p)
	}
	return result
}
