package binarysearch

func SearchInts(slice []int, key int) int {
	return search(slice, key, 0)
}

func search(slice []int, key, index int) int {
	if len(slice) == 0 {
		return -1
	}
	half := len(slice) / 2
	switch {
	case slice[half] < key:
		return search(slice[half+1:], key, index+half+1)
	case slice[half] > key:
		return search(slice[:half], key, index)
	default:
		return index + half
	}
}