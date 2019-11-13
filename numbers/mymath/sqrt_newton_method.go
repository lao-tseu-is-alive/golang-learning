package mymath

import (
	"fmt"
	"github.com/lao-tseu-is-alive/golog"
	"math"
	"time"
)

type ErrNegativeSqrt struct {
	When time.Time
	Num  float64
}

func (e ErrNegativeSqrt) Error() string {
	val := fmt.Sprint(float64(e.Num))
	err := fmt.Sprintf("ERROR at %v :cannot Sqrt negative number: %v\n", e.When, val)
	return err
}

/**

Computers typically compute the square root of x using a loop.
Starting with some guess z, we can adjust z based on how close z² is to x,
producing a better guess:
	z -= (z*z - x) / (2*z)
Repeating this adjustment makes the guess better and better until we reach
an answer that is as close to the actual square root as can be.
Experiment inspired by lesson 8 of flowcontrol in "A Tour of Go"

(Note: If you are interested in the details of the algorithm,
the z² − x above is how far away z² is from where it needs to be (x),
and the division by 2z is the derivative of z², to scale how much we adjust z
by how quickly z² is changing.
This general approach is called Newton's method.
https://en.wikipedia.org/wiki/Newton%27s_method
It works well for many functions but especially well for square root.)

*/
func Sqrt(x float64) (float64, error) {
	golog.Un(golog.Trace("In Sqrt(%v)", x))
	if x < 0 {
		return 0, &ErrNegativeSqrt{time.Now(), x}
	}
	const EPSILON float64 = 1e-12
	z := 1.0
	// let's say we allow a max of 50 steps
	for i := 1; i < 51; i++ {
		z = z - ((z*z - x) / (2 * z))
		delta := x - (z * z)
		// golog.Info("Iteration %d \t z=%v\t delta:%v", i, z, delta)
		if math.Abs(delta-1e-13) <= EPSILON {
			golog.Info("OK that's enough precision ! Let's get out after %d loops at delta:%v", i, delta)
			break
		}
	}
	return z, nil
}
