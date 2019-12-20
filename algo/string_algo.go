package algo

func ReverseString(s []byte) []byte {
	count := 0
	if len(s)%2 != 0 {
		count = (len(s) - 1) / 2
	} else {
		count = len(s) / 2
	}

	for i := 0; i < count; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}

	return s
}

func FirstUniqChar(s string) int {
	m := make(map[byte][]int)
	for i, _ := range s {
		m[s[i]] = append(m[s[i]], i)
	}

	min := -1
	for _, v := range m {
		if len(v) > 1 {
			continue
		}

		if len(v) == 1 {
			if min == -1 {
				min = v[0]
			} else if v[0] < min {
				min = v[0]
			}
		}
	}
	return min
}
