package handlers

import (
	"deteleng-backend/internal/models"
	"deteleng-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	service *services.ReviewService
}

func NewReviewHandler(service *services.ReviewService) *ReviewHandler {
	return &ReviewHandler{service: service}
}

// GetPublicReviews handles GET /api/reviews (public, approved reviews only)
func (h *ReviewHandler) GetPublicReviews(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "50")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit > 50 {
		limit = 50
	}

	reviews, err := h.service.GetApprovedReviews(limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// GetAllReviews handles GET /api/admin/reviews (admin only)
func (h *ReviewHandler) GetAllReviews(c *gin.Context) {
	reviews, err := h.service.GetAllReviews()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

// CreateReview handles POST /api/admin/reviews
func (h *ReviewHandler) CreateReview(c *gin.Context) {
	var review models.Review

	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateReview(&review); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Review created successfully",
		"review":  review,
	})
}

// UpdateReview handles PUT /api/admin/reviews/:id
func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	id := c.Param("id")

	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review.ID = id
	if err := h.service.UpdateReview(&review); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Review updated successfully",
		"review":  review,
	})
}

// DeleteReview handles DELETE /api/admin/reviews/:id
func (h *ReviewHandler) DeleteReview(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteReview(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review deleted successfully"})
}

// GetReviewSources handles GET /api/admin/review-sources
func (h *ReviewHandler) GetReviewSources(c *gin.Context) {
	sources, err := h.service.GetReviewSources()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"sources": sources})
}

// CreateReviewSource handles POST /api/admin/review-sources
func (h *ReviewHandler) CreateReviewSource(c *gin.Context) {
	var req models.AddReviewSourceRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	source := &models.ReviewSource{
		Name:     req.Name,
		URL:      req.URL,
		IsActive: true,
	}

	if err := h.service.CreateReviewSource(source); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Review source created successfully",
		"source":  source,
	})
}

// UpdateReviewSource handles PUT /api/admin/review-sources/:id
func (h *ReviewHandler) UpdateReviewSource(c *gin.Context) {
	id := c.Param("id")

	var source models.ReviewSource
	if err := c.ShouldBindJSON(&source); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	source.ID = id
	if err := h.service.UpdateReviewSource(&source); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Review source updated successfully",
		"source":  source,
	})
}

// DeleteReviewSource handles DELETE /api/admin/review-sources/:id
func (h *ReviewHandler) DeleteReviewSource(c *gin.Context) {
	id := c.Param("id")

	if err := h.service.DeleteReviewSource(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Review source deleted successfully"})
}

// ParseReviews handles POST /api/admin/reviews/parse
func (h *ReviewHandler) ParseReviews(c *gin.Context) {
	var req models.ParseReviewsRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Limit <= 0 || req.Limit > 50 {
		req.Limit = 50
	}

	reviews, err := h.service.ParseReviewsFromSources(req.SourceIDs, req.Limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Reviews parsed successfully",
		"count":   len(reviews),
		"reviews": reviews,
	})
}

// GetReviewStats handles GET /api/admin/reviews/stats
func (h *ReviewHandler) GetReviewStats(c *gin.Context) {
	stats, err := h.service.GetReviewStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"stats": stats})
}
