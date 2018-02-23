package exercise

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f",
		float64(e))
}

func sqrt(x float64) (float64, error) {

	if x >= 0 {
		z := float64(1.0)
		tmp := float64(1.0)
		for tmp > 0.000001 {
			if z*z-x > 0 {
				tmp = z*z - x
			} else {
				tmp = x - z*z
			}
			z -= (z*z - x) / (2 * z)
		}
		return z, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func MainExerciseError() {
	fmt.Println(sqrt(2))
	fmt.Println(sqrt(-2))
}
