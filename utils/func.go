package utils


func BuildPath(path string) string {
	rs := []rune(path)
	length := len(rs)
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
			diff = append(diff, a_item)
		}
	}
	return diff
}
