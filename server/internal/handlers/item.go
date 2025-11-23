package handlers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/BimaPDev/MixMatch/db"
	"github.com/gin-gonic/gin"
)

type createItemRequest struct {
	OwnerID  int32  `json:"owner_id" binding:"required"`
	ImageURL string `json:"image_url" binding:"required"`
	Category string `json:"category"`
	Color    string `json:"color"`
	Style    string `json:"style"`
}

// CreateItem handles POST /items
func (h *Handler) CreateItem(ctx *gin.Context) {
	var req createItemRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	arg := db.CreateItemParams{
		OwnerID:  req.OwnerID,
		ImageUrl: req.ImageURL,
		Category: intoNullString(req.Category), // Uses the helper in helpers.go
		Color:    intoNullString(req.Color),
		Style:    intoNullString(req.Style),
	}

	item, err := h.store.CreateItem(context.Background(), arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, item)
}

// ListItems handles GET /items?owner_id=1
func (h *Handler) ListItems(ctx *gin.Context) {
	ownerIDStr := ctx.Query("owner_id")
	if ownerIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "owner_id query parameter is required"})
		return
	}

	ownerID, err := strconv.Atoi(ownerIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid owner_id"})
		return
	}

	items, err := h.store.ListItemsByUser(context.Background(), int32(ownerID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch items"})
		return
	}

	ctx.JSON(http.StatusOK, items)
}
