package initialize

import (
	"github.com/TaiTitans/TrackingCoin/global"
	"github.com/TaiTitans/TrackingCoin/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
