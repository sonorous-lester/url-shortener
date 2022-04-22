package delivery

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"url-shortener/internal/domain"
	"url-shortener/internal/exception"
)

type originURL struct {
	Url      string `json:"url"`
	ExpireAt string `json:"expireAt"`
}

type shortURL struct {
	Id       string `json:"id"`
	ShortUrl string `json:"shortUrL"`
}

type Handler struct {
	shortDomainUrl    string
	validator         domain.UrlValidator
	shortenUrlUsecase domain.ShortenUrlUsecase
	redirectUsecase   domain.RedirectUsecase
	logger            domain.Logger
}

func NewDeliverHandler(
	c *gin.Engine,
	shortDomainUrl string,
	validator domain.UrlValidator,
	shortenUrl domain.ShortenUrlUsecase,
	redirect domain.RedirectUsecase,
	logger domain.Logger) {
	handler := &Handler{shortDomainUrl: shortDomainUrl, shortenUrlUsecase: shortenUrl, validator: validator, redirectUsecase: redirect, logger: logger}
	c.POST("/api/v1/urls", handler.urls)
	c.GET("/:url_id", handler.redirect)
}

func (h *Handler) urls(c *gin.Context) {
	var originUrl originURL
	err := c.Bind(&originUrl)
	if err != nil {
		raw, _ := c.GetRawData()
		h.logger.Debugf("convert request body to origin url failure. raw data is: %s", raw)
		c.JSON(http.StatusBadRequest, gin.H{"message": exception.IncorrectInput.Error()})
		return
	}

	isValid := h.validator.Valid(originUrl.Url)
	if !isValid {
		h.logger.Warnf("url is not valid. url is: %s", originUrl.Url)
		c.JSON(http.StatusBadRequest, gin.H{"message": exception.URLIsNotValid.Error()})
		return
	}

	urlId, shortUrl, err := h.shortenUrlUsecase.Process(h.shortDomainUrl, originUrl.Url, originUrl.ExpireAt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, shortURL{
		Id:       urlId,
		ShortUrl: shortUrl,
	})
}

func (h *Handler) redirect(c *gin.Context) {
	id := c.Param("url_id")
	location, err := h.redirectUsecase.Process(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if location == "" {
		c.JSON(http.StatusNotFound, gin.H{"message": "not found any matched short url"})
		return
	}
	c.Redirect(http.StatusFound, location)
}
