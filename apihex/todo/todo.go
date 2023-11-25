package todo

import (
	"database/sql"
	_ "modernc.org/sqlite"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeteledAt time.Time `json:"deteled_at"`
}

func (h *Handler) List(c *gin.Context) {

	c.JSON(http.StatusOK, []Todo{})
}

func (h *Handler) NewTask(c *gin.Context) {
	var t Todo
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
		return
	}
	_, err := h.db.Exec("INSERT INTO todos (title) values (?)", t.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}
