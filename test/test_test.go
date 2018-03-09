package test

import (
	"testing"
	"fmt"
)

func BuildPath(path string) string {
	rs := []rune(path)
	length := len(rs)
	fmt.Println(string(rs[length-1]) != string('/'))
	if string(rs[length-1]) != string('/') {
		rs = append(rs, '/')
	}
	return string(rs);
}

func DiffIntSlice(a []int, b []int) ([]int) {
	var diff []int
	var flag = 0
	for _, a_item := range a {
		for _, b_item := range b {
			if (a_item == b_item) {
				flag = 1;
				break
			} else {
				flag = 0;
			}
		}
		if flag == 0 {
			diff = append(diff,a_item)
		}
	}
	return diff
}

func Test_DiffArray(t *testing.T) {
	diff := DiffIntSlice([]int{3, 4},[]int{1, 2, 3})
	fmt.Println(diff)
	t.Log(diff)
}
