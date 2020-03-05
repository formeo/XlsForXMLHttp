package logger

import (
	"github.com/formeo/XlsForXMLHttp/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(conf *config.Config) (*zap.Logger, error) {
	level := zap.NewAtomicLevel()
	err := level.UnmarshalText([]byte(conf.LogLevel))
	if err != nil {
		return nil, err
	}
	logEncoding := "json"
	logEncodeLevel := zapcore.LowercaseLevelEncoder

	if conf.DevMode {
		logEncoding = "console"
		logEncodeLevel = zapcore.LowercaseColorLevelEncoder
	}

	logConfig := zap.Config{
		Level:       level,
		Development: conf.DevMode,
		Encoding:    logEncoding,
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "severity",
			TimeKey:        "timestamp",
			EncodeLevel:    logEncodeLevel,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
		},
		OutputPaths: []string{"stdout"},
	}
	log, err := logConfig.Build()
	if err != nil {
		return nil, err
	}
	log = log.With(getFields(conf)...)

	return log, nil
}

func getFields(conf *config.Config) []zapcore.Field {
	var fields []zapcore.Field

	if conf.Index != "" {
		fields = append(fields, zap.String("index", conf.Index))
	}

	return fields
}
