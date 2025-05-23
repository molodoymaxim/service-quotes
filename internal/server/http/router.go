package handler

import (
	"context"
	"github.com/molodoymaxim/service-quotes/internal/handler"
	"net/http"
	"strings"
)

type HTTPRouter interface {
	InitRoutes() http.Handler
}

type router struct {
	http *handler.Server
}

func New(h *handler.Server) HTTPRouter {
	return &router{
		http: h,
	}
}

// InitRoutes создаёт и возвращает роутер с отдельными маршрутами.
func (rt router) InitRoutes() http.Handler {
	mux := http.NewServeMux()

	// Создание цитаты — POST /quotes
	mux.HandleFunc("/quotes", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			rt.http.QuoteService.CreateQuote(w, r)
			return
		}
		if r.Method == http.MethodGet {
			rt.http.QuoteService.GetAllQuotes(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	// Получение случайной цитаты — GET /quotes/random
	mux.HandleFunc("/quotes/random", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			rt.http.QuoteService.GetRandomQuote(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	// Удаление цитаты по ID — DELETE /quotes/{id}
	// Отдельная функция-обработчик для динамического маршрута
	mux.HandleFunc("/quotes/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Извлечение id из URL пути
		id := strings.TrimPrefix(r.URL.Path, "/quotes/")
		if id == "" || strings.Contains(id, "/") {
			http.Error(w, "invalid quote id", http.StatusBadRequest)
			return
		}

		// Кладём id в контекст
		ctx := context.WithValue(r.Context(), "id", id)
		r = r.WithContext(ctx)

		rt.http.QuoteService.DeleteQuoteByID(w, r)
	})

	return mux
}
