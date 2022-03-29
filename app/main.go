package main

import (
	"database/sql"
	"log"
	_delivery "myapp/delivery"
	_repo "myapp/sticker/repository/mysql"
	_usecase "myapp/sticker/usecase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		log.Fatal("Unable to open connection to db")
	}
	defer db.Close()
	e := echo.New()
	repo := _repo.NewMysqlStickerRepository(db)
	usecase := _usecase.NewMysqlStickerUsecase(repo)
	_delivery.NewStickerHandler(e, usecase)
	e.Logger.Fatal(e.Start(":3000"))
}
