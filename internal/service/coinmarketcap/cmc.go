package coinmarketcap

import (
	"github.com/gin-gonic/gin"

	"github.com/TaiTitans/TrackingCoin/internal/dto"
	repository "github.com/TaiTitans/TrackingCoin/internal/repository/coinmarketcap"
)

type ICMCService interface {
	GetDetailFromCMC(ctx *gin.Context, slug string) (*dto.Currency, error)
	GetMarketPairFromCMC(ctx *gin.Context, slug string) (*dto.MarketPairsResponse, error)
	GetQuoteLastestFromCMC(ctx *gin.Context, id int) (*dto.QuoteLastestResponse, error)
}

type cmcService struct {
	repo repository.ICMCRepository
}

func NewCMCService(repo repository.ICMCRepository) ICMCService {
	return &cmcService{
		repo: repo,
	}
}
