# Benchmark Results (2023-03-26)
## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.14 | 0 | 0 |
| Logrus | 0.45 | 0 | 0 |
| Zerolog | 0.53 | 0 | 0 |
| Gokit | 10.67 | 32 | 1 |
| Seelog | 17.74 | 40 | 2 |
| Zap | 49.29 | 192 | 1 |
| Gologging | 74.47 | 144 | 2 |
| Log15 | 244.90 | 456 | 3 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.80 | 0 | 0 |
| Zerolog | 1.17 | 0 | 0 |
| Gokit | 29.98 | 128 | 1 |
| Zap | 50.64 | 192 | 1 |
| Logrus | 151.60 | 496 | 4 |
| Log15 | 294.40 | 648 | 5 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 53.05 | 0 | 0 |
| Zerolog | 74.44 | 0 | 0 |
| Buildin | 200.40 | 48 | 1 |
| Zap | 207.20 | 192 | 1 |
| Gokit | 212.50 | 256 | 4 |
| Gologging | 495.50 | 864 | 16 |
| Seelog | 1014.00 | 432 | 11 |
| Logrus | 1404.00 | 520 | 15 |
| Log15 | 1525.00 | 872 | 14 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 67.39 | 0 | 0 |
| Zerolog | 73.76 | 0 | 0 |
| Zap | 210.70 | 192 | 1 |
| Gokit | 806.10 | 1465 | 22 |
| Logrus | 2656.00 | 2133 | 32 |
| Log15 | 2771.00 | 1946 | 28 |

