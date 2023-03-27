package bench

import (
	"testing"

	"github.com/millken/golog"
)

func BenchmarkGoLog_TextPositive(b *testing.B) {
	stream := &blackholeStream{}
	cfg := golog.Config{
		Level:    golog.INFO,
		Encoding: golog.TextEncoding,
		TextEncoder: golog.TextEncoderConfig{
			DisableTimestamp: false,
			DisableColor:     true,
		},
		Handler: golog.HandlerConfig{
			Type:   "custom",
			Writer: stream,
		},
	}
	logger, err := golog.NewLoggerByConfig("test", cfg)
	if err != nil {
		b.Fatal(err)
	}
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

func BenchmarkGoLog_TextNegative(b *testing.B) {
	stream := &blackholeStream{}
	cfg := golog.Config{
		Level:    golog.ERROR,
		Encoding: golog.TextEncoding,
		TextEncoder: golog.TextEncoderConfig{
			DisableTimestamp: false,
		},
		Handler: golog.HandlerConfig{
			Type:   "custom",
			Writer: stream,
		},
	}
	logger, err := golog.NewLoggerByConfig("test", cfg)
	if err != nil {
		b.Fatal(err)
	}
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

func BenchmarkGoLog_JSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	cfg := golog.Config{
		Level:    golog.ERROR,
		Encoding: golog.JSONEncoding,
		JSONEncoder: golog.JSONEncoderConfig{
			DisableTimestamp: false,
		},
		Handler: golog.HandlerConfig{
			Type:   "custom",
			Writer: stream,
		},
	}
	logger, err := golog.NewLoggerByConfig("test", cfg)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", "15", "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGoLog_JSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	cfg := golog.Config{
		Level:    golog.INFO,
		Encoding: golog.JSONEncoding,
		JSONEncoder: golog.JSONEncoderConfig{
			DisableTimestamp: false,
		},
		Handler: golog.HandlerConfig{
			Type:   "custom",
			Writer: stream,
		},
	}
	logger, err := golog.NewLoggerByConfig("test", cfg)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog", "rate", "15", "low", 16, "high", 123.2)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
