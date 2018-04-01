package mlib

import "math"

// MinMaxInt returns min/max from int array
func MinMaxInt(aa ...interface{}) (int64, int64) {
	min := int64(math.MaxInt64)
	max := -min
	for _, a := range aa {
		v := Int(a)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// MinInt returns min from int array
func MinInt(aa ...interface{}) int64 {
	min, _ := MinMaxInt(aa...)
	return min
}

// MaxInt returns max from int array
func MaxInt(aa ...interface{}) int64 {
	_, max := MinMaxInt(aa...)
	return max
}

// MinMaxFloat returns min/max from float array
func MinMaxFloat(aa ...interface{}) (float64, float64) {
	min := math.MaxFloat64
	max := -min
	for _, a := range aa {
		v := Float(a)
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

// MinFloat returns min from float array
func MinFloat(aa ...interface{}) float64 {
	min, _ := MinMaxFloat(aa...)
	return min
}

// MaxFloat returns max from float array
func MaxFloat(aa ...interface{}) float64 {
	_, max := MinMaxFloat(aa...)
	return max
}

// Sum return sum
func Sum(aa ...interface{}) float64 {
	sum := 0.0
	for _, a := range aa {
		v := Float(a)
		sum += v
	}
	return sum
}

// Mean return mean value
func Mean(aa ...interface{}) float64 {
	if len(aa) == 0 {
		return 0
	}
	return Sum(aa...) / float64(len(aa))
}

// MSE return mean square error
// 方差：d2 = 1/n sum(xi-x)2
// 均方差=标准差：d = sqrt(d2)
func MSE(aa ...interface{}) float64 {
	if len(aa) == 0 {
		return 0
	}

	dd := 0.0
	m := Mean(aa...)
	for _, a := range aa {
		x := Float(a)
		dd += (x - m) * (x - m)
	}

	return math.Sqrt(dd / float64(len(aa)))
}

// LinearFit return (k, b) fit line y = kx + b
func LinearFit(sx []float64, sy []float64) (k float64, b float64) {
	if len(sx) != len(sy) {
		Fatalf("slice length not match x(%d) != y(%d)", len(sx), len(sy))
	}

	num := float64(len(sx))
	xy := 0.0
	xx := 0.0
	xs := 0.0
	ys := 0.0
	for i := 0; i < len(sx); i++ {
		x := sx[i]
		y := sy[i]
		xx += x * x
		xy += x * y
		xs += x
		ys += y
	}

	k = (num*xy - xs*ys) / (num*xx - xs*xs)
	b = (xx*ys - xs*xy) / (num*xx - xs*xs)

	return k, b
}

// CmpFloats compare slice
// a < b return -1
// a > b return 1
// else return 0
func CmpFloats(a []float64, b []float64) int {
	if len(a) != len(b) {
		Fatalf("length not same %d != %d", len(a), len(b))
	}

	aa := 0
	bb := 0
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			aa++
		} else if a[i] < b[i] {
			bb++
		} else {
			return 0
		}
		if aa != 0 && bb != 0 {
			return 0
		}
	}

	if aa == len(a) {
		return 1
	} else if bb == len(a) {
		return -1
	}
	Fatalf("unreachable code")
	return 0
}
