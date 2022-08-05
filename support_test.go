package bench

/*
	go test -benchmem -benchtime=5s -bench "Benchmark.*TextNegative" |tee TextNegative.txt
	benchstat -csv -sort -name TextNegative.txt > TextNegative.csv
  https://www.convertcsv.com/csv-to-markdown.htm
*/
import "sync/atomic"

type blackholeStream struct {
	writeCount uint64
}

func (s *blackholeStream) WriteCount() uint64 {
	return atomic.LoadUint64(&s.writeCount)
}

func (s *blackholeStream) Write(p []byte) (int, error) {
	atomic.AddUint64(&s.writeCount, 1)
	return len(p), nil
}
