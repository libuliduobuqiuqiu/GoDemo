package zapdemo

import (
	"time"

	"go.uber.org/zap"
)

func UseZapLogging() {
	logger, _ := zap.NewProduction()

	defer logger.Sync()
	suger := logger.Sugar()

	url := "http://www.google.com"

	suger.Infow(
		"failed to fetch the URL",
		"Url", url,
		"Attempt", 3,
		"Backoff", time.Second,
	)
}

func UseFasterZapLogging() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	url := "http://www.google.com"
	logger.Info("failed to featch the Url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

func UseZapExampleLogging() {
	logger := zap.NewExample()
	defer logger.Sync()

	url := "http://www.google.com"
	logger.Info("failed to featch the Url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)

}