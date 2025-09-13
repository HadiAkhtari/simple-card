package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"keycard_service/internal/database"
	"net/http"
	"time"
)

// KeyCardHandler یک struct است برای کار با KeyCard
type KeyCardHandler struct {
	repo *database.KeyCardModel // این repo دسترسی به دیتابیس را فراهم می‌کند
}

// تابع سازنده
func NewKeyCardHandler(repo *database.KeyCardModel) *KeyCardHandler {
	return &KeyCardHandler{repo: repo}
}

type CreatePOSRequest struct {
	User     UserRequest     `json:"user" binding:"required"`
	Terminal TerminalRequest `json:"terminal" binding:"required"`
	KeyCard  KeyCardRequest  `json:"keycard" binding:"required"`
}

type UserRequest struct {
	ID       uint   `json:"id"`
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type TerminalRequest struct {
	ID           uint   `json:"id"`
	SerialNumber string `json:"serial_number" binding:"required"`
	Model        string `json:"model"`
}

type KeyCardRequest struct {
	KeyEnc  string `json:"key_enc" binding:"required"`
	KeyType string `json:"key_type" binding:"required"`
}

func (h *KeyCardHandler) CreatePOSData(c *gin.Context) {
	var req CreatePOSRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1️⃣ ذخیره یا آپدیت کاربر
	user := database.User{
		ID:     req.User.ID,
		Status: "ACTIVE",
	}
	if err := h.repo.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2️⃣ ذخیره یا آپدیت ترمینال
	terminal := database.Terminal{
		ID:           req.Terminal.ID,
		UserID:       user.ID, // اگر می‌خوای ترمینال به کاربر وصل باشه
		SerialNumber: req.Terminal.SerialNumber,
		Model:        req.Terminal.Model,
		Status:       "ACTIVE",
	}
	if err := h.repo.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&terminal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 3️⃣ ذخیره کلید کارت
	key := database.KeyCard{
		UserID:     user.ID,
		TerminalID: terminal.ID,
		CreatedAt:  time.Now(),
	}
	if err := h.repo.DB.Create(&key).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":     "User, Terminal and Key saved successfully",
		"user_id":     user.ID,
		"terminal_id": terminal.ID,
		"key_id":      key.ID,
	})
}
