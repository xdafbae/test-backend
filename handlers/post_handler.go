package handlers

import (
	"net/http"
	"strconv"

	"sharingvision-backend/config"
	"sharingvision-backend/models"

	"github.com/gin-gonic/gin"
)

// PostRequest — request body dengan validasi
type PostRequest struct {
	Title    string `json:"title"    binding:"required,min=20"`
	Content  string `json:"content"  binding:"required,min=200"`
	Category string `json:"category" binding:"required,min=3"`
	Status   string `json:"status"   binding:"required,oneof=publish draft thrash"`
}

// CreateArticle — POST /article
func CreateArticle(c *gin.Context) {
	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title:    req.Title,
		Content:  req.Content,
		Category: req.Category,
		Status:   req.Status,
	}

	if err := config.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	c.JSON(http.StatusCreated, post)
}

// GetAllArticles — GET /article/:limit/:offset
func GetAllArticles(c *gin.Context) {
	limitStr := c.Param("limit")
	offsetStr := c.Param("offset")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	var posts []models.Post
	if err := config.DB.Limit(limit).Offset(offset).Order("created_date DESC").Find(&posts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch articles"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// GetArticleByID — GET /article/:id
func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// UpdateArticle — PUT /article/:id
func UpdateArticle(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	var req PostRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Category = req.Category
	post.Status = req.Status

	if err := config.DB.Save(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	c.JSON(http.StatusOK, post)
}

// DeleteArticle — DELETE /article/:id
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	var post models.Post
	if err := config.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if err := config.DB.Delete(&post).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
