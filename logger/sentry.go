package logger

import (
	"github.com/formeo/XlsForXMLHttp/config"
	"github.com/getsentry/sentry-go"
)

func InitSentry(conf *config.Config) error {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: conf.SentryDSN,
	})
	if err != nil {
		return err
	}
	return nil
}
