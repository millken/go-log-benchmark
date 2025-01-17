package bench

import (
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func BenchmarkZerolog_TextPositive(b *testing.B) {
	stream := &blackholeStream{}
	output := zerolog.ConsoleWriter{NoColor: true, Out: stream, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count got %d, want %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkZerolog_TextNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).
		Level(zerolog.ErrorLevel).
		With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZerolog_JSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).
		Level(zerolog.ErrorLevel).
		With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZerolog_JSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	logger := zerolog.New(stream).With().Timestamp().Logger()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info().
				Str("rate", "15").
				Int("low", 16).
				Float32("high", 123.2).
				Msg("The quick brown fox jumps over the lazy dog")
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count")
	}
}
