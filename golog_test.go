package bench

import (
	"testing"

	log "github.com/millken/golog"
)

func BenchmarkGoLogTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	handler := log.NewLoggerHandler(stream)
	formatter := log.NewTextFormatter()
	formatter.NoColor = true
	formatter.DisableTimestamp = true
	handler.SetFormatter(formatter)
	handler.SetLevel(log.InfoLevel)
	logger := log.NewLogger()
	logger.AddHandler(handler)
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

func BenchmarkGoLogTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	handler := log.NewLoggerHandler(stream)
	formatter := log.NewTextFormatter()
	formatter.NoColor = true
	formatter.DisableTimestamp = true
	handler.SetFormatter(formatter)
	handler.SetLevel(log.ErrorLevel)
	logger := log.NewLogger()
	logger.AddHandler(handler)
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

func BenchmarkGoLogJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	handler := log.NewLoggerHandler(stream)
	formatter := log.NewJSONFormatter()
	handler.SetFormatter(formatter)
	handler.SetLevel(log.ErrorLevel)
	logger := log.NewLogger()
	logger.AddHandler(handler)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.
				WithFields(log.F("rate", "15"), log.F("low", 16), log.F("high", 123.2)).
				Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGoLogJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	handler := log.NewLoggerHandler(stream)
	formatter := log.NewJSONFormatter()
	handler.SetFormatter(formatter)
	handler.SetLevel(log.InfoLevel)
	logger := log.NewLogger()
	logger.AddHandler(handler)
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.
				WithFields(log.F("rate", "15"), log.F("low", 16), log.F("high", 123.2)).
				Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
