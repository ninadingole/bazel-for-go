package pkg

import "net/http"

type Message struct {
	Message string `json:"message"`
}

func HelloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message": "Hello, world."}`))
	}
}
