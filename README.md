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

	go test -cpu=1,2,4 -benchmem -benchtime=5s -bench "Benchmark.*TextNegative" |tee TextNegative.txt
	benchstat -csv -sort -name TextNegative.txt > text.csv
  https://www.convertcsv.com/csv-to-markdown.htm
```

# Results

I ran these tests on Mac OSX 10.11.4 using 1.6 GHz Intel Core i5 Macbook Air
with 8 GB 1600 MHz DDR3 memory.

Overall, Go kit performed the best, though it has a quirky way of specifying
levels.

I was surprised with log15 because it seemed not to take advantage of the 2 or 4
goroutines. I believe this is because it probably performs most of its
operations inside some mutex-protected execution path right before flushing the
log to the output stream.

When it comes to negative tests for JSON, all loggers made too many allocations
for my taste. This is especially bad if you carpet bomb your program with
debug-level log lines but want to disable it in production: it puts too much
unecessary pressure on GC, leads to all sorts of problems including memory
fragmentation. This is something I would like to see improved in future
versions!

## Take 2 (September 2017)

Two years later, same test ran with Go 1.9 on Mac OSX 10.12.6 using 2.7GHz
Intel Core i7 Macbook Pro with 16 GB 2133 MHz LPDDR3 memory.

Overall, zerolog performed the best by a substantial margin with a constant
0 allocations for both text and JSON (output is always JSON) with positive and
negative tests.

## benchstat

### TextPositive

|name                   |time/op (ns/op)      |alloc/op (B/op)|allocs/op (allocs/op)|
|-----------------------|---------------------|---------------|---------------------|
|ZerologTextPositive-4  |102.8                |0              |0                    |
|ZerologTextPositive-2  |189.2                |0              |0                    |
|ZerologTextPositive    |346.6                |0              |0                    |
|SeelogTextPositive-4   |1199                 |432            |11                   |
|SeelogTextPositive-2   |1204                 |432            |11                   |
|SeelogTextPositive     |1009                 |432            |11                   |
|LogrusTextPositive-4   |1783                 |520            |15                   |
|LogrusTextPositive-2   |1776                 |520            |15                   |
|LogrusTextPositive     |1477                 |520            |15                   |
|Log15TextPositive-4    |1803                 |904            |14                   |
|Log15TextPositive-2    |1833                 |904            |14                   |
|Log15TextPositive      |1639                 |904            |14                   |
|GologgingTextPositive-4|353.4                |912            |16                   |
|GologgingTextPositive-2|584.6                |912            |16                   |
|GologgingTextPositive  |1035                 |912            |16                   |
|GokitTextPositive-4    |225.8                |256            |4                    |
|GokitTextPositive-2    |394.2                |256            |4                    |
|GokitTextPositive      |723.1                |256            |4                    |
|GoLogTextPositive-4    |128.8                |0              |0                    |
|GoLogTextPositive-2    |102.4                |0              |0                    |
|GoLogTextPositive      |124.5                |0              |0                    |

### TextNegative

|name                   |time/op (ns/op)      |alloc/op (B/op)|allocs/op (allocs/op)|
|-----------------------|---------------------|---------------|---------------------|
|ZerologTextNegative-4  |0.6935               |0              |0                    |
|ZerologTextNegative-2  |1.327                |0              |0                    |
|ZerologTextNegative    |2.575                |0              |0                    |
|ZapTextNegative-4      |1.678                |0              |0                    |
|ZapTextNegative-2      |3.199                |0              |0                    |
|ZapTextNegative        |6.274                |0              |0                    |
|SeelogTextNegative-4   |11.85                |40             |2                    |
|SeelogTextNegative-2   |20.88                |40             |2                    |
|SeelogTextNegative     |38.02                |40             |2                    |
|LogrusTextNegative-4   |0.5767               |0              |0                    |
|LogrusTextNegative-2   |1.102                |0              |0                    |
|LogrusTextNegative     |2.181                |0              |0                    |
|Log15TextNegative-4    |213.9                |440            |3                    |
|Log15TextNegative-2    |387.7                |440            |3                    |
|Log15TextNegative      |728.8                |440            |3                    |
|GologgingTextNegative-4|46                   |144            |2                    |
|GologgingTextNegative-2|70.91                |144            |2                    |
|GologgingTextNegative  |105.5                |144            |2                    |
|GokitTextNegative-4    |7.097                |32             |1                    |
|GokitTextNegative-2    |12.4                 |32             |1                    |
|GokitTextNegative      |21.83                |32             |1                    |
|GoLogTextNegative-4    |2.2                  |0              |0                    |
|GoLogTextNegative-2    |4.272                |0              |0                    |
|GoLogTextNegative      |8.202                |0              |0                    |


### JSONPositive

| test                    | op time        | op alloc sz | op alloc count |
|-------------------------|----------------|-------------|----------------|
| GokitJSONPositive-4     |  1.42µs ± 4%   | 1.55kB ± 0% |      24.0 ± 0% |
| Log15JSONPositive-4     |  6.56µs ± 1%   | 2.01kB ± 0% |      30.0 ± 0% |
| LogrusJSONPositive-4    |  1.81µs ± 3%   | 2.45kB ± 0% |      33.0 ± 0% |
| ZerologJSONPositive-4   | **195ns ± 3%** |**0.00B**    |    **0.00**    |

### JSONNegative

| test                    | op time         | op alloc sz | op alloc count |
|-------------------------|-----------------|-------------|----------------|
| GokitJSONNegative-4     |   27.3ns ± 2%   |   128B ± 0% |      1.00 ± 0% |
| Log15JSONNegative-4     |    189ns ± 2%   |   320B ± 0% |      3.00 ± 0% |
| LogrusJSONNegative-4    |    257ns ± 2%   |   752B ± 0% |      5.00 ± 0% |
| ZerologJSONNegative-4   | **6.39ns ± 2%** |**0.00B**    |    **0.00**    |

## Raw data

### TextPositive

| test                             | ops      | ns/op         | bytes/op    | allocs/op       |
|----------------------------------|----------|---------------|-------------|-----------------|
| BenchmarkGokitTextPositive-4     | 20000000 |   428 ns/op   |  256 B/op   |   4 allocs/op   |
| BenchmarkGologgingTextPositive-4 | 10000000 |   621 ns/op   |  920 B/op   |  15 allocs/op   |
| BenchmarkLog15TextPositive-4     |  2000000 |  3612 ns/op   | 1120 B/op   |  24 allocs/op   |
| BenchmarkLogrusTextPositive-4    | 10000000 |   657 ns/op   |  320 B/op   |  10 allocs/op   |
| BenchmarkSeelogTextPositive-4    |  3000000 |  2197 ns/op   |  440 B/op   |  11 allocs/op   |
| BenchmarkZerologTextPositive-4   | 50000000 | **125 ns/op** |  **0 B/op** | **0 allocs/op** |

### TextNegative

| test                             | ops         | ns/op          | bytes/op    | allocs/op       |
|----------------------------------|-------------|----------------|-------------|-----------------|
| BenchmarkGokitTextNegative-4     |   500000000 |   16.7 ns/op   |   32 B/op   |   1 allocs/op   |
| BenchmarkGologgingTextNegative-4 |   100000000 |   60.8 ns/op   |  144 B/op   |   2 allocs/op   |
| BenchmarkLog15TextNegative-4     |    50000000 |    146 ns/op   |  128 B/op   |   1 allocs/op   |
| BenchmarkLogrusTextNegative-4    | 10000000000 | **1.02 ns/op** |    0 B/op   |   0 allocs/op   |
| BenchmarkSeelogTextNegative-4    |   300000000 |   22.1 ns/op   |   48 B/op   |   2 allocs/op   |
| BenchmarkZerologTextNegative-4   |  2000000000 |   4.34 ns/op   |  **0 B/op** | **0 allocs/op** |

### JSONPositive

|name                   |time/op (ns/op)      |alloc/op (B/op)|allocs/op (allocs/op)|
|-----------------------|---------------------|---------------|---------------------|
|ZerologJSONPositive-4  |142.5                |0              |0                    |
|ZerologJSONPositive-2  |256.9                |0              |0                    |
|ZerologJSONPositive    |475.4                |0              |0                    |
|ZapJSONPositive-4      |251.7                |192            |1                    |
|ZapJSONPositive-2      |395.9                |192            |1                    |
|ZapJSONPositive        |721.9                |192            |1                    |
|LogrusJSONPositive-4   |3744                 |2258           |34                   |
|LogrusJSONPositive-2   |3514                 |2256           |34                   |
|LogrusJSONPositive     |3193                 |2256           |34                   |
|Log15JSONPositive-4    |3977                 |2056           |30                   |
|Log15JSONPositive-2    |3758                 |2056           |30                   |
|Log15JSONPositive      |3487                 |2056           |30                   |
|GokitJSONPositive-4    |746.8                |1592           |24                   |
|GokitJSONPositive-2    |1327                 |1592           |24                   |
|GokitJSONPositive      |2407                 |1592           |24                   |
|GoLogJSONPositive-4    |253.6                |64             |1                    |
|GoLogJSONPositive-2    |338.2                |64             |1                    |
|GoLogJSONPositive      |588                  |64             |1                    |


### JSONNegative

|name                   |time/op (ns/op)      |alloc/op (B/op)|allocs/op (allocs/op)|
|-----------------------|---------------------|---------------|---------------------|
|ZerologJSONNegative-4  |1.749                |0              |0                    |
|ZerologJSONNegative-2  |3.257                |0              |0                    |
|ZerologJSONNegative    |6.391                |0              |0                    |
|ZapJSONNegative-4      |24.59                |192            |1                    |
|ZapJSONNegative-2      |34.25                |192            |1                    |
|ZapJSONNegative        |51.37                |192            |1                    |
|LogrusJSONNegative-4   |98.2                 |496            |4                    |
|LogrusJSONNegative-2   |161.2                |496            |4                    |
|LogrusJSONNegative     |286.1                |496            |4                    |
|Log15JSONNegative-4    |238.7                |632            |5                    |
|Log15JSONNegative-2    |423.6                |632            |5                    |
|Log15JSONNegative      |776.8                |632            |5                    |
|GokitJSONNegative-4    |14.55                |128            |1                    |
|GokitJSONNegative-2    |20.91                |128            |1                    |
|GokitJSONNegative      |31.1                 |128            |1                    |
|GoLogJSONNegative-4    |26.71                |64             |1                    |
|GoLogJSONNegative-2    |27.54                |64             |1                    |
|GoLogJSONNegative      |41.9                 |64             |1                    |

