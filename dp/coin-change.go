package dp

import "math"

/*
Coin Change (Medium)
Given coins and an amount, return the fewest number of coins needed to make up that amount, or -1 if impossible.

Time Complexity: O(amount * len(coins))
Space Complexity: O(amount)
*/

func CoinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ { dp[i] = math.MaxInt/4 }
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if c <= i && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] >= math.MaxInt/8 { return -1 }
	return dp[amount]
}
