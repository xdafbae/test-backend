package routes

import (
	"sharingvision-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORS Middleware — izinkan semua origin (dev mode)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))

	// Routes
	r.POST("/article", handlers.CreateArticle)
	r.GET("/article/:limit/:offset", handlers.GetAllArticles)
	r.GET("/article/:id", handlers.GetArticleByID)
	r.PUT("/article/:id", handlers.UpdateArticle)
	r.DELETE("/article/:id", handlers.DeleteArticle)

	return r
}
