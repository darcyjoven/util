package zapper

import (
	"encoding/json"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type constantClock time.Time

func (c constantClock) Now() time.Time { return time.Time(c) }
func (c constantClock) NewTicker(d time.Duration) *time.Ticker {
	return &time.Ticker{}
}

func Exec(file string) (logger *zap.Logger, err error) {
	rawJSON := []byte(`{
		"level": "debug",
		"encoding": "console"
	  }`)
	var cfg zap.Config

	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	cfg.OutputPaths = append(cfg.OutputPaths, "stdout", file)
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("06-01-02 15:04:05.00")
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	local, _ := time.LoadLocation("Asia/Shanghai")
	date := time.Now().In(local)
	clock := constantClock(date)
	logger = zap.Must(cfg.Build(
		zap.WithClock(clock),
	))
	// defer logger.Sync()
	return logger, nil
}
