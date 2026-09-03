# Numbers

This theme covers Go's numeric types, conversions, formatting, floating-point
behavior, arbitrary precision, and random numbers. Run commands from the
repository root.

[Back to the learning path](../README.md#learning-path)

## Mental model

Go has distinct integer, floating-point, and complex types. Conversions between
them are explicit because they can lose range or precision. Integer arithmetic
is exact within the type's range; overflow does not produce an error, so choose
a type that fits the domain and validate untrusted values before arithmetic.

Floating-point values approximate real numbers in a finite binary format. Many
decimal fractions cannot be represented exactly. Equality may be appropriate
for values with an exact representation, but measurements and calculated
results often need a tolerance chosen for the problem—not one universal
epsilon.

Parsing turns text into a number and can fail; formatting turns a number into a
representation for people or protocols. Keep the numeric value separate from
its displayed form.

Use `math/big` when ordinary types cannot provide the required range or
precision. Use `math/rand` for simulations and reproducible sequences, and
`crypto/rand` for tokens, keys, and other security-sensitive randomness.

## Core path

| Example | What it teaches | Level |
| --- | --- | --- |
| [`strconv2num`](strconv2num/) | Parsing decimal, binary, hexadecimal, and floating-point text | Beginner |
| [`convert`](convert/) | Converting integer representations between bases | Beginner |
| [`formatnumber`](formatnumber/) | `fmt` verbs, width, precision, and numeric bases | Beginner |
| [`round`](round/) | Why integer conversion is not rounding | Beginner |
| [`tolerance`](tolerance/) | Comparing calculated floating-point values | Intermediate |
| [`angles`](angles/) | Named numeric types, methods, radians, and degrees | Intermediate |
| [`complex`](complex/) | Complex arithmetic and `math/cmplx` | Intermediate |
| [`random`](random/) | Deterministic pseudo-randomness versus cryptographic randomness | Intermediate |

```sh
go run ./numbers/strconv2num
go run ./numbers/convert
go run ./numbers/formatnumber
go run ./numbers/round
go run ./numbers/tolerance
go run ./numbers/angles
go run ./numbers/random
```

## Further examples

| Example | What it teaches | Requirements or caution |
| --- | --- | --- |
| [`bignumber`](bignumber/) | Large constants and configurable `big.Float` precision | Precision is measured in bits |
| [`verybigfloat`](verybigfloat/) | High-precision arithmetic with `math/big` | Uses the external `golog` module |
| [`logarithms`](logarithms/) | Natural, base-2, base-10, and change-of-base logarithms | Includes an invalid negative input that produces `NaN` |
| [`plurals`](plurals/) | Locale-aware plural selection | Uses `golang.org/x/text` |
| [`sqrt-newton-testing`](sqrt-newton-testing/) | Comparing Newton's method with `math.Sqrt` | Uses the local [`mymath`](mymath/) package and `golog` |

## Check your understanding

- Parse user input without discarding the returned error.
- Explain why `0.1 + 0.2` needs care in a comparison.
- Choose whether a seeded pseudo-random generator or `crypto/rand` fits a given
  use case.
- Add table-driven tests for a conversion or rounding function.

Continue with [`time`](../time/README.md).

## Further reading

- [Numeric types in the Go specification](https://go.dev/ref/spec#Numeric_types)
- [`strconv` package](https://pkg.go.dev/strconv)
- [`math` package](https://pkg.go.dev/math)
- [`math/big` package](https://pkg.go.dev/math/big)
- [`crypto/rand` package](https://pkg.go.dev/crypto/rand)
