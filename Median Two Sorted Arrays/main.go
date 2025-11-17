package main

func main() {

	arr1 := []int{1, 2}
	arr2 := []int{3, 4}

	median := findMedianSortedArrays(arr1, arr2)
	println("Median is:", median)

}

func findMedianSortedArrays(arr1, arr2 []int) float64 {

	m := len(arr1)
	n := len(arr2)
	totalLength := m + n
	merged := make([]int, 0, totalLength)
	i, j := 0, 0
	for i < m && j < n {
		if arr1[i] < arr2[j] {
			merged = append(merged, arr1[i])
			i++
		} else {
			merged = append(merged, arr2[j])
			j++
		}
	}

	for i < m {
		merged = append(merged, arr1[i])
		i++
	}

	for j < n {
		merged = append(merged, arr2[j])
		j++
	}

	if totalLength%2 == 0 {
		mid1 := totalLength/2 - 1
		mid2 := totalLength / 2
		return float64(merged[mid1]+merged[mid2]) / 2.0
	} else {
		mid := totalLength / 2
		return float64(merged[mid])
	}
}
