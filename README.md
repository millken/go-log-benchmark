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
make

```

# Results
`
$ go test -v -benchmem -run=^$$ -bench=. ./...
goos: darwin
goarch: arm64
pkg: go-log-benchmark
BenchmarkBuildinTextPositive
BenchmarkBuildinTextPositive-8     	 5539008	       205.9 ns/op	      48 B/op	       1 allocs/op
BenchmarkGokitJSONPositive
BenchmarkGokitJSONPositive-8       	 1627938	       727.5 ns/op	    1592 B/op	      24 allocs/op
BenchmarkGokitJSONNegative
BenchmarkGokitJSONNegative-8       	64103988	        21.22 ns/op	     128 B/op	       1 allocs/op
BenchmarkGokitTextPositive
BenchmarkGokitTextPositive-8       	 5339086	       225.5 ns/op	     256 B/op	       4 allocs/op
BenchmarkGokitTextNegative
BenchmarkGokitTextNegative-8       	135220968	         9.079 ns/op	      32 B/op	       1 allocs/op
BenchmarkGoLogTextPositive
BenchmarkGoLogTextPositive-8       	21435716	        57.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoLogTextNegative
BenchmarkGoLogTextNegative-8       	1000000000	         0.1634 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoLogJSONNegative
BenchmarkGoLogJSONNegative-8       	50011893	        21.27 ns/op	      96 B/op	       1 allocs/op
BenchmarkGoLogJSONPositive
BenchmarkGoLogJSONPositive-8       	 2530756	       401.7 ns/op	     722 B/op	       6 allocs/op
BenchmarkGologgingTextNegative
BenchmarkGologgingTextNegative-8   	17435313	        68.88 ns/op	     144 B/op	       2 allocs/op
BenchmarkGologgingTextPositive
BenchmarkGologgingTextPositive-8   	 2293288	       522.4 ns/op	     912 B/op	      16 allocs/op
BenchmarkLog15TextNegative
BenchmarkLog15TextNegative-8       	 6115641	       217.7 ns/op	     440 B/op	       3 allocs/op
BenchmarkLog15TextPositive
BenchmarkLog15TextPositive-8       	  639520	      1877 ns/op	     904 B/op	      14 allocs/op
BenchmarkLog15JSONNegative
BenchmarkLog15JSONNegative-8       	 5217562	       244.1 ns/op	     632 B/op	       5 allocs/op
BenchmarkLog15JSONPositive
BenchmarkLog15JSONPositive-8       	  311836	      4133 ns/op	    2057 B/op	      30 allocs/op
BenchmarkLogrusTextPositive
BenchmarkLogrusTextPositive-8      	  664591	      1804 ns/op	     520 B/op	      15 allocs/op
BenchmarkLogrusTextNegative
BenchmarkLogrusTextNegative-8      	1000000000	         0.4587 ns/op	       0 B/op	       0 allocs/op
BenchmarkLogrusJSONNegative
BenchmarkLogrusJSONNegative-8      	10244847	       117.2 ns/op	     496 B/op	       4 allocs/op
BenchmarkLogrusJSONPositive
BenchmarkLogrusJSONPositive-8      	  330482	      3652 ns/op	    2260 B/op	      34 allocs/op
BenchmarkSeelogTextNegative
BenchmarkSeelogTextNegative-8      	76014715	        13.38 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeelogTextPositive
BenchmarkSeelogTextPositive-8      	  973884	      1190 ns/op	     432 B/op	      11 allocs/op
BenchmarkZapTextPositive
BenchmarkZapTextPositive-8         	 8952968	       139.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkZapTextNegative
BenchmarkZapTextNegative-8         	953091988	         1.337 ns/op	       0 B/op	       0 allocs/op
BenchmarkZapJSONPositive
BenchmarkZapJSONPositive-8         	 5434003	       229.6 ns/op	     192 B/op	       1 allocs/op
BenchmarkZapJSONNegative
BenchmarkZapJSONNegative-8         	39555006	        31.56 ns/op	     192 B/op	       1 allocs/op
BenchmarkZerologTextPositive
BenchmarkZerologTextPositive-8     	13272378	        90.87 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologTextNegative
BenchmarkZerologTextNegative-8     	1000000000	         0.5163 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologJSONNegative
BenchmarkZerologJSONNegative-8     	920908063	         1.414 ns/op	       0 B/op	       0 allocs/op
BenchmarkZerologJSONPositive
BenchmarkZerologJSONPositive-8     	10910830	       107.1 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	go-log-benchmark	37.940s
$ benchstat -sort -name benchmark.txt 
name                     time/op
ZerologTextPositive-8    90.9ns ± 0%
ZerologTextNegative-8    0.52ns ± 0%
ZerologJSONPositive-8     107ns ± 0%
ZerologJSONNegative-8    1.41ns ± 0%
ZapTextPositive-8         140ns ± 0%
ZapTextNegative-8        1.34ns ± 0%
ZapJSONPositive-8         230ns ± 0%
ZapJSONNegative-8        31.6ns ± 0%
SeelogTextPositive-8     1.19µs ± 0%
SeelogTextNegative-8     13.4ns ± 0%
LogrusTextPositive-8     1.80µs ± 0%
LogrusTextNegative-8     0.46ns ± 0%
LogrusJSONPositive-8     3.65µs ± 0%
LogrusJSONNegative-8      117ns ± 0%
Log15TextPositive-8      1.88µs ± 0%
Log15TextNegative-8       218ns ± 0%
Log15JSONPositive-8      4.13µs ± 0%
Log15JSONNegative-8       244ns ± 0%
GologgingTextPositive-8   522ns ± 0%
GologgingTextNegative-8  68.9ns ± 0%
GokitTextPositive-8       226ns ± 0%
GokitTextNegative-8      9.08ns ± 0%
GokitJSONPositive-8       728ns ± 0%
GokitJSONNegative-8      21.2ns ± 0%
GoLogTextPositive-8      57.3ns ± 0%
GoLogTextNegative-8      0.16ns ± 0%
GoLogJSONPositive-8       402ns ± 0%
GoLogJSONNegative-8      21.3ns ± 0%
BuildinTextPositive-8     206ns ± 0%

name                     alloc/op
ZerologTextPositive-8     0.00B     
ZerologTextNegative-8     0.00B     
ZerologJSONPositive-8     0.00B     
ZerologJSONNegative-8     0.00B     
ZapTextPositive-8         0.00B     
ZapTextNegative-8         0.00B     
ZapJSONPositive-8          192B ± 0%
ZapJSONNegative-8          192B ± 0%
SeelogTextPositive-8       432B ± 0%
SeelogTextNegative-8      40.0B ± 0%
LogrusTextPositive-8       520B ± 0%
LogrusTextNegative-8      0.00B     
LogrusJSONPositive-8     2.26kB ± 0%
LogrusJSONNegative-8       496B ± 0%
Log15TextPositive-8        904B ± 0%
Log15TextNegative-8        440B ± 0%
Log15JSONPositive-8      2.06kB ± 0%
Log15JSONNegative-8        632B ± 0%
GologgingTextPositive-8    912B ± 0%
GologgingTextNegative-8    144B ± 0%
GokitTextPositive-8        256B ± 0%
GokitTextNegative-8       32.0B ± 0%
GokitJSONPositive-8      1.59kB ± 0%
GokitJSONNegative-8        128B ± 0%
GoLogTextPositive-8       0.00B     
GoLogTextNegative-8       0.00B     
GoLogJSONPositive-8        722B ± 0%
GoLogJSONNegative-8       96.0B ± 0%
BuildinTextPositive-8     48.0B ± 0%

name                     allocs/op
ZerologTextPositive-8      0.00     
ZerologTextNegative-8      0.00     
ZerologJSONPositive-8      0.00     
ZerologJSONNegative-8      0.00     
ZapTextPositive-8          0.00     
ZapTextNegative-8          0.00     
ZapJSONPositive-8          1.00 ± 0%
ZapJSONNegative-8          1.00 ± 0%
SeelogTextPositive-8       11.0 ± 0%
SeelogTextNegative-8       2.00 ± 0%
LogrusTextPositive-8       15.0 ± 0%
LogrusTextNegative-8       0.00     
LogrusJSONPositive-8       34.0 ± 0%
LogrusJSONNegative-8       4.00 ± 0%
Log15TextPositive-8        14.0 ± 0%
Log15TextNegative-8        3.00 ± 0%
Log15JSONPositive-8        30.0 ± 0%
Log15JSONNegative-8        5.00 ± 0%
GologgingTextPositive-8    16.0 ± 0%
GologgingTextNegative-8    2.00 ± 0%
GokitTextPositive-8        4.00 ± 0%
GokitTextNegative-8        1.00 ± 0%
GokitJSONPositive-8        24.0 ± 0%
GokitJSONNegative-8        1.00 ± 0%
GoLogTextPositive-8        0.00     
GoLogTextNegative-8        0.00     
GoLogJSONPositive-8        6.00 ± 0%
GoLogJSONNegative-8        1.00 ± 0%
BuildinTextPositive-8      1.00 ± 0%
`
