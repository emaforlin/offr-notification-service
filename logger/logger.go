package logger

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func ProvideZapLogger() fx.Option {
	return fx.Provide(zap.NewExample)
}
