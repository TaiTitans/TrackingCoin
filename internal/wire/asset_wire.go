//go:build wireinject

package wire

import (
	"github.com/google/wire"

	"github.com/TaiTitans/TrackingCoin/internal/handler"
	assetRepo "github.com/TaiTitans/TrackingCoin/internal/repository/asset"
	assetService "github.com/TaiTitans/TrackingCoin/internal/service/asset"
)

func InitAssetRouterHandler() (*handler.AssetHandler, error) {
	wire.Build(
		assetRepo.NewAssetRepo,
		assetService.NewAssetService,
		handler.NewAssetHandler,
	)
	return new(handler.AssetHandler), nil
}
