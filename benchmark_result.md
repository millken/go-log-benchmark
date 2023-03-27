# Benchmark Results (2023-03-27)
## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.14 | 0 | 0 |
| Logrus | 0.46 | 0 | 0 |
| Zerolog | 0.53 | 0 | 0 |
| Gokit | 10.85 | 32 | 1 |
| Seelog | 18.96 | 40 | 2 |
| Zap | 47.94 | 192 | 1 |
| Gologging | 74.14 | 144 | 2 |
| Log15 | 237.00 | 456 | 3 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.78 | 0 | 0 |
| Zerolog | 1.34 | 0 | 0 |
| Gokit | 29.25 | 128 | 1 |
| Zap | 48.58 | 192 | 1 |
| Logrus | 168.10 | 496 | 4 |
| Log15 | 297.20 | 648 | 5 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 57.75 | 0 | 0 |
| Zerolog | 68.60 | 0 | 0 |
| Buildin | 203.30 | 48 | 1 |
| Zap | 204.50 | 192 | 1 |
| Gokit | 209.30 | 256 | 4 |
| Gologging | 475.00 | 864 | 16 |
| Seelog | 1021.00 | 432 | 11 |
| Logrus | 1399.00 | 520 | 15 |
| Log15 | 1544.00 | 872 | 14 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| Zerolog | 69.57 | 0 | 0 |
| GoLog | 76.09 | 0 | 0 |
| Zap | 199.80 | 192 | 1 |
| Gokit | 797.80 | 1465 | 22 |
| Logrus | 2624.00 | 2133 | 32 |
| Log15 | 2785.00 | 1946 | 28 |

