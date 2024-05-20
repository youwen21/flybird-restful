package libutils

import "strconv"

func IsInSliceInt(value int, s []int) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func IsInSliceStr(value string, s []string) bool {
	for _, v := range s {
		if v == value {
			return true
		}
	}

	return false
}

func IntSliceToStringSlice(s []int) []string {
	strS := make([]string, len(s), len(s))
	for idx, v := range s {
		strS[idx] = strconv.Itoa(v)
	}

	return strS
}

func StringSliceToIntSlice(s []string) ([]int, error) {
	intS := make([]int, len(s), len(s))
	var err error
	for idx, v := range s {
		intV, err1 := strconv.Atoi(v)
		if err == nil && err1 != nil {
			err = err1
		}
		intS[idx] = intV
	}

	return intS, err
}
