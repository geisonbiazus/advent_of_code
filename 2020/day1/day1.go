package day1

func FindEntriesThatSumTo2020(numberOfEntries int, input []int) []int {
	entries, _ := findEntriesThatSumToAmount(numberOfEntries, 2020, input, []int{}, 0)
	return entries
}

func findEntriesThatSumToAmount(numberOfEntries, amountToSum int, input, output []int, fromIndex int) ([]int, bool) {
	if numberOfEntries == 0 {
		return output, amountToSum == 0
	}

	for i := fromIndex; i < len(input); i++ {
		v := input[i]
		if v <= amountToSum {
			out := append([]int{}, output...)
			out = append(out, v)

			if entries, ok := findEntriesThatSumToAmount(numberOfEntries-1, amountToSum-v, input, out, i+1); ok {
				return entries, ok
			}
		}
	}

	return output, false
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
