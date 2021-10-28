package sort

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func Imports(path string) error {
	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	fmt.Println(string(f))
	reg := regexp.MustCompile(`import\s*\(([\s\S]*?)\)`)
	imports := string(reg.Find(f))
	for _, v := range imports {
		fmt.Print(string(v))
	}
	imports = imports[9 : len(imports)-1]
	fmt.Println()
	imps := strings.Split(imports, "\n")
	sort.Strings(imps)
	trim(&imps)
	imps = append(imps, ")")
	imps = append([]string{"import ("}, imps...)

	changed := make([]byte, 0)
	for i, v := range imps {
		temp := []byte(v)
		if i != len(imps)-1 {
			temp = append(temp, '\n')
		}
		changed = append(changed, temp...)
	}
	fmt.Println(string(changed))

	overwrite := reg.ReplaceAll(f, changed)
	fmt.Println(string(overwrite))
	err = os.WriteFile(path, overwrite, 0)
	if err != nil {
		return err
	}
	return nil

}

func trim(arr *[]string) {
	for len((*arr)[0]) == 0 {
		*arr = (*arr)[1:]
	}
	for len((*arr)[len(*arr)-1])==0 {
		*arr = (*arr)[:len(*arr)-1]
	}
}