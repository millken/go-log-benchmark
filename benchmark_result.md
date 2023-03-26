## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.14 | 0 | 0 |
| Zerolog | 0.48 | 0 | 0 |
| Logrus | 0.54 | 0 | 0 |
| Gokit | 10.88 | 32 | 1 |
| Seelog | 20.37 | 40 | 2 |
| Zap | 52.90 | 192 | 1 |
| Gologging | 84.54 | 144 | 2 |
| Log15 | 244.70 | 456 | 3 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.81 | 0 | 0 |
| Zerolog | 1.26 | 0 | 0 |
| Gokit | 32.02 | 128 | 1 |
| Zap | 49.19 | 192 | 1 |
| Logrus | 161.60 | 496 | 4 |
| Log15 | 300.20 | 648 | 5 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 55.41 | 0 | 0 |
| Zerolog | 68.64 | 0 | 0 |
| Buildin | 192.20 | 48 | 1 |
| Zap | 208.30 | 192 | 1 |
| Gokit | 244.70 | 256 | 4 |
| Gologging | 567.00 | 864 | 16 |
| Seelog | 1032.00 | 432 | 11 |
| Logrus | 1399.00 | 520 | 15 |
| Log15 | 1522.00 | 872 | 14 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| Zerolog | 71.18 | 0 | 0 |
| GoLog | 87.54 | 16 | 1 |
| Zap | 201.40 | 192 | 1 |
| Gokit | 795.50 | 1465 | 22 |
| Logrus | 2662.00 | 2133 | 32 |
| Log15 | 2812.00 | 1946 | 28 |

