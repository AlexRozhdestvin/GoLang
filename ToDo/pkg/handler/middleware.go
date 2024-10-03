package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = `Authorization`
	userCTX             = `userId`
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorMessage(c, http.StatusUnauthorized, "empty auth header")
		return
	}
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorMessage(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])

	if err != nil {
		newErrorMessage(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCTX, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCTX)

	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)

	if !ok {
		newErrorMessage(c, http.StatusInternalServerError, "user id is of invalid type")
		return 0, errors.New("user id not found")
	}

	return idInt, nil
}
