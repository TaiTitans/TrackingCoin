//go:build wireinject

package wire

import (
	"github.com/google/wire"

	handler "github.com/TaiTitans/TrackingCoin/internal/handler"
	repository "github.com/TaiTitans/TrackingCoin/internal/repository/user"
	service "github.com/TaiTitans/TrackingCoin/internal/service/auth"
)

func InitAuthRouterHandler() (*handler.AuthHandler, error) {
	wire.Build(
		repository.NewUserRepo,
		service.NewAuthService,
		handler.NewAuthHandler,
	)
	return new(handler.AuthHandler), nil
}
