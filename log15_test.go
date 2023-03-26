package bench

import (
	"testing"

	log "gopkg.in/inconshreveable/log15.v2"
)

func BenchmarkLog15_TextNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.SetHandler(log.LvlFilterHandler(
		log.LvlError,
		log.StreamHandler(stream, log.LogfmtFormat())),
	)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLog15_TextPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.SetHandler(log.StreamHandler(stream, log.LogfmtFormat()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count got %d, want %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkLog15_JSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.SetHandler(log.LvlFilterHandler(
		log.LvlError,
		log.StreamHandler(stream, log.JsonFormat())),
	)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkLog15_JSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := log.New()
	logger.SetHandler(log.StreamHandler(stream, log.JsonFormat()))
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", 15, "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
