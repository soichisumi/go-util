package slice

import "reflect"

// slice and element must be typed to []V and V respectively
func Contains(slice, element interface{}) bool {
	sValue := reflect.ValueOf(slice)

	if sValue.Kind() == reflect.Slice {
		for i := 0; i < sValue.Len(); i++ {
			if sValue.Index(i).Interface() == element {
				return true
			}
		}
	}
	return false
}

func Sequence(start, times, step uint64) []uint64 {
	if times == 0 {
		return []uint64{}
	}
	res := make([]uint64, times)
	for i, _ := range res {
		res[i] = start
		start += step
	}
	return res
}