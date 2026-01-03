package order

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

/* ---------- CREATE ORDER ---------- */

type createOrderRequest struct {
	OrderID string `json:"order_id"`
	UserID  string `json:"user_id"`
}

func (h *Handler) CreateOrder(c *gin.Context) {
	var req createOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.service.CreateOrder(c, req.OrderID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created",
		"status":  StatusPlaced,
	})
}

/* ---------- CHANGE STATUS (already exists) ---------- */

type changeStatusRequest struct {
	Status OrderStatus `json:"status"`
}

func (h *Handler) ChangeOrderStatus(c *gin.Context) {
	orderID := c.Param("id")

	var req changeStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	err := h.service.ChangeStatus(c, orderID, req.Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order status updated"})
}
func (h *Handler) GetOrder(c *gin.Context) {
	id := c.Param("id")

	order, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(404, gin.H{"error": "order not found"})
		return
	}

	c.JSON(200, gin.H{
		"id":     order.ID,
		"userId": order.UserID,
		"status": order.Status,
	})
}

