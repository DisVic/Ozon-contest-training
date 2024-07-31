package main

import (
	"fmt"
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
	var resultInt int
	var resultFloat float64
	for i := 0; i < n; i++ {
		func(p *int) {
			var number int
			var err error
			_, err = fmt.Scan(&number)
			if err != nil {
				fmt.Println("не целое число!")
				return
			}
			resultInt += int(float64(number) * float64(*p) / 100)

			resultFloat += float64(number) * float64(*p) / 100

		}(&p)
		fmt.Println(resultFloat)
	}
	return resultFloat - float64(resultInt)
}
