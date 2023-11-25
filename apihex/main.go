package main

import (
	"context"
	"database/sql"
	"log"
	"log/slog"
	_ "modernc.org/sqlite"
	"net/http"
	"os/signal"
	"syscall"
	"time"
	"todoapi/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	db, err := sql.Open("sqlite", "./database.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := gin.Default()
	message := "pong"
	// r.GET("/ping",handler(message).pingPongHandler)
	r.GET("/ping", func(ctx *gin.Context) {
		pingpong(message, ctx)
	})

	handler := todo.NewHandler(db)
	r.GET("/todos", handler.List)
	r.GET("/transfer/:id", func(c *gin.Context) {
		id := c.Param("id")
		slog.Info("parsing...", slog.String("id", id))
		time.Sleep(time.Millisecond * 200)
		slog.Info("validating...", slog.String("id", id))
		time.Sleep(time.Millisecond * 100)
		slog.Info("staging...", slog.String("id", id))
		time.Sleep(time.Millisecond * 200)
		slog.Info("transection starting...", slog.String("id", id))
		time.Sleep(time.Millisecond * 300)
		slog.Info("drawing...", slog.String("id", id))
		time.Sleep(time.Millisecond * 400)
		slog.Info("depositing...", slog.String("id", id))
		time.Sleep(time.Millisecond * 400)
		slog.Info("transection ending...", slog.String("id", id))
		time.Sleep(time.Millisecond * 100)
		slog.Info("responding...", slog.String("id", id))
		time.Sleep(time.Millisecond * 100)
		slog.Info("finish", slog.String("id", id))
		c.JSON(http.StatusOK, map[string]string{
			"message": "success",
		})
	})
	// r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

func (h handler) pingPongHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": string(h),
	})
}

type handler string

func pingpong(s string, c *gin.Context) {
	c.JSON(200, gin.H{
		"message": s,
	})
}
