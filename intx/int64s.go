package intx

// Int64s array of int64.
type Int64s []int64

// Exists returns true is input value found in array.
func (v Int64s) Exists(in int64) bool {
	return v.Index(in) >= 0
}

// Index returns the position in the array of the first matching value.
func (v Int64s) Index(in int64) int {
	for idx, item := range v {
		if item == in {
			return idx
		}
	}
	return -1
}

// Len returns size of array.
func (v Int64s) Len() int {
	return len(v)
}

// Sort sorts values ​​in ascending order (without allocate).
//
// Insertion sorting method.
func (v Int64s) Sort() {
	for j := 1; j < len(v); j++ {
		// Invariant: v[:j] contains the same elements as
		// the original slice v[:j], but in sorted order.
		key := v[j]
		i := j - 1
		for i >= 0 && v[i] > key {
			v[i+1] = v[i]
			i--
		}
		v[i+1] = key
	}
}

// Copy creates a copy of the array.
func (v Int64s) Copy() Int64s {
	return v.Filter(func(in int64) bool { return true })
}

// Uniq returns a unique set (without allocate).
//
// NOTE: source array becomes sorted. If do not need to change the source array, first create a copy.
func (v Int64s) Uniq() Int64s {
	v.Sort()
	var last int64
	res := v[:0]
	for _, item := range v {
		if item > last {
			res = append(res, item)
			last = item
		}
	}
	return res
}

// Filter returns a filtered array by custom function (without allocate).
func (v Int64s) Filter(f func(in int64) bool) Int64s {
	res := v[:0]
	if f == nil {
		return res
	}
	for _, item := range v {
		if f(item) {
			res = append(res, item)
		}
	}
	return res
}
