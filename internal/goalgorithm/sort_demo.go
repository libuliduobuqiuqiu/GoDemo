package goalgorithm

// 冒泡排序：比较相邻的元素，将最大的元素冒泡到列表的末尾
func BubbleSort(unSorted []int) []int {
	for i := 0; i < len(unSorted); i++ {
		for j := 0; j < len(unSorted)-1-i; j++ {
			if unSorted[j] > unSorted[j+1] {
				unSorted[j], unSorted[j+1] = unSorted[j+1], unSorted[j]
			}
		}
	}
	return unSorted
}

// 插入排序: 从后向前扫描，当遇到大于基准元素的就向右移动，循环已排序的部分，当找到小于基准元素时跳出循环
func InsertionSort(unSorted []int) []int {

	// 从第一个元素开始
	for i := 1; i < len(unSorted); i++ {
		var j int
		tmp := unSorted[i]

		// 从后向前扫描，将大于tmp的元素向右移动
		for j = i - 1; j >= 0 && unSorted[j] > tmp; {
			unSorted[j+1] = unSorted[j]
			j -= 1
		}

		unSorted[j+1] = tmp
	}
	return unSorted
}

// 归并排序（分治法）：将切片分成两部分，直到切片的长度为1，然后再把有序的切片合并
func MergeSort(unSorted []int) []int {
	if len(unSorted) <= 1 {
		return unSorted
	}

	half := len(unSorted) / 2
	left := MergeSort(unSorted[:half])
	right := MergeSort(unSorted[half:])

	return Merge(left, right)
}

func Merge(left, right []int) []int {
	var sorted []int
	var i, j int

	// 通过对比左右切片的元素添加到新的已排序的切片
	for i, j = 0, 0; i < len(left) && j < len(right); {

		if left[i] < right[j] {
			sorted = append(sorted, left[i])
			i += 1
		} else {
			sorted = append(sorted, right[j])
			j += 1
		}
	}

	if i < len(left) {
		sorted = append(sorted, left[i:]...)
	}

	if j < len(right) {
		sorted = append(sorted, right[j:]...)
	}
	return sorted
}

// 快速排序: 确定基准元素，分区，递归排序分区里面的元素
func QuickSort(unSorted []int) []int {

	if len(unSorted) <= 1 {
		return unSorted
	}

	tmp := unSorted[0]
	var left, right, sorted []int
	for i := 1; i < len(unSorted); i++ {
		if unSorted[i] > tmp {
			right = append(right, unSorted[i])
		} else {
			left = append(left, unSorted[i])
		}
	}

	sorted = append(sorted, QuickSort(left)...)
	sorted = append(sorted, tmp)
	sorted = append(sorted, QuickSort(right)...)
	return sorted
}

// 选择排序: 从未排序部分中确定最小元素的位置，将最小元素放到已排序的末尾
func SelectionSort(unSorted []int) []int {

	if len(unSorted) <= 1 {
		return unSorted
	}

	for i := 0; i < len(unSorted)-1; i++ {
		exchange := i

		for j := i + 1; j < len(unSorted); j++ {
			if unSorted[exchange] > unSorted[j] {
				exchange = j
			}
		}

		unSorted[i], unSorted[exchange] = unSorted[exchange], unSorted[i]
	}
	return unSorted
}
