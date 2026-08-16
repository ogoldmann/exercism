package binarysearch
func SearchInts(list []int, key int) int {
	low, high := 0, len(list)-1
	for low <= high {
		mid := low + (high-low) / 2
		if list[mid] < key {
			low = mid + 1
		} else if list[mid] > key {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}