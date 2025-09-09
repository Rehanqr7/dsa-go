package patterns

/*
Binary Search on Answer Pattern
Example: Capacity To Ship Packages Within D Days

Given weights and days, find the minimal ship capacity to ship within given days.

Time Complexity: O(n log sum(weights))
Space Complexity: O(1)
*/

func ShipWithinDays(weights []int, days int) int {
	low, high := 0, 0
	for _, w := range weights {
		if w > low { low = w }
		high += w
	}
	for low < high {
		mid := (low + high) / 2
		if canShip(weights, days, mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func canShip(weights []int, days int, cap int) bool {
	need := 1
	sum := 0
	for _, w := range weights {
		if sum + w > cap {
			need++
			sum = 0
		}
		sum += w
		if need > days { return false }
	}
	return true
}
