package bench

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	callerEnabled = false
)

func BenchmarkZap_TextPositive(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	logger := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderCfg),
		w,
		zap.InfoLevel,
	), zap.WithCaller(callerEnabled))

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count got %d, want %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkZap_TextNegative(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.ErrorLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}

func BenchmarkZap_JSONPositive(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})

	if stream.WriteCount() != uint64(b.N) {
		b.Fatalf("Log write count got %d, want %d", stream.WriteCount(), b.N)
	}
}

func BenchmarkZap_JSONNegative(b *testing.B) {
	stream := &blackholeStream{}
	w := zapcore.AddSync(stream)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.ErrorLevel,
	)
	logger := zap.New(core)

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Info("The quick brown fox jumps over the lazy dog",
				zap.String("rate", "15"),
				zap.Int("low", 16),
				zap.Float32("high", 123.2),
			)
		}
	})

	if stream.WriteCount() != uint64(0) {
		b.Fatalf("Log write count")
	}
}
