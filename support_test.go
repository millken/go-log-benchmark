package bench

/*
	go test -benchmem -benchtime=5s -bench "Benchmark.*TextNegative" |tee TextNegative.txt
	benchstat -csv -sort -name TextNegative.txt > TextNegative.csv
  https://www.convertcsv.com/csv-to-markdown.htm
*/
import (
	"os"
	"sync/atomic"
	"testing"
	"time"

	kit "github.com/go-kit/kit/log"
	"github.com/millken/golog"
	"github.com/rs/zerolog"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

func TestJSONPositive(t *testing.T) {
	t.Run("Zerolog", func(t *testing.T) {
		logger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
		logger.Info().
			Str("rate", "15").
			Int("low", 16).
			Float32("high", 123.2).
			Msg("The quick brown fox jumps over the lazy dog")
	})
	t.Run("Zap", func(t *testing.T) {
		atom := zap.NewAtomicLevel()
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		logger := zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		), zap.WithCaller(false))
		logger.Info("The quick brown fox jumps over the lazy dog",
			zap.String("rate", "15"),
			zap.Int("low", 16),
			zap.Float32("high", 123.2),
		)
	})
	t.Run("Golog", func(t *testing.T) {
		cfg := golog.Config{
			Level:    golog.INFO,
			Encoding: golog.JSONEncoding,
			JSONEncoder: golog.JSONEncoderConfig{
				DisableTimestamp: false,
			},
			Handler: golog.HandlerConfig{
				Type:   "custom",
				Writer: os.Stdout,
			},
			CallerLevels: []golog.Level{golog.INFO},
		}
		logger, _ := golog.NewLoggerByConfig("test", cfg)
		logger.Info("The quick brown fox jumps over the lazy dog", "rate", "15", "low", 16, "high", 123.2)
	})
	t.Run("Logrus", func(t *testing.T) {
		logger := logrus.New()
		logger.Formatter = &logrus.JSONFormatter{}
		logger.Out = os.Stdout
		logger.WithFields(logrus.Fields{
			"rate": "15",
			"low":  16,
			"high": 123.2,
		}).Info("The quick brown fox jumps over the lazy dog")
	})
	t.Run("Gokit", func(t *testing.T) {
		logger := kit.With(kit.NewJSONLogger(os.Stdout), "ts", kit.DefaultTimestampUTC)
		lvllog := newLeveledLogger(logger, true)
		lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog", "rate", "15", "low", 16, "high", 123.2)
	})
}

func TestTextPositive(t *testing.T) {
	t.Run("Zerolog", func(t *testing.T) {
		output := zerolog.ConsoleWriter{NoColor: true, Out: os.Stdout, TimeFormat: time.RFC3339}
		logger := zerolog.New(output).With().Timestamp().Logger()
		logger.Info().
			Str("rate", "15").
			Int("low", 16).
			Float32("high", 123.2).
			Msg("The quick brown fox jumps over the lazy dog")
	})
	t.Run("Zap", func(t *testing.T) {
		atom := zap.NewAtomicLevel()
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		logger := zap.New(zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoderCfg),
			zapcore.Lock(os.Stdout),
			atom,
		), zap.WithCaller(callerEnabled))
		logger.Info("The quick brown fox jumps over the lazy dog",
			zap.String("rate", "15"),
			zap.Int("low", 16),
			zap.Float32("high", 123.2),
		)
	})
	t.Run("Golog", func(t *testing.T) {
		cfg := golog.Config{
			Level:    golog.INFO,
			Encoding: golog.TextEncoding,
			TextEncoder: golog.TextEncoderConfig{
				DisableTimestamp: false,
				DisableColor:     true,
			},
			Handler: golog.HandlerConfig{
				Type:   "custom",
				Writer: os.Stdout,
			},
			//CallerLevels: []golog.Level{golog.INFO},
		}
		logger, _ := golog.NewLoggerByConfig("test", cfg)
		logger.Info("The quick brown fox jumps over the lazy dog", "rate", "15", "low", 16, "high", 123.2)
	})
	t.Run("Logrus", func(t *testing.T) {
		logger := logrus.New()
		logger.Formatter = &logrus.TextFormatter{}
		logger.Out = os.Stdout
		logger.WithFields(logrus.Fields{
			"rate": "15",
			"low":  16,
			"high": 123.2,
		}).Info("The quick brown fox jumps over the lazy dog")
	})
	t.Run("Gokit", func(t *testing.T) {
		logger := kit.With(kit.NewLogfmtLogger(os.Stdout), "ts", kit.DefaultTimestampUTC)
		lvllog := newLeveledLogger(logger, true)
		lvllog.Info.Log("msg", "The quick brown fox jumps over the lazy dog", "rate", "15", "low", 16, "high", 123.2)
	})
}
