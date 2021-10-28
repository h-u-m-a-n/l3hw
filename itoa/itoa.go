package itoa

import "github.com/h-u-m-a-n/l3hw/revstr"

func Itoa(i int) string {
	if i == 0 {
		return "0"
	}
	res := make([]byte, 0)
	negative := false
	if i < 0 {
		i *= -1
		negative = true
	}
	for i > 0 {
		r := byte(i % 10)
		res = append(res, byte(r+'0'))
		i /= 10
	}
	if negative {
		res = append(res, '-')
	}
	return revstr.Reverse(string(res))
}
