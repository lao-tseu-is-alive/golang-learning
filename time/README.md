# Time

This theme covers instants, durations, formatting, parsing, time zones, timers,
and tickers. Some output depends on the current clock or local time zone. Run
commands from the repository root.

[Back to the learning path](../README.md#learning-path)

## Mental model

A `time.Time` represents an instant and carries location information used to
present its calendar fields. A `time.Duration` represents an elapsed amount of
time as an integer number of nanoseconds. Do not use a duration to mean a
calendar month: months and daylight-saving transitions do not have fixed
lengths. Use `Add` for elapsed time and `AddDate` for calendar arithmetic.

Go formats and parses time using the reference instant
`Mon Jan 2 15:04:05 MST 2006`. The layout describes how that particular instant
would look in the desired format. Prefer standard layouts such as RFC 3339 for
data exchanged between systems.

An instant and its displayed time zone are different concerns. Calling `In`
changes the presentation location, not the instant. Time-zone rules come from a
database and can change; handle `LoadLocation` errors.

Timers deliver one event, while tickers deliver repeated events until stopped.
Long-running code should stop timers or tickers it no longer needs and should
have an explicit cancellation path. For elapsed-time measurements, operations
on values returned by `time.Now` can use Go's monotonic clock reading.

## Core path

| Example | What it teaches | Behavior |
| --- | --- | --- |
| [`now`](now/) | Reading and displaying the current time | Output changes on every run |
| [`format`](format/) | Layouts, predefined formats, padding, and fractions | Deterministic fixed instant |
| [`parse`](parse/) | `Parse`, `ParseInLocation`, and layouts | Includes UTC and local-time parsing |
| [`units`](units/) | Extracting calendar fields | Deterministic fixed instant |
| [`arithmetics`](arithmetics/) | `Add`, negative durations, and `AddDate` | Requires the `Europe/Vienna` zone |
| [`timezones`](timezones/) | Presenting one instant in another location | Requires system or bundled zone data |
| [`epoch`](epoch/) | Converting to and from Unix timestamps | Includes the current clock |
| [`serialize`](serialize/) | JSON/RFC 3339 and Unix timestamp representations | Requires the `Europe/Vienna` zone |

```sh
go run ./time/format
go run ./time/parse
go run ./time/units
go run ./time/arithmetics
go run ./time/timezones
go run ./time/serialize
```

## Clock-driven examples

| Example | What it teaches | Behavior or caution |
| --- | --- | --- |
| [`datediff`](datediff/) | `Sub`, `Since`, and `Until` | Some results depend on the current date |
| [`greetings`](greetings/) | Branching on the local hour | Output depends on local time |
| [`how-long-until-saturday`](how-long-until-saturday/) | Working with `Weekday` | Simplified exercise, not full calendar arithmetic |
| [`wait4delay`](wait4delay/) | `Timer`, `AfterFunc`, and `After` | Takes about 14 seconds |
| [`run-every-5-seconds`](run-every-5-seconds/) | A ticker stopped by an OS signal | Runs until interrupted with Ctrl-C |
| [`timeout`](timeout/) | Selecting on a timeout channel | Busy-loops and allocates heavily for five seconds; demonstration only |

## Check your understanding

- Explain the difference between adding 24 hours and adding one calendar day.
- Parse a timestamp with an explicit zone and serialize it as RFC 3339.
- Inject a clock or fixed `time.Time` into code so its tests are deterministic.
- Replace a polling loop with a timer, ticker, or cancellation signal.

Continue with [`filesystem`](../filesystem/README.md).

## Further reading

- [`time` package](https://pkg.go.dev/time)
- [`time.Time.Format` examples](https://pkg.go.dev/time#example-Time.Format)
- [Monotonic clocks in `time.Time`](https://pkg.go.dev/time#hdr-Monotonic_Clocks)
