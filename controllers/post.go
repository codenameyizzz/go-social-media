package controllers

import (
	"fmt"
	"go-socialmedia/models"
	"net/http"
	"path/filepath"
	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadPost(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	userID := userIDVal.(int)

	caption := c.PostForm("caption")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Gambar wajib diunggah"})
		return
	}

	filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
	err = c.SaveUploadedFile(file, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan gambar"})
		return
	}

	_, err = models.DB.Exec("INSERT INTO posts (user_id, caption, image_path) VALUES (?, ?, ?)", userID, caption, filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "Post berhasil diupload",
		"imagePath": filename,
	})
}

func GetAllPosts(c *gin.Context) {
	rows, err := models.DB.Query(`
			SELECT posts.id, posts.user_id, posts.caption, posts.image_path, posts.created_at, users.username
			FROM posts
			JOIN users ON posts.user_id = users.id
			ORDER BY posts.created_at DESC
	`)

	row := models.DB.QueryRow("SELECT COUNT(*) FROM posts")
	var total int
	row.Scan(&total)
	fmt.Println("JUMLAH POST TERDETEKSI:", total)

	if err != nil {
			fmt.Println("DB Query error:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
	}
	defer rows.Close()

	var posts []gin.H

	for rows.Next() {
			var id, userID int
			var caption, imagePath, username string
			var createdAt time.Time

			err := rows.Scan(&id, &userID, &caption, &imagePath, &createdAt, &username)
			if err != nil {
					fmt.Println("Scan error:", err)
					continue
			}

			posts = append(posts, gin.H{
					"id":         id,
					"user_id":    userID,
					"username":   username,
					"caption":    caption,
					"image":      "/" + imagePath,
					"created_at": createdAt.Format("2006-01-02 15:04"),
			})
	}

	if posts == nil {
			posts = []gin.H{}
	}
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
