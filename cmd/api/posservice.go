package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"keycard_service/internal/database"
	"net/http"
	"time"
)

type CreateKeyServiceRequest struct {
	UserID     uint   `json:"user_id" binding:"required"`     // کاربر صاحب کارت
	UserName   string `json:"user_name" binding:"required"`   // نام کاربر
	Email      string `json:"email"`                          // اختیاری
	Phone      string `json:"phone"`                          // اختیاری
	TerminalID uint   `json:"terminal_id" binding:"required"` // دستگاه POS
	KeyEnc     string `json:"key_enc" binding:"required"`     // کارت رمزنگاری شده
	KeyType    string `json:"key_type" binding:"required"`    // نوع کلید کارت (PIN, Track)
}

func (h *KeyCardHandler) CreateKeyService(c *gin.Context) {
	var req CreateKeyServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 1️⃣ ذخیره یا به‌روزرسانی اطلاعات کاربر
	user := database.User{
		ID:       req.UserID,
		Username: req.UserName,
		Email:    req.Email,
		Phone:    req.Phone,
		Status:   "ACTIVE",
	}
	if err := h.repo.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 2️⃣ ذخیره کلید کارت مرتبط با کاربر و دستگاه POS
	key := database.KeyCard{
		UserID:     req.UserID,
		TerminalID: req.TerminalID,
		KeyEnc:     []byte(req.KeyEnc), // رمزنگاری واقعی بهتر است
		KeyType:    req.KeyType,
		CreatedAt:  time.Now(),
	}
	if err := h.repo.DB.Create(&key).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User and Key saved successfully", "key_id": key.ID})
}
func (h *KeyCardHandler) GetAllKeyService(c *gin.Context) {
	var keys []database.KeyCard
	if err := h.repo.DB.Preload("User").Preload("Terminal").Find(&keys).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"keys": keys})
}
