package main

import (
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
    "sakata-server-rest/internal/di"
    "time"
)

func main() {
    r := gin.Default()

    // CORS 設定
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"}, // フロントエンドのURL
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // 依存関係の初期化
    userHandler := di.InitializeUserHandler()
    themeHandler := di.InitializeThemeHandler()

    // ユーザー関連のルーティング
    r.GET("/user/:id", userHandler.GetUser)

    // テーマ関連のルーティング
    r.POST("/theme", themeHandler.CreateTheme)
    r.GET("/theme/:id", themeHandler.GetTheme)

    r.Run(":8085") // サーバー起動
}
