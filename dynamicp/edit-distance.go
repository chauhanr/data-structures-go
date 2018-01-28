package dynamicp

import "math"

func EditDistance(x string, y string) int {
	a := []string{}
	b := []string{}
	for _, xd := range x {
		a = append(a, string(xd))
	}
	for _, yd := range y {
		b = append(b, string(yd))
	}
	return Distance(a, b)
}

func Distance(a []string, b []string) int {
	distance := math.MaxInt16
	if len(a) == 0 || len(b) == 0 {
		if len(a) > len(b) {
			return len(a) - len(b)
		} else {
			return len(b) - len(a)
		}
	}
	constant := 0
	if a[len(a)-1] == b[len(b)-1] {
		constant = 0
	} else {
		constant = 1
	}
	distance = multiMin(Distance(a[0:len(a)-1], b)+1, Distance(a, b[0:len(b)-1])+1, Distance(a[0:len(a)-1], b[0:len(b)-1])+constant)
	return distance
}

func multiMin(values ...int) int {
	min := math.MaxInt16
	for _, val := range values {
		if min > val {
			min = val
		}
	}
	return min
}
