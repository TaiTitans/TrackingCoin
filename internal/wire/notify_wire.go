//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/TaiTitans/TrackingCoin/internal/handler"
	"github.com/TaiTitans/TrackingCoin/internal/service/notification"
)

func InitNotifyRouterHandler() (*handler.NotifyHandler, error) {
	wire.Build(
		notification.NewNotifyService,
		handler.NewNotifyHandler,
	)
	return new(handler.NotifyHandler), nil
}
