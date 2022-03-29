package domain

type Sticker struct {
	Id           int
	Name         string
	Priority     int
	Start_hour   int
	End_hour     int
	Published_at string
	Created_at   string
	Updated_at   string
	Url          string
}

type StickerUsecase interface {
	FetchTrendingStickers() []map[string]string
}

type StickerRepository interface {
	FetchTrendingStickers() (result []Sticker)
}
