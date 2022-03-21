package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		fmt.Println("执行前时间（ms）：", time.Now().UnixNano()/1e6)
		line := input.Text()
		for loop := 0; loop < 100; loop++ {
			ans := make([]int, 0)
			dp := make([]int, len(line))
			var charMap [128]int
			maxLen := 1
			dp[0] = 1
			for i := 0; i < len(line); i++ {
				charMap[i] = -1
			}
			charMap[line[0]] = 0
			for i := 1; i < len(line); i++ {
				dp[i] = int(math.Min(float64(dp[i-1]+1), float64(i-charMap[line[i]])))
				maxLen = int(math.Max(float64(maxLen), float64(dp[i])))
				charMap[line[i]] = i

			}
			for i, v := range dp {
				if v == maxLen {
					ans = append(ans, i)
				}
			}
			fmt.Println(maxLen, ",", ans)
		}
		fmt.Println("执行后时间（ms）：", time.Now().UnixNano()/1e6)
	}

}
