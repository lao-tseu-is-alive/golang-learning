package benchmark

import (
	"log"
	"testing"
)

/*
Besides the pure test support, the testing package also provides the mechanisms for measuring the code performance.
For this purpose, the B struct pointer as the argument is used, and the benchmarking functions
in the test file are named as BenchmarkXXXX.

The essential part of the benchmark function is the manipulation
with the timer and usage of the loop iteration counter N.

As you can see, the timer is manipulated with the methods,
Reset/Start/StopTimer. By these, the result of the benchmark is influenced.
Note that the timer starts running with the beginning of the benchmark function and
the ResetTimer function just restarts it.

The N field of B is the iteration count within the measurement loop.
The N value is set to a value high enough to reliably measure the result of the benchmark.
The result in the benchmark log then displays the value of iterations and measured time per one iteration.
Execute the benchmark by running : 	go test -bench=.
*/

func BenchmarkSampleOne(b *testing.B) {
	logger := log.New(devNull{}, "test", log.Llongfile)
	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		logger.Println("This si awesome")
	}
	b.StopTimer()
}

type devNull struct{}

func (d devNull) Write(b []byte) (int, error) {
	return 0, nil
}
