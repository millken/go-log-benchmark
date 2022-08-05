package bench

import (
	"testing"

	"github.com/millken/golog"
	"github.com/millken/golog/config"
	"github.com/millken/golog/log"
)

func BenchmarkGoLogTextPositive(b *testing.B) {
	stream := &blackholeStream{}
	cfg := config.Config{
		Level:    log.INFO,
		Encoding: "console",
		ConsoleEncoderConfig: config.ConsoleEncoderConfig{
			DisableTimestamp: true,
		},
		Writer: config.WriterConfig{
			Type:         "custom",
			CustomWriter: stream,
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

func BenchmarkGoLogTextNegative(b *testing.B) {
	stream := &blackholeStream{}
	cfg := config.Config{
		Level:    log.ERROR,
		Encoding: "console",
		ConsoleEncoderConfig: config.ConsoleEncoderConfig{
			DisableTimestamp: true,
		},
		Writer: config.WriterConfig{
			Type:         "custom",
			CustomWriter: stream,
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

func BenchmarkGoLogJSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	cfg := config.Config{
		Level:    log.ERROR,
		Encoding: "json",
		JSONEncoderConfig: config.JSONEncoderConfig{
			DisableTimestamp: true,
		},
		Writer: config.WriterConfig{
			Type:         "custom",
			CustomWriter: stream,
		},
	}
	logger, err := golog.NewLoggerByConfig("test", cfg)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.
				WithFields(golog.F("rate", "15"), golog.F("low", 16), golog.F("high", 123.2)).
				Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkGoLogJSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	cfg := config.Config{
		Level:    log.INFO,
		Encoding: "json",
		JSONEncoderConfig: config.JSONEncoderConfig{
			DisableTimestamp: true,
		},
		Writer: config.WriterConfig{
			Type:         "custom",
			CustomWriter: stream,
		},
	}
	logger, err := golog.NewLoggerByConfig("test", cfg)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.
				WithFields(golog.F("rate", "15"), golog.F("low", 16), golog.F("high", 123.2)).
				Info("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
