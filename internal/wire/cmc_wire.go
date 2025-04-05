//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/TaiTitans/TrackingCoin/internal/handler"
	cmcRepo "github.com/TaiTitans/TrackingCoin/internal/repository/coinmarketcap"
	cmcService "github.com/TaiTitans/TrackingCoin/internal/service/coinmarketcap"
)

func InitCMCRouterHandler() (*handler.CMCHandler, error) {
	wire.Build(
		cmcRepo.NewCMCRepo,
		cmcService.NewCMCService,
		handler.NewCMCHandler,
	)
	return new(handler.CMCHandler), nil
}
