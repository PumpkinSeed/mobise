package mobise

func Search(arr []int, low, high int) int {
	if low > high {
		return -1
	}

	if arr[low] > arr[high] {
		var mid = (low + high) / 2
		if arr[mid] > arr[high] {
			return Search(arr, mid+1, high)
		}
		return Search(arr, low, mid)
	}
	return low
}
