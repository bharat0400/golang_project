package usecase

import (
	"myapp/domain"
	"strconv"
)

type stickerUsecase struct {
	stickerRepo domain.StickerRepository
}

func NewMysqlStickerUsecase(s domain.StickerRepository) domain.StickerUsecase {
	return &stickerUsecase{stickerRepo: s}
}

func (m *stickerUsecase) FetchTrendingStickers() []map[string]string {
	stickers := m.stickerRepo.FetchTrendingStickers()
	var stickers_json []map[string]string
	for _, s := range stickers {
		sticker_json := make(map[string]string)
		sticker_json["id"] = strconv.Itoa(s.Id)
		sticker_json["name"] = s.Name
		sticker_json["priority"] = strconv.Itoa(s.Priority)
		sticker_json["url"] = s.Url
		stickers_json = append(stickers_json, sticker_json)
	}
	return stickers_json
}
