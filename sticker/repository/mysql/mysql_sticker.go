package mysql

import (
	"database/sql"
	"log"
	"myapp/domain"
	"time"

	"github.com/spf13/viper"
)

var (
	configName  = "stickers"
	configPaths = []string{
		".",
	}
)

type mysqlStickerRepository struct {
	Conn *sql.DB
}

func NewMysqlStickerRepository(Conn *sql.DB) domain.StickerRepository {
	return &mysqlStickerRepository{Conn}
}

func (m *mysqlStickerRepository) FetchTrendingStickers() (result []domain.Sticker) {
	var stickers []domain.Sticker
	results, err := m.Conn.Query("select * from stickers where end_hour > ? and start_hour < ? order by priority desc", time.Now().Hour(), time.Now().Hour())
	if err != nil {
		log.Fatal("Error when fetching stickers table rows:", err)
	}
	defer results.Close()
	for results.Next() {
		var sticker domain.Sticker
		err = results.Scan(&sticker.Id, &sticker.Name, &sticker.Priority, &sticker.Start_hour, &sticker.End_hour, &sticker.Published_at, &sticker.Created_at, &sticker.Updated_at)
		if err != nil {
			log.Fatal("Unable to parse row:", err)
		}
		stickers = append(stickers, sticker)
	}
	viper.SetConfigName(configName)
	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}
	err2 := viper.ReadInConfig()
	if err2 != nil {
		log.Fatalf("could not read config file: %v", err2)
	}
	for i := 0; i < len(stickers); i++ {
		stickers[i].Url = viper.GetString("stickers.raw_url") + stickers[i].Name
	}
	return stickers
}
