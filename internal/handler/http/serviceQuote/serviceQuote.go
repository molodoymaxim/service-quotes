package serviceQuote

import (
	"encoding/json"
	"github.com/molodoymaxim/service-quotes/internal/types"
	"net/http"
	"strconv"
)

func (qh *QuoteServiceHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote types.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if err := qh.quotaService.Create(r.Context(), &quote); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quote)
}

func (qh *QuoteServiceHandler) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes, err := qh.quotaService.GetAll(r.Context(), author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quotes)
}

func (qh *QuoteServiceHandler) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := qh.quotaService.GetRandom(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(quote)
}

func (qh *QuoteServiceHandler) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	idValue := r.Context().Value("id")
	if idValue == nil {
		http.Error(w, "missing id in context", http.StatusBadRequest)
		return
	}

	idStr, ok := idValue.(string)
	if !ok {
		http.Error(w, "invalid id format", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := qh.quotaService.DeleteByID(r.Context(), id); err != nil {
		// Например, если не найдено — можно вернуть 404
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
