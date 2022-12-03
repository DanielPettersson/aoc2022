package common

func SetOfChars(input string) map[rune]interface{} {
	m := make(map[rune]interface{})

	for _, c := range input {
		m[c] = nil
	}

	return m
}

func Intersection[T comparable, K interface{}](input1 map[T]K, input2 map[T]K) map[T]K {
	ret := make(map[T]K)

	for k, v := range input1 {
		if _, present := input2[k]; present {
			ret[k] = v
		}
	}

	return ret
}
