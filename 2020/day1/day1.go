package day1

func FindEntries(input []int) []int {
	for _, v1 := range input {
		for _, v2 := range input {
			if v1+v2 == 2020 {
				return []int{v1, v2}
			}
		}
	}

	return []int{}
}

func SolvePuzzle(input []int) int {
	entries := FindEntries(input)

	if len(entries) == 2 {
		return entries[0] * entries[1]
	}

	return 0
}
