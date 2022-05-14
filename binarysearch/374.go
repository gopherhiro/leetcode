package binarysearch

// 374. Guess Number Higher or Lower
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		answer := guess(mid)
		if answer == 1 {
			left = mid + 1
		} else if answer == -1 {
			right = mid - 1
		} else if answer == 0 {
			return mid
		}
	}
	return 0
}

func guess(num int) int {
	return 0
}
