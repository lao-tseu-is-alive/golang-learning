# Structs, methods, interfaces, and errors

This theme moves from concrete data types to behavior expressed through
methods and interfaces. Follow the first group in order; use the remaining
examples to deepen individual concepts. Run commands from the repository root.

[Back to the learning path](../README.md#learning-path)

## Mental model

A struct groups related fields into one concrete value. It describes what data
the value contains. A method associates behavior with a named type through a
receiver; the receiver is an explicit participant in the method, not a hidden
class object.

Choose a value receiver when a method can work on a copy and the type is small.
Choose a pointer receiver when the method must mutate the value, when copying is
undesirable, or when the type's method set requires it. Be consistent for a
given type unless there is a clear reason not to be.

An interface describes required behavior as a method set. Types satisfy it
implicitly: there is no `implements` declaration. This allows the consumer of a
behavior to define the smallest interface it needs. Prefer accepting small
interfaces and returning useful concrete types.

An interface value carries both a dynamic type and a dynamic value. A nil
pointer stored inside an interface therefore does not make the interface itself
nil—an important distinction explored in `interface-values`.

Finally, `error` is just a standard interface. Returning an error lets the
caller decide whether to handle, wrap, log, or propagate a failure.

## Core path

| Example | What it teaches | Level |
| --- | --- | --- |
| [`struct-literals`](struct-literals/) | Struct values, keyed fields, zero values, and pointers | Beginner |
| [`methods`](methods/) | Value receivers and pointer receivers | Beginner |
| [`stringer-interface`](stringer-interface/) | Implementing `fmt.Stringer` implicitly | Beginner |
| [`my-errors`](my-errors/) | Defining a type that implements `error` | Beginner |
| [`interface-values`](interface-values/) | Dynamic interface values and nil concrete pointers | Intermediate |
| [`type-assertions`](type-assertions/) | Inspecting an interface value with the comma-ok form | Intermediate |
| [`type-switch`](type-switch/) | Branching on an interface value's dynamic type | Intermediate |

```sh
go run ./struct/struct-literals
go run ./struct/methods
go run ./struct/stringer-interface
go run ./struct/my-errors
go run ./struct/interface-values
go run ./struct/type-assertions
go run ./struct/type-switch
```

## Further examples

| Example | What it teaches | Notes |
| --- | --- | --- |
| [`ipaddr-stringer`](ipaddr-stringer/) | Formatting a custom value with `fmt.Stringer` | Go Tour exercise |
| [`empty-interface`](empty-interface/) | Storing values of different types in an interface | Uses the older `interface{}` spelling; modern code normally writes `any` |
| [`my-reader-type`](my-reader-type/) | Implementing the `io.Reader` contract | Uses the external `golog` module |
| [`rot13-reader`](rot13-reader/) | Wrapping one `io.Reader` with another | Intermediate composition exercise |
| [`testing-myerrors`](testing-myerrors/) | Consuming errors from the local `numbers/mymath` package | Uses the external `golog` module |

## Check your understanding

- Explain when changing a value receiver to a pointer receiver changes program
  behavior.
- Add another type that satisfies an existing interface without editing the
  interface.
- Test an error with `errors.Is` or `errors.As` instead of comparing its text.

Continue with [`strings`](../strings/README.md).

## Further reading

- [Methods and interfaces](https://go.dev/tour/methods/1) in A Tour of Go
- [Implicit interface implementation](https://go.dev/tour/methods/10)
- [Interface values](https://go.dev/tour/methods/11)
- [Effective Go: interfaces and types](https://go.dev/doc/effective_go#interfaces_and_types)
