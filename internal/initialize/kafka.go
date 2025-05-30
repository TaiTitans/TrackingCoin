package initialize

import (
	"github.com/TaiTitans/TrackingCoin/global"
	"github.com/TaiTitans/TrackingCoin/internal/kafka/consumer"
	"github.com/TaiTitans/TrackingCoin/internal/kafka/producer"
)

func InitKafka() {
	producer := producer.NewProducer([]string{"localhost:9092"})
	global.Logger.Info("init producer successfully")
	defer producer.Close()

	consumer := consumer.NewConsumer(global.Config.Kafka.GroupID, []string{"notification_topic"})
	global.Logger.Info("init consumer successfully")
	defer consumer.Close()
}
