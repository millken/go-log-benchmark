## TextNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.14 | 0 | 0 |
| Logrus | 0.42 | 0 | 0 |
| Zerolog | 0.50 | 0 | 0 |
| Gokit | 11.85 | 32 | 1 |
| Seelog | 17.59 | 40 | 2 |
| Zap | 49.72 | 192 | 1 |
| Gologging | 80.96 | 144 | 2 |
| Log15 | 242.40 | 456 | 3 |

## JSONNegative
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 0.78 | 0 | 0 |
| Zerolog | 1.17 | 0 | 0 |
| Gokit | 30.91 | 128 | 1 |
| Zap | 50.45 | 192 | 1 |
| Logrus | 160.40 | 496 | 4 |
| Log15 | 307.30 | 648 | 5 |

## TextPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| GoLog | 54.20 | 0 | 0 |
| Zerolog | 68.35 | 0 | 0 |
| Zap | 196.50 | 192 | 1 |
| Buildin | 198.70 | 48 | 1 |
| Gokit | 214.90 | 256 | 4 |
| Gologging | 526.60 | 864 | 16 |
| Seelog | 1014.00 | 432 | 11 |
| Logrus | 1394.00 | 520 | 15 |
| Log15 | 1522.00 | 872 | 14 |

## JSONPositive
| Name | ns/op | B/op | allocs/op |
| --------- | --------- | --------- | --------- |
| Zerolog | 68.08 | 0 | 0 |
| GoLog | 71.51 | 16 | 1 |
| Zap | 200.20 | 192 | 1 |
| Gokit | 798.60 | 1465 | 22 |
| Logrus | 2629.00 | 2133 | 32 |
| Log15 | 2823.00 | 1946 | 28 |

