# Benchmarking logging libraries for Go

I compared a varity of logging libraries for Go, and performed 2 types of tests
for 2 logging formats.

Types of test:

- Positive tests: everything will be logged because the baseline log level is
  set to lower or equal the level I am logging with.
- Negative tests: nothing will be logged because I set the baseline log level
  to ```ERROR``` but log wih ```INFO```. This is important, because the faster
  it bails-out, and the less pressure the library puts on the GC, the better!

Formats:

- Text: there are a few minor differences among loggers, but I tried to follow
  a ```Time Level Message``` format).
- JSON: few loggers support this but it is still interesting.

# Running the tests

On a terminal, just execute:

```shell
make bench

```

# Benchmark Results (2023-03-27)
## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.15 | 0 | 0 |
| Logrus | 0.43 | 0 | 0 |
| Zerolog | 0.49 | 0 | 0 |
| Gokit | 10.68 | 32 | 1 |
| Seelog | 18.03 | 40 | 2 |
| Zap | 48.92 | 192 | 1 |
| Gologging | 72.65 | 144 | 2 |
| Log15 | 239.30 | 456 | 3 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.79 | 0 | 0 |
| Zerolog | 1.18 | 0 | 0 |
| Gokit | 29.29 | 128 | 1 |
| Zap | 47.61 | 192 | 1 |
| Logrus | 157.00 | 496 | 4 |
| Log15 | 292.60 | 648 | 5 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 57.51 | 0 | 0 |
| Zerolog | 69.40 | 0 | 0 |
| Zap | 193.80 | 192 | 1 |
| Buildin | 200.90 | 48 | 1 |
| Gokit | 207.80 | 256 | 4 |
| Gologging | 483.80 | 864 | 16 |
| Seelog | 1013.00 | 432 | 11 |
| Logrus | 1391.00 | 520 | 15 |
| Log15 | 1520.00 | 872 | 14 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 66.60 | 0 | 0 |
| Zerolog | 68.81 | 0 | 0 |
| Zap | 196.80 | 192 | 1 |
| Gokit | 783.90 | 1465 | 22 |
| Logrus | 2630.00 | 2133 | 32 |
| Log15 | 2774.00 | 1946 | 28 |

