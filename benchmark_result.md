# Benchmark Results (2023-03-27)
## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.15 | 0 | 0 |
| Logrus | 0.43 | 0 | 0 |
| Zerolog | 0.53 | 0 | 0 |
| Gokit | 11.06 | 32 | 1 |
| Zap | 48.98 | 192 | 1 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.82 | 0 | 0 |
| Zerolog | 1.23 | 0 | 0 |
| Gokit | 34.01 | 128 | 1 |
| Zap | 48.86 | 192 | 1 |
| Logrus | 155.40 | 496 | 4 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 57.60 | 0 | 0 |
| Zerolog | 73.78 | 0 | 0 |
| Buildin | 199.30 | 48 | 1 |
| Gokit | 216.50 | 256 | 4 |
| Zap | 220.90 | 192 | 1 |
| Logrus | 1415.00 | 520 | 15 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| Zerolog | 74.32 | 0 | 0 |
| GoLog | 80.66 | 0 | 0 |
| Zap | 197.30 | 192 | 1 |
| Gokit | 783.20 | 1465 | 22 |
| Logrus | 2605.00 | 2132 | 32 |

