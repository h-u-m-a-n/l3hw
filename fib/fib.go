package fib

func Fibonacci() func() int {
	fib := []int{0, 1}
	i := 0
	return func() int {
		fib = append(fib, fib[i]+fib[i+1])
		i++
		return fib[i-1]
	}
}
