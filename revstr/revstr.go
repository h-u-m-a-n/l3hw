package revstr

func Reverse(s string) string {
	arr := []rune(s)
	n := len(arr)
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-i-1] = arr[n-i-1], arr[i]
	}
	return string(arr)
}
