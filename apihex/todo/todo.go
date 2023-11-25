package todo

import (
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

func List(c *gin.Context) {
	c.JSON(http.StatusOK, []Todo{})
}
