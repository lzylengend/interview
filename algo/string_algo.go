package algo

import (
	"strings"
)

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

func IsPalindrome(s string) bool {
	s2 := strings.ToLower(s)
	i := 0
	end := len(s) - 1

	for {
		if i > end {
			break
		}

		if (s2[i] < 'a' || s2[i] > 'z') && (s2[i] < 'A' || s2[i] > 'Z') && (s2[i] < '0' || s2[i] > '9') {
			i++
			continue
		}

		if (s2[end] < 'a' || s2[end] > 'z') && (s2[end] < 'A' || s2[end] > 'Z') && (s2[end] < '0' || s2[end] > '9') {
			end--
			continue
		}

		if s2[i] != s2[end] {
			return false
		}

		i++
		end--
	}

	return true
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	res := 0
	start := 0
	end := 1

	for ; end < len(s); end++ {
		tmp := start
		for ; tmp < end; tmp++ {
			if s[tmp] == s[end] {
				start = tmp + 1
			}
		}

		if end-start > res {
			res = end - start
		}
	}

	return res + 1
}
