package zapper

import (
	"encoding/json"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

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

	logger = zap.Must(cfg.Build())
	// defer logger.Sync()
	return logger, nil
}
