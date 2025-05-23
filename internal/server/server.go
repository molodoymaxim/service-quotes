package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
)

type HTTPServer interface {
	Start(c chan os.Signal, router http.Handler) error
}

type server struct {
	port int
}

func New(p int) HTTPServer {
	return &server{
		port: p,
	}
}

func (s *server) Start(c chan os.Signal, handler http.Handler) error {
	addr := fmt.Sprintf(":%d", s.port)

	srv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	// Запуск сервера в отдельной горутине
	errCh := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- fmt.Errorf("failed to start server: %w", err)
		}
	}()

	// Ждём сигнал для graceful shutdown
	<-c

	// Контекст с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	return nil
}
