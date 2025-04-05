package initialize

import (
	"fmt"

	"github.com/TaiTitans/TrackingCoin/global"
	"github.com/TaiTitans/TrackingCoin/internal/cronjob"
)

func Run() {
	LoadConfig()
	InitLogger()
	InitDB()
	//InitRedis()
	InitKafka()
	go cronjob.StartHourlyJob()
	r := InitRouter()
	serverAddr := fmt.Sprintf(":%v", global.Config.Server.Port)
	r.Run(serverAddr)
}
