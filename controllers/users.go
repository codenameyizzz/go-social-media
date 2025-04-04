package controllers

import (
	// "database/sql"
	"go-socialmedia/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	// "github.com/golang-jwt/jwt/v5"
	// "golang.org/x/crypto/bcrypt"
)

func GetMyProfile(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(int)

	var username, email string
	err := models.DB.QueryRow("SELECT username, email FROM users WHERE id = ?", userID).Scan(&username, &email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data user"})
		return
	}

	// Ambil semua post user juga
	rows, err := models.DB.Query(`
		SELECT id, caption, image_path, created_at
		FROM posts
		WHERE user_id = ?
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal ambil post user"})
		return
	}
	defer rows.Close()

	var posts []gin.H
	for rows.Next() {
		var id int
		var caption, imagePath string
		var createdAt time.Time

		err := rows.Scan(&id, &caption, &imagePath, &createdAt)
		if err != nil {
			continue
		}

		posts = append(posts, gin.H{
			"id":         id,
			"caption":    caption,
			"image":      "/" + imagePath,
			"created_at": createdAt.Format("2006-01-02 15:04"),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"email":    email,
		"posts":    posts,
	})
}
