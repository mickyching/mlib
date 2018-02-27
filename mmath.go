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
