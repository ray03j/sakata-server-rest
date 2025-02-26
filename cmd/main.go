package main

import (
    "github.com/gin-gonic/gin"
    "sakata-server-rest/internal/di"
)

func main() {
    r := gin.Default()

    userHandler := di.InitializeUserHandler()

    // ルーティング設定
    r.GET("/user/:id", userHandler.GetUser)

    r.Run(":8085") // サーバー起動
}
