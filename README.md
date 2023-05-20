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

# Benchmark Results (2023-05-19)
## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.14 | 0 | 0 |
| Logrus | 0.42 | 0 | 0 |
| Zerolog | 0.51 | 0 | 0 |
| Gokit | 11.42 | 32 | 1 |
| Zap | 49.89 | 192 | 1 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.79 | 0 | 0 |
| Zerolog | 1.24 | 0 | 0 |
| Gokit | 31.00 | 128 | 1 |
| Zap | 51.31 | 192 | 1 |
| Logrus | 161.00 | 496 | 4 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| Buildin | 203.30 | 48 | 1 |
| Gokit | 209.20 | 256 | 4 |
| GoLog | 288.80 | 408 | 4 |
| Zap | 532.70 | 530 | 7 |
| Logrus | 1421.00 | 520 | 15 |
| Zerolog | 1597.00 | 2340 | 71 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 70.05 | 0 | 0 |
| Zerolog | 71.42 | 0 | 0 |
| Zap | 222.00 | 192 | 1 |
| Gokit | 745.30 | 1465 | 22 |
| Logrus | 2668.00 | 2132 | 32 |

