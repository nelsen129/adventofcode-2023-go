package algorithm

func quickSortPartition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	i := lo

	for j := lo; j < hi; j += 1 {
		if nums[j] <= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i += 1
		}
	}

	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

func quickSortInPlace(nums []int, lo, hi int) {
	if lo >= hi || lo < 0 {
		return
	}

	p := quickSortPartition(nums, lo, hi)

	quickSortInPlace(nums, lo, p-1)
	quickSortInPlace(nums, p+1, hi)
}

func QuickSort(nums []int) []int {
	n := make([]int, len(nums))
	copy(n, nums)

	quickSortInPlace(n, 0, len(n)-1)
	return n
}

func FindOverlapInSortedLists(nums1, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	overlap := []int{}

	p1 := 0
	p2 := 0

	for p1 < len(nums1) && p2 < len(nums2) {
		if nums1[p1] < nums2[p2] {
			p1 += 1
			continue
		}
		if nums1[p1] > nums2[p2] {
			p2 += 1
			continue
		}
		target := nums1[p1]
		p1count := 0
		for p1 < len(nums1) && nums1[p1] == target {
			p1count += 1
			p1 += 1
		}
		p2count := 0
		for p2 < len(nums2) && nums2[p2] == target {
			p2count += 1
			p2 += 1
		}
		o := make([]int, p1count*p2count)
		for i := range p1count * p2count {
			o[i] = target
		}
		overlap = append(overlap, o...)
	}

	return overlap
}
