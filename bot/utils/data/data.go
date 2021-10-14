package data

func Int64InSlice(e int64, s []int64) bool {
	for _, x := range s {
		if x == e {
			return true
		}
	}
	return false
}
