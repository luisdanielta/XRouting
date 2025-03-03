// HTTP Handlers para Comment
package api

import (
	"net/http"
	"xrouting/internal/domain/entities"
	"xrouting/internal/ports"

	"github.com/labstack/echo/v4"
)

type CommentHandler struct {
	commentService ports.CommentService
}

func NewCommentHandler(commentService ports.CommentService) *CommentHandler {
	return &CommentHandler{commentService: commentService}
}

func (h *CommentHandler) RegisterComment(c echo.Context) error {
	comment := new(entities.Comment)
	if err := c.Bind(comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	uid := GenerateRandomUID()
	comment.ID = uid

	tableName := "comments"

	err := h.commentService.RegisterComment(c.Request().Context(), tableName, comment)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, comment)
}

func (h *CommentHandler) FindComment(c echo.Context) error {
	commentID := c.Param("id")
	tableName := "comments"

	comment, err := h.commentService.FindComment(c.Request().Context(), tableName, commentID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Comment not found"})
	}

	return c.JSON(http.StatusOK, comment)
}

func (h *CommentHandler) RemoveComment(c echo.Context) error {
	commentID := c.Param("id")
	tableName := "comments"

	err := h.commentService.RemoveComment(c.Request().Context(), tableName, commentID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Comment deleted successfully"})
}
