package bench

import (
	"log"
	"testing"
)

func BenchmarkBuildinTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New(stream, "", 0)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Print("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
