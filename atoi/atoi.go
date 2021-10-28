package atoi

import "errors"

func Atoi(s string) (int, error) {
	if len(s) == 0 {
		return 0, errors.New("string is empty")
	}
	negative := s[0] == '-'
	if negative {
		s = s[1:]
	}
	res := 0
	for _, v := range s {
		if !(v >= '0' && v <= '9') {
			return 0, errors.New("string contains invalid symbol")
		}
		res = res*10 + int(v-'0')
	}
	if negative {
		res *= -1
	}
	return res, nil
}
