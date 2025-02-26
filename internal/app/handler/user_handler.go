package handler

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "sakata-server-rest/internal/app/service"
)

// UserHandler は HTTP リクエストを処理
type UserHandler struct {
    userService service.UserService
}

// NewUserHandler は UserHandler のインスタンスを作成
func NewUserHandler(userService service.UserService) *UserHandler {
    return &UserHandler{userService: userService}
}

// GetUser は /user/:id に対するリクエストを処理
func (h *UserHandler) GetUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    userName := h.userService.GetUserName(id)
    c.JSON(http.StatusOK, gin.H{"id": id, "name": userName})
}
