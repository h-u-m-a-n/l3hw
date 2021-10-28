package main

import (
	"github.com/h-u-m-a-n/l3hw/sort"
	"github.com/h-u-m-a-n/l3hw/atoi"
	"github.com/h-u-m-a-n/l3hw/itoa"


	"fmt"
	"github.com/h-u-m-a-n/l3hw/revstr"

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
		fmt.Println(err)
	} else {
		fmt.Println(t4)
	}
	// task 5
	sort.Imports("main.go")
	// task 6

	// task 7

}
