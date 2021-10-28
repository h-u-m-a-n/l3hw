package main

import (
	"fmt"
	"github.com/h-u-m-a-n/l3hw/atoi"
	"github.com/h-u-m-a-n/l3hw/fib"
	"github.com/h-u-m-a-n/l3hw/itoa"
	"github.com/h-u-m-a-n/l3hw/revstr"
	"github.com/h-u-m-a-n/l3hw/runebyindex"
	"github.com/h-u-m-a-n/l3hw/sort"
	"log"
)

func main() {
	// task 2
	fmt.Println(itoa.Itoa(-10))
	// task 3
	fmt.Println(revstr.Reverse("It's in English, а это на русском"))
	// task 4
	fmt.Println(atoi.Atoi("-13"))
	fmt.Println(atoi.Atoi("25"))
	t4, err := atoi.Atoi("58q")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(t4)
	}
	//task 5
	sort.Imports("tosort.go")
	//task 6
	gen := fib.Fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Print(gen(), " ")
	}
	fmt.Println()
	// task 7
	s := "word слово"
	i := 10
	res, err := runebyindex.RuneByInd(&s, &i)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(string(res))
	}
}
