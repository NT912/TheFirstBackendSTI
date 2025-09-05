package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"nhatruong/firstGoBackend/internal/config"
	"nhatruong/firstGoBackend/internal/controllers"
	"nhatruong/firstGoBackend/internal/repository"
	"nhatruong/firstGoBackend/internal/routes"
	"nhatruong/firstGoBackend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// 1. Load cấu hình
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 2. Kết nối PostgreSQL bằng pgxpool (tương thích với repository của bạn)
	if err != nil {
		log.Fatalf("❌ failed to load config: %v", err)
	}

	ctxConn, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ✅ kết nối DB trực tiếp từ DATABASE_URL
	pool, err := pgxpool.New(ctxConn, cfg.DBUrl)
	if err != nil {
		log.Fatalf("❌ failed to connect DB: %v", err)
	}
	defer pool.Close()

	// 3. Khởi tạo repository, service, controller
	userRepo := repository.NewUserRepository(pool)
	authService := services.NewAuthService(userRepo, cfg.JWT.Secret)
	authController := controllers.NewAuthController(authService)

	// 4. Tạo router (gắn controller)
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := routes.SetupRouter(authController)

	// 5. Cấu hình và chạy HTTP server với graceful shutdown
	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      r,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  60 * time.Second,
	}

	// chạy server trong goroutine để có thể xử lý shutdown
	go func() {
		log.Printf("🚀 Server starting on %s\n", serverAddr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// chờ signal để tắt sạch sẽ
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()
	if err := srv.Shutdown(ctxShutdown); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server stopped gracefully")
}
