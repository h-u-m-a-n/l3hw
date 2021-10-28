package runebyindex

import (
	"fmt"
)

func RuneByInd(s *string, i *int) (r rune, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic recovered: %v", r)
		}
	}()

	r = []rune(*s)[*i]
	return
}
