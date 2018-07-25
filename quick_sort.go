package main

func quickSort(list []int) {
	if len(list) <= 1 {
		return
	}

	val := list[0]
	start, end := 0, len(list)-1
	for i := 1; i <= end; {
		if list[i] > val {
			list[i], list[end] = list[end], list[i]
			end--
		} else {
			list[i], list[start] = list[start], list[i]
			start++
			i++
		}
	}

	quickSort(list[:start])
	quickSort(list[start+1:])
}
