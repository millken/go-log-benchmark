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

# Benchmark Results (2023-03-26)
## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.14 | 0 | 0 |
| Logrus | 0.42 | 0 | 0 |
| Zerolog | 0.49 | 0 | 0 |
| Gokit | 10.75 | 32 | 1 |
| Seelog | 16.10 | 40 | 2 |
| Zap | 47.18 | 192 | 1 |
| Gologging | 71.63 | 144 | 2 |
| Log15 | 241.70 | 456 | 3 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.79 | 0 | 0 |
| Zerolog | 1.17 | 0 | 0 |
| Gokit | 30.77 | 128 | 1 |
| Zap | 46.93 | 192 | 1 |
| Logrus | 158.80 | 496 | 4 |
| Log15 | 290.00 | 648 | 5 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 55.04 | 0 | 0 |
| Zerolog | 68.52 | 0 | 0 |
| Zap | 195.60 | 192 | 1 |
| Buildin | 200.00 | 48 | 1 |
| Gokit | 223.50 | 256 | 4 |
| Gologging | 470.70 | 864 | 16 |
| Seelog | 1019.00 | 432 | 11 |
| Logrus | 1401.00 | 520 | 15 |
| Log15 | 1522.00 | 872 | 14 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| Zerolog | 69.05 | 0 | 0 |
| GoLog | 70.79 | 16 | 1 |
| Zap | 197.60 | 192 | 1 |
| Gokit | 904.40 | 1465 | 22 |
| Logrus | 2626.00 | 2133 | 32 |
| Log15 | 2768.00 | 1946 | 28 |

