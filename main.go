package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var t int
	//var err error
	_, _ = fmt.Fscan(in, &t)
	// if err != nil {
	// 	fmt.Print("ошибка")
	// 	return
	// }
	var n, p int
	for i := 1; i <= t; i++ {
		//var er error
		_, _ = fmt.Fscan(in, &n, &p)
		// if er != nil {
		// 	fmt.Println("ошибка ввода")
		// 	return
		// }
		fmt.Fprintf(out, "%.2f \n", roundSum(n, p, in))
	}
}

func roundSum(n, p int, in *bufio.Reader) float64 {
	var resultInt int
	var resultFloat, result float64
	for i := 0; i < n; i++ {
		func(p *int) {
			var number int
			//var err error
			_, _ = fmt.Fscan(in, &number)
			// if err != nil {
			// 	fmt.Println("не целое число!")
			// 	return
			// }
			resultInt = int(float64(number) * float64(*p) / 100)
			resultFloat = float64(number) * float64(*p) / 100
			result += resultFloat - float64(resultInt)
		}(&p)
	}
	return result
}
