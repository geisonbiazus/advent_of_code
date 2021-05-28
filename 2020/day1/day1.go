package day1

func FindEntriesThatSumTo2020(numberOfEntries int, input []int) []int {
	if numberOfEntries == 1 {
		for _, entry := range input {
			if entry == 2020 {
				return []int{entry}
			}
		}
	}

	if numberOfEntries == 2 {
		for _, v1 := range input {
			for _, v2 := range input {
				if v1+v2 == 2020 {
					return []int{v1, v2}
				}
			}
		}
	}

	for _, v1 := range input {
		for _, v2 := range input {
			for _, v3 := range input {
				if v1+v2+v3 == 2020 {
					return []int{v1, v2, v3}
				}
			}
		}
	}

	return []int{}
}

func SolvePuzzle(numberOfEntries int, input []int) int {
	entries := FindEntriesThatSumTo2020(numberOfEntries, input)

	result := 0

	if len(entries) > 0 {
		result = 1
	}

	for _, v := range entries {
		result *= v
	}

	return result
}
