package global

import (
	"FXIm/config"
	"go.uber.org/zap"
)

var (
	FXConfig config.Config
	FXLog    *zap.Logger
)
