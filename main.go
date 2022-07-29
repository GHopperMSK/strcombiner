package strcombiner

type Combination struct {
	Length int
	Words  map[int]string
}

func Combine(dict []string, start int, index int, desiredLength int) Combination {
	if len(dict) == 0 {
		return Combination{
			Length: -1,
			Words:  make(map[int]string),
		}
	}

	data := make(map[int]string)
	for length := desiredLength; length > 0; length-- {
		res := findCombination(dict, data, start, index, length, 0)
		if len(res) > 0 {
			return Combination{
				Length: length,
				Words:  res,
			}
		}
	}

	return Combination{
		Length: -1,
		Words:  make(map[int]string),
	}
}

func RemoveProcessed(dict *[]string, comb Combination) {
	for _, v := range comb.Words {
		for i := 0; i < len(*dict); i++ {
			if (*dict)[i] == v {
				*dict = append((*dict)[:i], (*dict)[i+1:]...)
				break
			}
		}
	}
}

func findCombination(dict []string, data map[int]string, start int, index int, desiredLength int, currentLength int) map[int]string {
	if currentLength == desiredLength {
		return data
	}

	for i := start; i < len(dict); i++ {
		data[index] = dict[i]
		c := findCombination(dict, data, i+1, index+1, desiredLength, currentLength+len([]rune(dict[i])))
		if len(c) != 0 {
			return c
		}
		delete(data, index)
	}

	return make(map[int]string)
}
