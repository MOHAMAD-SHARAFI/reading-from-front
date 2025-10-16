package handlers

import (
	"net/http"
	"project/internal/config"
	"project/internal/models"
	"project/internal/usecase"

	"github.com/gin-gonic/gin"
)

type handler struct {
	uc *usecase.Usecase
}

func Init(conf *config.Config, uc *usecase.Usecase) (err error) {
	h := &handler{
		uc: uc,
	}

	server := gin.Default()
	server.POST("/users", h.createUser)
	return server.Run(conf.ListenAddr)
}

func (h *handler) createUser(body *gin.Context) {
	var user models.User
	if err := body.ShouldBindJSON(&user); err != nil {
		body.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.uc.CreateUser(user)
	if err != nil {
		body.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	body.JSON(http.StatusOK, gin.H{"message": "user created successfully", "user": resp})

}
