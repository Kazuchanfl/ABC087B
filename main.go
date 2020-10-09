package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b, c, x int
	var result int = 0

	fmt.Fscan(r, &a)
	fmt.Fscan(r, &b)
	fmt.Fscan(r, &c)
	fmt.Fscan(r, &x)
	// 500円玉を0からAまで指定する
	for i := 0; i < a; i++ {
		// 100円玉を0からBまで指定する
		for j := 0; j < b; j++ {
			// 50円玉を0からCまで指定する
			for k := 0; k < c; k++ {
				// 500 * a + 100 * b + 50 * c == x なら result ++
			}
		}
	}
	fmt.Fprint(w, result)
}
