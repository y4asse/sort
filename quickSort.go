package sort

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]

	left := []int{}
	right := []int{}

	for _, v := range arr[1:] {
		if v > pivot {
			right = append(right, v)
		} else {
			left = append(left, v)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	var ret []int
	ret = append(left, pivot)
	ret = append(ret, right...)

	return ret
}
