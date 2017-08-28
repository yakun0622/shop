package tools

//RemoveDuplicates 去除slice中重复的数据
func RemoveDuplicates(s []int) (d []int) {
	m := map[int]bool{}

	// walk the slice and for each value we've not seen so far
	// move it to slot K here K i the number of unique values
	// we've seen so far. below, K is represented by `len(m)`
	// after the loop you're left with a slice all unique values at the front
	// in their original order so you simply to re-slice to K to get only the unique values
	for _, v := range s {
		if _, seen := m[v]; !seen {
			s[len(m)] = v
			m[v] = true
		}
	}
	// re-slice s to the number of unique values
	d = s[:len(m)]
	return
}

//多个slice合并
func Slice_merge(slice1, slice2 []interface{}) (c []interface{}) {
	c = append(slice1, slice2...)
	return
}

//判断是否存在于slice中
func In_slice(val interface{}, slice []interface{}) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

//判断是否存在于slice中
func InIntSlice(val int, slice []int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func Slice_unique(slice []interface{}) (uniqueslice []interface{}) {
	for _, v := range slice {
		if !In_slice(v, uniqueslice) {
			uniqueslice = append(uniqueslice, v)
		}
	}
	return
}