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

