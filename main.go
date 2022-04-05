package main

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/delivery"
	"url-shortener/internal/service/config"
	"url-shortener/internal/service/db"
	"url-shortener/internal/service/urld"
	"url-shortener/internal/usecase"
)

func main() {
	cfg := config.Read()

	idGenerator := urld.NewGenerator()
	storage := db.NewSql(cfg)
	cache := db.NewNosql(cfg)
	shortenUrlUsecaseImp := usecase.NewShortenUrlUsecase(idGenerator, storage, cache)
	redirectUsecaseImp := usecase.NewRedirectUsecase(cache)

	defer storage.Close()
	defer cache.Close()

	router := gin.Default()
	delivery.NewDeliverHandler(router, cfg.Server.Host, shortenUrlUsecaseImp, redirectUsecaseImp)

	router.Run(cfg.Server.Port)
}