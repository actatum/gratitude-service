package logger

import (
	"go.uber.org/zap"
)

// NewZapLogger returns a new zap logger object
func NewZapLogger() (*zap.Logger, error) {
	l, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return l, nil
}
