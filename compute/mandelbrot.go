package compute

import (
	"math/cmplx"
)

const (
	Threshold  uint8 = 255
	Upperbound uint8 = 4
)

// Computes the number of iteration before the complex number exceeds its Upperbound.
func Compute(c complex128) uint8 {
	// helper function for the recursive compute call
	var computeIter func(complex128, uint8) uint8 // seperate the decleration from implementation to be able to use it within current scope

	computeIter = func(cc complex128, iter uint8) uint8 {
		if iter >= Threshold {
			return Threshold
		}
		n := cc*cc + c
		absolute := cmplx.Abs(n * n)
		isAboveUpperBound := absolute >= float64(Upperbound)
		if isAboveUpperBound {
			return iter
		}

		return computeIter(n, iter+1)
	}

	return computeIter(c, 1)
}
