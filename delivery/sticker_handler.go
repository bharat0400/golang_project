package delivery

import (
	"myapp/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type stickerHandler struct {
	sUsecase domain.StickerUsecase
}

func NewStickerHandler(e *echo.Echo, us domain.StickerUsecase) {
	handler := &stickerHandler{
		sUsecase: us,
	}
	stickers_json := handler.sUsecase.FetchTrendingStickers()
	e.GET("/v1/stickers", func(c echo.Context) error {
		return c.JSON(http.StatusOK, stickers_json)
	})
}
