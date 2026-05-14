package core_logger

import "go.uber.org/zap"

func New(isProd bool) (*zap.Logger, error) {
	if isProd {
		return zap.NewProduction()
	}
	return zap.NewDevelopment()
}
