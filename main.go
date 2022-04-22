package main

import (
	"github.com/gin-gonic/gin"
	"url-shortener/internal/delivery"
	"url-shortener/internal/service/config"
	"url-shortener/internal/service/db"
	"url-shortener/internal/service/logger"
	"url-shortener/internal/service/urld"
	"url-shortener/internal/service/validator"
	"url-shortener/internal/usecase"
)

func main() {
	l := logger.NewLogger("log.txt")

	cfg, err := config.Read()
	if err != nil {
		l.Panicf("read config failure.\n message: %s", err.Error())
	}

	idGenerator := urld.NewGenerator()

	storage, err := db.NewSql(cfg)
	defer storage.Close()
	if err != nil {
		l.Panicf("init sql failure.\n message: %s", err.Error())
	}

	cache := db.NewNosql(cfg)
	defer cache.Close()

	shortenUrlUsecaseImp := usecase.NewShortenUrlUsecase(idGenerator, storage, cache, l)
	redirectUsecaseImp := usecase.NewRedirectUsecase(cache, l)

	router := gin.Default()
	v, err := validator.NewUrlValidator()
	if err != nil {
		l.Panicf("init validator failure.\n message: %s", err.Error())
	}

	delivery.NewDeliverHandler(router, cfg.Server.Host, v, shortenUrlUsecaseImp, redirectUsecaseImp, l)

	router.Run(cfg.Server.Port)
}
