package remove_element

func RemoveElement(num []int, val int) int {
	res := 0

	for i := 0; i < len(num); i++ {
		if num[i] != val {
			num[res] = num[i]
			res++
		}
	}

	return res
}
