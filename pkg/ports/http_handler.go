package ports

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	appError "urlShortener/error"
	"urlShortener/pkg/app/command"
	"urlShortener/pkg/utils"

	"github.com/gin-gonic/gin"
)

type shorteningUrlRequestBody struct {
	OriginalUrl string `json:"origianlUrl"`
}

func (s HttpServer) HandleShorteningLink(c *gin.Context) {
	username := c.GetString("username")
	if username == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	var body shorteningUrlRequestBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	su, err := s.app.Commands.CreateShortLink.Handle(c.Request.Context(), command.CreateShortLink{
		OriginalUrl: body.OriginalUrl,
		Name:        username,
	})
	if err != nil {
		handleAppError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"newUrl": su.FullShortUrl(),
	})

}

func (s HttpServer) HandleGetOriginalLink(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	originalLink, err := s.app.Queries.GetOriginalLink.Handle(c.Request.Context(), id)
	if err != nil {
		handleAppError(c, err)
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, originalLink)
}

type RegisterForm struct {
	Name string `json:"name"`
}

func (s HttpServer) RegisterOrLogin(c *gin.Context) {
	var form RegisterForm
	err := c.ShouldBindJSON(&form)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = s.app.Commands.Register.Handle(c.Request.Context(), form.Name)
	if errors.Is(err, appError.ErrorAlreadyExists{}) {
		//Do Nothing
	} else if err != nil {
		handleAppError(c, err)
		return
	}

	token, err := utils.GenerateToken(form.Name)
	if err != nil {
		err = fmt.Errorf("error generating token: %w", err)
		log.Println(err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

func (s HttpServer) authorize(c *gin.Context) {
	tokenStr := c.GetHeader("authorization")
	if tokenStr == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	username, err := utils.ParseToken(tokenStr)
	if err != nil {
		stdErr := fmt.Errorf("error parsing token: %w", err)
		c.AbortWithError(http.StatusForbidden, stdErr)
		return
	}

	u, err := s.app.Queries.GetUser.Handle(c.Request.Context(), username)
	if err != nil {
		handleAppError(c, err)
		return
	}

	c.Set("username", u.Name())
}

func handleAppError(c *gin.Context, err error) {
	if e, has := err.(appError.AppError); has {
		c.AbortWithStatus(e.HttpStatus())
		return
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}
}
