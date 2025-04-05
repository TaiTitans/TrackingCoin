package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"

	database "github.com/TaiTitans/TrackingCoin/database/sqlc"
	"github.com/TaiTitans/TrackingCoin/pkg/logger"
	"github.com/TaiTitans/TrackingCoin/pkg/setting"
)

var (
	Config        setting.Config
	Logger        *logger.LoggerZap
	Db            *database.Store
	Rdb           *redis.Client
	KafkaProducer *kafka.Writer
	KafkaConsumer *kafka.Reader
)
